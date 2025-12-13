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
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/xyproto/randomstring"
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
// @Failure 403 {string} string "Неверный код входа"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id} [post]
func (h *NamespaceHandler) EnterNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	parts := strings.Split(r.URL.Path, "/")

	namespace := parts[3]
	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	invite_code := r.PostFormValue("invite_code")

	var namespaceInviteCode database.NamespaceInvite

	err := h.DB.Table("namespace_invite").First(&namespaceInviteCode, "namespace = ?", namespace).Error
	if err != nil {
		log.Printf("failed to find namespace code! %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if invite_code != namespaceInviteCode.InviteCode {
		log.Println("wrong invite code!")
		http.Error(w, "internal server error", http.StatusForbidden)
		return
	}

	err = h.DB.Table("user_namespace").Create(map[string]interface{}{
		"namespace": namespace, "user": tgId,
	}).Error
	if err != nil {
		log.Printf("failed to enter namespace! %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`you entered namespace`))
	w.Write([]byte("\n\n"))
}

// EnterNamespace
// @Summary Выйти из неймспейса
// @Param namespace query string true "id неймспейса"
// @Success 200 {string} string "Успешный вход"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id} [delete]
func (h *NamespaceHandler) LeaveNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	parts := strings.Split(r.URL.Path, "/")

	namespace := parts[3]

	// type UserNamespace struct{}
	err := h.DB.Debug().Table("user_namespace").Where(`"user" = ? AND namespace = ?`, tgId, namespace).Delete(&struct{}{}).Error
	if err != nil {
		log.Println("failed to leave namespace")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`you leaved namespace`))
	w.Write([]byte("\n\n"))
}

