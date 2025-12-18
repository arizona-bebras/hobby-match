package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"

	"shumi/internal/database"
)

type TgUsersHandler struct {
	DB *gorm.DB
}

func (h *TgUsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var user database.User
	err := h.DB.Select("id", "hide").First(&user, "id = ?", id).Error
	if err != nil {
		log.Printf("failed to get user %v", err)
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Printf("failed to marshal user data %v", err)
		http.Error(w, "failed to marshal user data", http.StatusInternalServerError)
		return
	}

	w.Write(userJSON)
	w.Write([]byte("\n\n"))
}

func (h *TgUsersHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body %v", err)
		http.Error(w, "failed to read body", http.StatusInternalServerError)
		return
	}

	var regData RegisterData
	err = json.Unmarshal(body, &regData)
	if err != nil {
		log.Printf("failed to unmarshal reg data %v", err)
		http.Error(w, "failed to unmarshal reg data", http.StatusInternalServerError)
		return
	}

	result := h.DB.Find(&database.User{}, "id = ?", regData.TgId)
	if result.Error != nil {
		log.Printf("failed to find user %v", result.Error)
		http.Error(w, "failed to find user", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected > 0 {
		log.Printf("user already exsists")
		http.Error(w, "user already exsists", http.StatusBadRequest)
		return
	}

	tx := h.DB.Begin()
	err = tx.Create(&database.User{
		Id: regData.TgId,
		Hide: false,
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

func (h *TgUsersHandler) UpdateHideStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	tx := h.DB.Begin()

	var user database.User
	err := tx.First(&user, "id = ?", id).Error
	if err != nil {
		tx.Rollback()
		log.Printf("failed to get user %v", err)
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}

	user.Hide = !user.Hide
	err = tx.Save(&user).Error
	if err != nil {
		tx.Rollback()
		log.Printf("failed to update hide status %v", err)
		http.Error(w, "failed to update hide status", http.StatusInternalServerError)
		return
	}

	tx.Commit()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\n\n"))
}

func (h *TgUsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := h.DB.Delete(&database.User{
		Id: id,
	}).Error
	if err != nil {
		log.Printf("failed to delete user %v", err)
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\n\n"))
}

func (h TgUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.RegisterUser(w, r)
	case http.MethodPatch:
		h.UpdateHideStatus(w, r)
	case http.MethodDelete:
		h.DeleteUser(w ,r)
	case http.MethodGet:
		h.GetUser(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
