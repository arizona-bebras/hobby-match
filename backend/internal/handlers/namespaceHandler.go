package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"shumi/internal/database"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NamespaceHandler struct {
	DB *gorm.DB
}

type View struct {
	Id string `json:"id"`
}

type Match struct {
	Id string `json:"id"`
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

	err = h.DB.Table("namespaces").Create(database.Namespace{
		Id: uuid.NewString(),
		Title: title,
		Picture: pictureBytes,
		Description: description,
		Admin: tgId,
	}).Error
	if err != nil {
		log.Println("failed to create namespace!")
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

	err = h.DB.Model(&database.Namespace{ Id: id }).Updates(database.Namespace{
		Title: title,
		Picture: pictureBytes,
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

func getUserById(tx *gorm.DB, id string) (database.User, error){
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

	namespace := strings.Split(r.URL.Path, "/")[2]

	var views []View
	q := `
      WITH user_view_dates AS (
            SELECT
                u.id AS user_id,
                MAX(v.created) AS last_viewed
            FROM users u
                LEFT JOIN views v
                    ON v.pageOwner = u.id AND v.viewer = $1
				INNER JOIN user_namespaces un
					ON un.user = u.id AND un.namespace = $2
            WHERE u.id != $1
              AND u.miniapp_name <> ''
              AND u.user_photo <> ''
              AND u.interests <> '[]'
              AND NOT u.hide
              AND NOT EXISTS (
                SELECT 1
                FROM bans b
                WHERE b.telegram_id = u.telegram_id
              )
            GROUP BY u.id
        ),
        total_users AS (
            SELECT COUNT(*) AS total FROM users WHERE id != $1
        ),
        limited_users AS (
            SELECT * FROM user_view_dates 
                     ORDER BY last_viewed
                     ASC NULLS FIRST 
                     LIMIT (
                         SELECT MIN(100, MAX(3, CAST(total * 0.2 AS INT))) FROM total_users
                     )
        ) SELECT * FROM limited_users;
  `
	err := h.DB.Raw(q, tgId, namespace).Scan(&views).Error
	if err != nil {
		log.Printf("failed to get views %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	viewsJSON, err := json.Marshal(views)
	if err != nil {
		log.Printf("failed to get views %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/feed/query", os.Getenv("WORKER_ENDPOINT")), strings.NewReader(fmt.Sprintf(`{ "id": "%s", "from": %s }`, q, viewsJSON)))
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

	var matches []Match
	err = json.Unmarshal(body, &matches)
	if err != nil {
		log.Printf("worker/feed: failed to read response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	pages := []database.User{}
	for _, match := range matches {
		page, err := getUserById(h.DB, match.Id)
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