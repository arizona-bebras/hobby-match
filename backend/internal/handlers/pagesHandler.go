package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"

	"shumi/internal/database"
)

type PagesHandler struct {
	DB *gorm.DB
}

// GetPage
// @Summary Получить анкету(страницу) одного пользователя
// @Produce json
// @Param page_id path string true "id анкеты"
// @Success 200 {array} handlers.PageData
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/pages/{page_id} [get]
func (h *PagesHandler) GetPage(w http.ResponseWriter, r *http.Request) {
	pageId := r.PathValue("page_id")

	var user PageData

	err := h.DB.Table("users").
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
		First(&user, "users.id = ?", pageId).Error
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

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Printf("pages handler: failed to marshal user data, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("pages handler: failed to marshal user data, %v", err),
			),
		http.StatusInternalServerError,
		)
		return
	}

	w.Write(userJSON)
}

func (h PagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		h.GetPage(w ,r)
	}
}