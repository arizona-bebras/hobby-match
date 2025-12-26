package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shumi/internal/database"
	"strings"
	"time"

	"gorm.io/gorm"
)

type NamespaceHandler struct {
	DB *gorm.DB
}

type WorkerRequest struct {
	Id   string   `json:"id"`
	From []string `json:"from"`
}

type UserView struct {
	UserID     string     `json:"user_id" db:"user_id"`
	LastViewed *time.Time `json:"last_viewed" db:"last_viewed"`
}

type Match struct {
	Id    string  `json:"id"`
	Score float64 `json:"score"`
}

// EnterNamespace
// @Summary Зайти в неймспейс
// @Accept multipart/form-data
// @Param namespace_id path string true "id неймспейса"
// @Param invite_code formData string true "инвайт код для входа в неймспейс"
// @Success 200 {object} nil "Успешный вход"
// @Failure 403 {object} database.Error "Неверный код входа"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id} [post]
func (h *NamespaceHandler) EnterNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	//namespace := r.PathValue("namespace_id")
	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	invite_code := r.PostFormValue("invite_code")

	var namespace database.Namespace

	err := h.DB.Table("namespace_invite").
		Joins("left join namespaces on namespaces.id = namespace_invite.namespace_id").
		Where("namespace_invite.invite_code = ?", invite_code).
		Scan(&namespace).Error
	if err != nil {
		log.Println("namespace handler: wrong invite code!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusForbidden,
				"namespace handler: wrong invite code!",
			),
			http.StatusForbidden,
		)
		return
	}

	err = h.DB.Model(&namespace).Omit("Members.*").Association("Members").Append(&database.User{Id: tgId})
	if err != nil {
		log.Printf("namespace handler: failed to enter namespace, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to enter namespace, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}
}

// LeaveNamespace
// @Summary Выйти из неймспейса
// @Param namespace_id path string true "id неймспейса"
// @Success 200 {object} nil "Успешный вход"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id} [delete]
func (h *NamespaceHandler) LeaveNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	namespace := r.PathValue("namespace_id")

	// type UserNamespace struct{}
	err := h.DB.Model(database.Namespace{Id: namespace}).Association("Members").Delete(database.User{Id: tgId})
	if err != nil {
		log.Println("namespace handler: failed to leave namespace")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"namespace handler: failed to leave namespace",
			),
			http.StatusInternalServerError,
		)
		return
	}
}

