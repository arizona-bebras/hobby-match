package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"shumi/internal/database"
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
// @Param namespace formData string true "id неймспейса"
// @Param invite_code formData string true "инвайт код для входа в неймспейс"
// @Success 200 {string} string "Успешный вход"
// @Failure 403 {object} database.Error "Неверный код входа"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id} [post]
func (h *NamespaceHandler) EnterNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	namespace := r.PathValue("namespace_id")
	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	invite_code := r.PostFormValue("invite_code")

	var namespaceInviteCode database.NamespaceInvite

	err := h.DB.Table("namespace_invite").First(&namespaceInviteCode, "namespace = ?", namespace).Error
	if err != nil {
		log.Printf("namespace handler: failed to find namespace code, %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("namespace handler: failed to find namespace code, %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	if invite_code != namespaceInviteCode.InviteCode {
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

	err = h.DB.Table("user_namespace").Create(map[string]interface{}{
		"namespace": namespace, "user": tgId,
	}).Error
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

	w.Write([]byte("you entered namespace"))
	w.Write([]byte("\n\n"))
}

// EnterNamespace
// @Summary Выйти из неймспейса
// @Param namespace query string true "id неймспейса"
// @Success 200 {string} string "Успешный вход"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id} [delete]
func (h *NamespaceHandler) LeaveNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	namespace := r.PathValue("namespace_id")

	// type UserNamespace struct{}
	err := h.DB.Debug().Table("user_namespace").Where(`"user" = ? AND namespace = ?`, tgId, namespace).Delete(&struct{}{}).Error
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

	w.Write([]byte("you leaved namespace"))
	w.Write([]byte("\n\n"))
}

func getUserById(tx *gorm.DB, id string) (PageData, error) {
	var user PageData
	result := tx.Table("users").
		Select(
			"users.id", 
			"users.name",
			"users.location",
			"users.gender",
			"users.birth_date",
			"users.interests",
			"users.photo",
			"users.info",
			"users.hide",
			"tg_users.username",
		).
		Joins("JOIN tg_users ON users.id = tg_users.id").
		First(&user, "users.id = ?", id)
	if result.Error != nil {
		return user, fmt.Errorf("failed to get user: %s", result.Error.Error())
	}

	var widgets []database.Widget
	result = tx.Table("widgets").Find(&widgets, "widgets.user = ?", id)
	if result.Error != nil {
		return user, fmt.Errorf("failed to get widgets: %s", result.Error.Error())
	}
	user.Widgets = widgets

	return user, nil
}

// GetFeed
// @Summary Получить ленту из анкет пользователей неймспейса
// @Produce json
// @Param namespace_id path string true "id неймспейса"
// @Success 200 {array} PageData
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id} [get]
func (h *NamespaceHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	namespace := r.PathValue("namespace_id")

	var viewsResp []UserView
	q := `
      WITH user_view_dates AS (
            SELECT
                u.id AS user_id,
                MAX(v.date) AS last_viewed
            FROM users u
                LEFT JOIN views v
                    ON v.page_owner_id = u.id AND v.viewer_id = $1
				INNER JOIN user_namespace un
					ON un.user_id = u.id AND un.namespace_id = $2
            WHERE u.id != $1
              AND u.name != ''
              AND u.interests IS NOT NULL
            GROUP BY u.id
        ),
        total_users AS (
            SELECT COUNT(*) AS total FROM users WHERE id != $1
        ) 
		SELECT * FROM user_view_dates
		WHERE last_viewed IS NULL;
  `
	err := h.DB.Raw(q, tgId, namespace).Scan(&viewsResp).Error
	if err != nil {
		log.Printf("namespace handler: failed to get views %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("namespace handler: failed to get views %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	var views []string

	for _, view := range viewsResp {
		views = append(views, view.UserID)
	}

	workerRequestBody := WorkerRequest{
		Id: tgId,
		From: views,
	}

	workerRequestJSON, err := json.Marshal(workerRequestBody)
	if err != nil {
		log.Printf("namespace handler: failed to get views %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("namespace handler: failed to get views %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	log.Println(string(workerRequestJSON))

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/feed/query", os.Getenv("WORKER_ENDPOINT")), bytes.NewReader(workerRequestJSON))
	if err != nil {
		log.Printf("worker/feed: failed to create request: %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("worker/feed: failed to create request: %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WORKER_SECRET")))

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("worker/feed: failed to process request: %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("worker/feed: failed to create request: %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("worker/feed: failed to read response: %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("worker/feed: failed to read response: %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	var respJSON map[string]interface{}
	err = json.Unmarshal(body, &respJSON)

	log.Println(respJSON)
	if err != nil {
		log.Printf("worker/feed: failed to read response: %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("worker/feed: failed to read response: %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	matches := respJSON["matches"].([]interface{})

	pages := []PageData{}
	for _, match := range matches {
		pageId, ok := match.(map[string]interface{})["id"].(string)
		if !ok {
			log.Printf("worker/feed: failed to get user page: %v", err)
			http.Error(
				w, 
				database.JSONErr(
					http.StatusInternalServerError, 
					fmt.Sprintf("worker/feed: failed to get user page: %v", err),
				), 
				http.StatusInternalServerError,
			)
			return
		}
		page, err := getUserById(h.DB, pageId)
		if err != nil {
			log.Printf("namespace handler: failed to get user page: %v", err)
			http.Error(
				w, 
				database.JSONErr(
					http.StatusInternalServerError, 
					fmt.Sprintf("namespace handler: failed to get user page: %v", err),
				), 
				http.StatusInternalServerError,
			)
			return
		}
		pages = append(pages, page)
	}

	pagesJSON, err := json.Marshal(pages)
	if err != nil {
		log.Printf("namespace handler: failed to serialize user: %s", err.Error())
		http.Error(
				w, 
				database.JSONErr(
					http.StatusInternalServerError, 
					fmt.Sprintf("namespace handler: failed to serialize user: %s", err.Error()),
				), 
				http.StatusInternalServerError,
			)
		return
	}

	w.Write(pagesJSON)
	w.Write([]byte("\n\n"))
}

func (h NamespaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.EnterNamespace(w, r)
	case http.MethodDelete:
		h.LeaveNamespace(w, r)
	case http.MethodGet:
		h.GetFeed(w, r)
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
