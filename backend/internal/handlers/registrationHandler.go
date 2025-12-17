package handlers

import (
	"net/http"
	"log"

	"gorm.io/gorm"

	"shumi/internal/database"
)

type RegistrationHandler struct {
	DB *gorm.DB
}

func (h *RegistrationHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5000000)
	id := r.PostFormValue("id")
	result := h.DB.Find(&database.User{}, "id = ?", id)
	if result.Error != nil {
		log.Printf("failed to find user %v", result.Error)
		http.Error(w, "failed to find user", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected > 1 {
		log.Printf("user already exsists")
		http.Error(w, "user already exsists", http.StatusBadRequest)
		return
	}

	tx := h.DB.Begin()
	err := tx.Create(&database.User{
		Id: id,
	}).Error
	if err != nil {
		tx.Rollback()
		log.Printf("failed to create user %v", err)
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	err = tx.Create(&database.TgUser{
		UserId: id,
		TgUsername: r.PostFormValue("username"),
		TgFirstname: r.PostFormValue("firstname"),
	}).Error
	if err != nil {
		tx.Rollback()
		log.Printf("failed to create tg user %v", err)
		http.Error(w, "failed to create tg user", http.StatusInternalServerError)
		return
	}
	tx.Commit()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\n\n"))
}

func (h RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.RegisterUser(w, r)
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}