// CreateNamespace
// @Summary Создать неймспейс
// @Accept multipart/form-data
// @Param title formData string true "Название неймспейса"
// @Param photo formData file true "Картинка неймспейса"
// @Param description formData string true "Описание неймспейса"
// @Success 200 {string} string "Неймспейс успешно создан"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/namespace [post]
func (h *NamespaceHandler) CreateNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	title := r.PostFormValue("title")

	picture, _, err := r.FormFile("photo")
	if err != nil {
		log.Printf("failed to get file: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	pictureBytes, err := io.ReadAll(picture)
	if err != nil {
		log.Printf("failed to read picture bytes: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	picture.Close()

	description := r.PostFormValue("description")

	namespaceId := uuid.NewString()

	tx := h.DB.Begin()
	err = tx.Table("namespaces").Create(database.Namespace{
		Id:          namespaceId,
		Title:       title,
		Picture:     pictureBytes,
		Description: description,
		Admin:       tgId,
	}).Error
	if err != nil {
		tx.Rollback()
		log.Println("failed to create namespace!")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = tx.Table("namespace_invite").Create(database.NamespaceInvite{
		Namespace:         namespaceId,
		InviteCode: randomstring.CookieFriendlyString(20),
	}).Error
	if err != nil {
		tx.Rollback()
		log.Println("failed to create namespace invite code!")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	tx.Commit();

	w.Write([]byte(`namespace created`))
	w.Write([]byte("\n\n"))
}

// UpdateNamespace
// @Summary Обновить информацию о неймспейсе
// @Accept multipart/form-data
// @Param title formData string false "Название неймспейса"
// @Param photo formData file false "Картинка неймспейса"
// @Param description formData string false "Описание неймспейса"
// @Success 200 {string} string "Неймспейс успешно обновлен"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/namespace [put]
func (h *NamespaceHandler) UpdateNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	id := r.PostFormValue("id")

	var namespace database.Namespace
	err := h.DB.Table("namespaces").First(&namespace, "id = ?", id).Error
	if err != nil {
		log.Println("failed to get namespace!")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if namespace.Admin != tgId {
		log.Println("access denied!")
		http.Error(w, "you are not an admin", http.StatusForbidden)
		return
	}
	title := r.PostFormValue("title")

	picture, _, err := r.FormFile("photo")
	var pictureBytes []byte
	if err != nil {
		log.Printf("no file file in form: %s", err.Error())
	} else {
		pictureBytes, err = io.ReadAll(picture)
		if err != nil {
			log.Printf("failed to read picture bytes: %s", err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		picture.Close()
	}

	description := r.PostFormValue("description")

	err = h.DB.Model(&database.Namespace{Id: id}).Updates(database.Namespace{
		Title:       title,
		Picture:     pictureBytes,
		Description: description,
	}).Error

	if err != nil {
		log.Println("failed to update namespace!")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`namespace updated`))
	w.Write([]byte("\n\n"))
}

// DeleteNamespace
// @Summary Удалить неймспейс
// @Param id query string true "Название неймспейса"
// @Success 200 {string} string "Неймспейс успешно удален"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Failure 403 {string} string "Этот пользователь не админ неймспейса"
// @Router /api/namespace [delete]
func (h *NamespaceHandler) DeleteNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey)

	id := r.URL.Query().Get("id")
	var namespace database.Namespace
	err := h.DB.Table("namespaces").First(&namespace, "id = ?", id).Error
	if err != nil {
		log.Printf("failed to get namespace! %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if namespace.Admin != tgId {
		log.Println("access denied!")
		http.Error(w, "you are not an admin", http.StatusForbidden)
		return
	}

	err = h.DB.Delete(namespace).Error
	if err != nil {
		log.Println("failed to get namespace!")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`namespace deleted`))
	w.Write([]byte("\n\n"))
}

func getUserById(tx *gorm.DB, id string) (database.User, error) {
	var user database.User
	result := tx.Table("users").First(&user, "tg_id = ?", id)
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
// @Param id path string true "id неймспейса"
// @Success 200 {array} database.User
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/namespace/{namespace_id} [get]
func (h *NamespaceHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	namespace := strings.Split(r.URL.Path, "/")[3]
	log.Println(namespace)

	var viewsResp []UserView
	q := `
      WITH user_view_dates AS (
            SELECT
                u.tg_id AS user_id,
                MAX(v.date) AS last_viewed
            FROM users u
                LEFT JOIN views v
                    ON v.page_owner = u.tg_id AND v.viewer = $1
				INNER JOIN user_namespace un
					ON un.user = u.tg_id AND un.namespace = $2
            WHERE u.tg_id != $1
              AND u.name != ''
              AND u.interests IS NOT NULL
            GROUP BY u.tg_id
        ),
        total_users AS (
            SELECT COUNT(*) AS total FROM users WHERE tg_id != $1
        ),
        limited_users AS (
            SELECT * FROM user_view_dates 
                     ORDER BY last_viewed
                     ASC NULLS FIRST 
                     LIMIT (
                         SELECT LEAST(100, GREATEST(3, CAST(total * 0.2 AS INT))) FROM total_users
                     )
        ) SELECT * FROM limited_users;
  `
	err := h.DB.Raw(q, tgId, namespace).Scan(&viewsResp).Error
	if err != nil {
		log.Printf("failed to get views %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
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
		log.Printf("failed to get views %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	log.Println(string(workerRequestJSON))

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/feed/query", os.Getenv("WORKER_ENDPOINT")), bytes.NewReader(workerRequestJSON))
	if err != nil {
		log.Printf("worker/feed: failed to create request: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WORKER_SECRET")))

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("worker/feed: failed to process request: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("worker/feed: failed to read response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var respJSON map[string]interface{}
	err = json.Unmarshal(body, &respJSON)

	log.Println(respJSON)
	if err != nil {
		log.Printf("worker/feed: failed to read response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	matches := respJSON["matches"].([]interface{})

	pages := []database.User{}
	for _, match := range matches {
		pageId, ok := match.(map[string]interface{})["id"].(string)
		if !ok {
			log.Printf("worker/feed: failed to get user page: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		page, err := getUserById(h.DB, pageId)
		if err != nil {
			log.Printf("failed to get user page: %v", err)
			log.Printf("worker/feed: failed to read response: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		pages = append(pages, page)
	}

	pagesJSON, err := json.Marshal(pages)
	if err != nil {
		log.Printf("failed to serialize user: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(pagesJSON)
	w.Write([]byte("\n\n"))
}

func (h NamespaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var namespace string
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) >= 4 && parts[3] != "" {
		namespace = parts[3]
	}
	if namespace == "" {
		switch r.Method {
		case http.MethodPost:
			h.CreateNamespace(w, r)
		case http.MethodDelete:
			h.DeleteNamespace(w, r)
		case http.MethodPut:
			h.UpdateNamespace(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		switch r.Method {
		case http.MethodPost:
			h.EnterNamespace(w, r)
		case http.MethodDelete:
			h.LeaveNamespace(w, r)
		case http.MethodGet:
			h.GetFeed(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
