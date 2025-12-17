package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"

	"shumi/internal/database"
)

type RegistrationHandler struct {
	DB *gorm.DB
}

type RegisterData struct {
	TgId      string `json:"tg_id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
}

func (h *RegistrationHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body %v", err)
		http.Error(w, "failed to read body", http.StatusInternalServerError)
		return
	}

	var regData RegisterData
	err = json.Unmarshal(body, &regData)
	if err != nil {
		log.Printf("failed to marshal reg data %v", err)
		http.Error(w, "failed to marshal reg data", http.StatusInternalServerError)
		return
	}

	result := h.DB.Find(&database.User{}, "id = ?", regData.TgId)
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
	err = tx.Create(&database.User{
		Id: regData.TgId,
	}).Error
	if err != nil {
		tx.Rollback()
		log.Printf("failed to create user %v", err)
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	err = tx.Create(&database.TgUser{
		UserId:      regData.TgId,
		TgUsername:  regData.Username,
		TgFirstname: regData.Firstname,
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
	if r.Method == http.MethodPost {
		h.RegisterUser(w, r)
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
