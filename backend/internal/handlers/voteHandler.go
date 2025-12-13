package handlers

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"log"

	"shumi/internal/database"
)

type VoteHandler struct {
	DB *gorm.DB
}

// Vote
// @Summary Проголосовать в опросе
// @Accept multipart/form-data
// @Param survey formData string false "id виджета опроса"
// @Param option formData string false "вариант опроса"
// @Success 200 {string} string "Голос засчитан"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/vote [post]
func (h *VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {
	tgID := r.Context().Value(database.AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	survey := r.PostFormValue("survey")
	option, err := strconv.Atoi(r.PostFormValue("option"))
	if err != nil {
		log.Println("failed to vote")
		http.Error(w, "interbal server error", http.StatusInternalServerError)
		return
	}

	vote := database.Vote{
		User: tgID,
		Survey: survey,
		Option: option,
	}

	result := h.DB.Create(&vote)
	if result.Error != nil {
		log.Println("failed to vote")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`voted`))
	w.Write([]byte("\n\n"))
}

func (h VoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/api/vote/":
		switch r.Method {
		case http.MethodPost:
			h.Vote(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "endpoint not found", http.StatusNotFound)
	}
}