// GetFeed
// @Summary Получить ленту из анкет пользователей неймспейса
// @Produce json
// @Param namespace_id path string true "id неймспейса"
// @Success 200 {array} handlers.PageData
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id}/feed [get]
func (h *NamespaceHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	selfId := r.Context().Value(database.AuthContextKey).(string)
	namespace := r.PathValue("namespace_id")

	var views []string
	err := h.DB.Raw(`
		WITH user_vectors AS (
		  SELECT u.id, info_embedding, personality_test, AVG(i.embedding) as avg_interest_embedding FROM users AS u
		  LEFT JOIN user_interests ui ON u.id = ui.user_id
		  LEFT JOIN interests i ON i.id = ui.interest_id
		  LEFT JOIN user_namespace un ON u.id = un.user_id
		  WHERE un.namespace_id = @namespace
		  GROUP BY u.id
		), self_user AS (
		  SELECT * FROM user_vectors WHERE id = @self
		), scores AS (
		  SELECT id,
			coalesce(1 - (info_embedding <=> (SELECT info_embedding FROM self_user)), 0) AS info_score,
			coalesce(1 - (avg_interest_embedding <=> (SELECT avg_interest_embedding FROM self_user)), 0) AS interests_score,
			coalesce(1 - ((personality_test <-> (SELECT personality_test FROM self_user)) / SQRT(500)), 0) AS personality_score
		  FROM user_vectors WHERE id != @self
		), ranking AS (
		  SELECT id, info_score, interests_score, personality_score,
			(0.2 * info_score + 0.5 * interests_score + 0.4 * personality_score) AS similarity 
		  FROM scores
		  ORDER BY similarity DESC
		), total_views AS (
		  SELECT id AS page_owner_id, MAX(v.date) as date FROM users 
		  LEFT JOIN views v ON v.page_owner_id = users.id AND v.viewer_id = @self
		  GROUP BY users.id
		), ranked_recommendations AS (
			SELECT
				u.id,
				s.similarity,
				v.date,
				CASE WHEN v.date IS NULL THEN 1 ELSE 2 END AS priority
			FROM users u
			JOIN ranking s ON s.id = u.id
			LEFT JOIN total_views v ON v.page_owner_id = u.id
			WHERE u.id != @self
		)
		SELECT id FROM ranked_recommendations
		ORDER BY
			priority ASC,
			CASE WHEN priority = 1 THEN similarity END DESC,
			CASE WHEN priority = 1 THEN id END DESC,
			CASE WHEN priority = 2 THEN date END ASC,
			CASE WHEN priority = 2 THEN similarity END DESC
		LIMIT 3;
  	`, sql.Named("self", selfId), sql.Named("namespace", namespace)).Scan(&views).Error
	if err != nil {
		log.Printf("namespace handler: failed to get recommendations %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to get recommendations %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	viewRecords := make([]database.View, len(views))
	for i, view := range views {
		viewRecords[i] = database.View{
			ViewerId:    selfId,
			PageOwnerId: view,
		}
	}

	err = h.DB.Create(&viewRecords).Error
	if err != nil {
		log.Printf("namespace handler: failed to create views: %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to create views: %s", err.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}
	pages := make([]PageData, len(views))
	for i, view := range views {
		var page *PageData
		page, err = GetPageById(h.DB, selfId, view)
		if err != nil {
			log.Printf("pages handler: failed to get user, %v", err)
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("pages handler: failed to get user, %v", err),
				),
				http.StatusInternalServerError,
			)
			return
		}
		pages[i] = *page
	}

	pagesJSON, err := json.Marshal(pages)
	if err != nil {
		log.Printf("namespace handler: failed to serialize recommendations: %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to serialize recommendations: %s", err.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}

	w.Write(pagesJSON)
}

func checkMember(namespaceMembers []NamespaceMember, memberId string) bool {
	for _, namespaceMember := range namespaceMembers {
		if memberId == namespaceMember.Id {
			return true
		}
	}
	return false
}

// GetPages
// @Summary Получить пользователей неймспейса
// @Produce json
// @Param namespace_id path string true "id неймспейса"
// @Success 200 {object} handlers.NamespaceMembers
// @Failure 403 {object} database.Error "Пользователь не участник этого неймспейса"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id}/pages [get]
func (h *NamespaceHandler) GetPages(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	namespaceId := r.PathValue("namespace_id")
	var namespace database.Namespace
	var members []NamespaceMember

	err := h.DB.Model(&database.Namespace{}).First(&namespace, "id = ?", namespaceId).Error
	if err != nil {
		log.Printf("namespace handler: failed to get namespace data, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to get namespace data, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	err = h.DB.Model(&database.UserNamespace{}).
		Joins(`LEFT JOIN users ON users.id = user_namespace.user_id`).
		Select("users.id, users.name, user_namespace.date").
		Find(&members, "namespace_id = ?", namespaceId).Error
	if err != nil {
		log.Printf("namespace handler: failed to get namespace member`s ids, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to get namespace member`s ids, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	if !checkMember(members, tgId) {
		log.Println("namespace handler: access denied you aren`t namespace member!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusForbidden,
				"namespace handler: access denied you aren`t namespace member!",
			),
			http.StatusForbidden,
		)
		return
	}

	JSONNamespaceMembers, err := json.Marshal(NamespaceMembers{
		Namespace:        namespace,
		NamespaceMembers: members,
	})
	if err != nil {
		log.Printf("namespace handler: failed to marshal namespace data, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to marshal namespace data, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	w.Write(JSONNamespaceMembers)
}

func (h NamespaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.EnterNamespace(w, r)
	case http.MethodDelete:
		h.LeaveNamespace(w, r)
	case http.MethodGet:
		switch strings.Split(r.URL.Path, "/")[4] {
		case "feed":
			h.GetFeed(w, r)
		case "pages":
			h.GetPages(w, r)
		default:
			http.Error(
				w,
				database.JSONErr(
					http.StatusNotFound,
					"namespace handler: endpoint not found",
				),
				http.StatusNotFound,
			)
		}
	default:
		http.Error(
			w,
			database.JSONErr(
				http.StatusMethodNotAllowed,
				"namespace handler: method not allowed",
			),
			http.StatusMethodNotAllowed,
		)
	}
}
