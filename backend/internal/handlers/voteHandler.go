package handlers

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"log"
	"fmt"

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
// @Success 200 {object} nil "Голос засчитан"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/vote [post]
func (h *VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {
	tgID := r.Context().Value(database.AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	survey := r.PostFormValue("survey")
	option, err := strconv.Atoi(r.PostFormValue("option"))
	if err != nil {
		log.Println("failed to vote")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	vote := database.Vote{
		UserId: tgID,
		WidgetId: survey,
		Option: option,
	}

	result := h.DB.Create(&vote)
	if result.Error != nil {
		log.Printf("vote handler: failed to vote, %s", result.Error.Error())
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("vote handler: failed to vote, %s", result.Error.Error()),
			), 
			http.StatusInternalServerError,
		)
		return
	}
}

func (h VoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/api/vote":
		switch r.Method {
		case http.MethodPost:
			h.Vote(w, r)
		default:
			http.Error(
			w, 
			database.JSONErr(
				http.StatusMethodNotAllowed, 
				"vote handler: method not allowed",
			), 
			http.StatusMethodNotAllowed,
		)
		}
	}
}