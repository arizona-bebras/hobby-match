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

	"gorm.io/gorm"
)

type WorkerHandler struct {
	DB *gorm.DB
}

func (_ *WorkerHandler) Autocomplete(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("query")
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/interests/query", os.Getenv("WORKER_ENDPOINT")), strings.NewReader(fmt.Sprintf(`{ "query": "%s" }`, q)))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WORKER_SECRET")))
	if err != nil {
		log.Printf("worker/autocomplete: failed to create request: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("worker/autocomplete: failed to process request: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("worker/autocomplete: failed to read response body: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf(`{ "query": "%s", "response": %s}`, q, string(body))))
	w.Write([]byte("\n\n"))
}

type View struct {
	Id string `json:"id"`
}

type Match struct {
	Id string `json:"id"`
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

func (h *WorkerHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	var views []View
	q := `WITH user_view_dates AS ( SELECT u.id AS user_id, MAX(v.created) AS last_viewed FROM users u LEFT JOIN views v ON v.pageOwner = u.id AND v.viewer = ? WHERE u.id != ? AND u.miniapp_name <> '' AND u.user_photo <> '' AND u.interests <> '[]' AND NOT u.hide AND NOT EXISTS ( SELECT 1 FROM bans b WHERE b.telegram_id = u.telegram_id ) GROUP BY u.id ), total_users AS ( SELECT COUNT(*) AS total FROM users WHERE id != ? ), limited_users AS ( SELECT * FROM user_view_dates ORDER BY last_viewed ASC NULLS FIRST LIMIT ( SELECT MIN(100, MAX(3, CAST(total * 0.2 AS INT))) FROM total_users ) ) SELECT * FROM limited_users;`
	err := h.DB.Raw(q, tgId).Scan(&views).Error
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

func (h WorkerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	switch path {
	case "/api/worker/autocomplete":
		switch r.Method {
		case http.MethodGet:
			h.Autocomplete(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "unknow endpoint", http.StatusNotFound)
	}
}