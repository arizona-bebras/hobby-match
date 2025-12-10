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

func (h *NamespaceHandler) EnterNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	namespace := r.PostFormValue("namespace")

	err := h.DB.Table("namespace_user").Create(map[string]interface{}{
		"namespace": namespace, "user": tgId,
	}).Error
	if err != nil {
		log.Println("failed to enter namespace!")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`you entered namespace`))
	w.Write([]byte("\n\n"))
}

func (h *NamespaceHandler) LeaveNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	namespace := r.URL.Query().Get("namespace")

	type NamespaceUser struct{}
	err := h.DB.Delete(&NamespaceUser{}, "user = $1 AND namespace = $2", tgId, namespace).Error
	if err != nil {
		log.Println("failed to leave namespace")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`you entered namespace`))
	w.Write([]byte("\n\n"))
}

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

	err = h.DB.Table("namespaces").Create(database.Namespace{
		Id:          namespaceId,
		Title:       title,
		Picture:     pictureBytes,
		Description: description,
		Admin:       tgId,
	}).Error
	if err != nil {
		log.Println("failed to create namespace!")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = h.DB.Table("namespace_invite").Create(database.NamespaceInvite{
		Id:         namespaceId,
		InviteCode: randomstring.String(20),
	}).Error
	if err != nil {
		log.Println("failed to create namespace invite code!")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`namespace created`))
	w.Write([]byte("\n\n"))
}

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

func (h *NamespaceHandler) DeleteNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey)

	id := r.URL.Query().Get("id")
	var namespace database.Namespace
	err := h.DB.Table("namespaces").First(&namespace, "id = ?", id)
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

	err = h.DB.Delete(namespace)
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
