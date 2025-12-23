package handlers

import (
	"encoding/json"
	"io"
	"log"
	"fmt"
	"time"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"crypto/rand"

	"shumi/internal/database"
)

type TgUsersHandler struct {
	DB *gorm.DB
}

// GetUser
// @Summary Проверить существование пользователя в бд
// @Produce json
// @Param id query string true "id пользователя"
// @Success 200 {array} database.User
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/tg [get]
func (h *TgUsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var user database.User
	err := h.DB.Select("id", "hide").First(&user, "id = ?", id).Error
	if err != nil {
		log.Printf("tg handler: failed to get user %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to get user %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Printf("tg handler: failed to marshal user data %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to marshal user data %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	w.Write(userJSON)
	w.Write([]byte("\n\n"))
}

// RegisterUser
// @Summary Проверить существование пользователя в бд
// @Produce json
// @Param id body string true "tg id пользователя"
// @Param username body string true "tg username пользователя"
// @Param firstname body string true "tg firstname пользователя"
// @Success 200 {object} nil
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/tg [post]
func (h *TgUsersHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("tg handler: failed to read body %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to read body %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	var regData RegisterData
	err = json.Unmarshal(body, &regData)
	if err != nil {
		log.Printf("tg handler: failed to unmarshal reg data %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to unmarshal reg data %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	result := h.DB.Find(&database.User{}, "id = ?", regData.TgId)
	if result.Error != nil {
		log.Printf("tg handler: failed to find user %v", result.Error)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to find user %v", result.Error),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	if result.RowsAffected > 0 {
		log.Println("tg handler: user already exsists")
		http.Error(
			w, 
			database.JSONErr(
				http.StatusBadRequest, 
				"tg handler: user already exsists",
			), 
			http.StatusBadRequest,
		)
		return
	}

	tx := h.DB.Begin()
	err = tx.Create(&database.User{
		Id: regData.TgId,
		Hide: false,
	}).Error
	if err != nil {
		tx.Rollback()
		log.Printf("tg handler: failed to create user %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to find user %v", result.Error),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	err = tx.Create(&database.TgUser{
		UserId:      regData.TgId,
		TgUsername:  regData.Username,
		TgFirstname: regData.Firstname,
	}).Error
	if err != nil {
		tx.Rollback()
		log.Printf("tg handler: failed to create tg user %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to find user %v", result.Error),
			), 
			http.StatusInternalServerError,
		)
		return
	}
	tx.Commit()
}

// UpdateHideStatus
// @Summary Скрыть/показывать анкету другим пользователям
// @Produce json
// @Param id query string true "id пользователя"
// @Success 200 {object} nil
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/tg [patch]
func (h *TgUsersHandler) UpdateHideStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	tx := h.DB.Begin()

	var user database.User
	err := tx.First(&user, "id = ?", id).Error
	if err != nil {
		tx.Rollback()
		log.Printf("tg handler: failed to get user, %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to get user %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	user.Hide = !user.Hide
	err = tx.Save(&user).Error
	if err != nil {
		tx.Rollback()
		log.Printf("tg handler: failed to update hide status, %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to update hide status, %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	tx.Commit()
}

// DeleteUser
// @Summary Скрыть/показывать анкету другим пользователям
// @Produce json
// @Param id query string true "id пользователя"
// @Success 200 {object} nil
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/tg [delete]
func (h *TgUsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := h.DB.Delete(&database.User{
		Id: id,
	}).Error
	if err != nil {
		log.Printf("tg handler: failed to delete user %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to delete user %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}
}

// EnterNamespace
// @Summary Записать пользоватея в неймспейс
// @Produce json
// @Param data body TgEnterNamespaceData true "данные для записи пользователя"
// @Success 200 {object} nil
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/tg/namespace/add [post]
func (h *TgUsersHandler) EnterNamespace(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("tg handler: failed to read body %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to read body %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	var enterData TgEnterNamespaceData
	err = json.Unmarshal(body, &enterData)
	if err != nil {
		log.Printf("tg handler: failed to unmarshal reg data %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("tg handler: failed to unmarshal reg data %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	rows := h.DB.Table("user_namespace").First(&database.UserNamespace{}, "user_id = ? AND namespace_id = ?", enterData.UserId, enterData.NamespaceId).RowsAffected
	if rows != 0 {
		log.Println("tg handler: user already in namespace")
		return
	}

	err = h.DB.Create(&database.UserNamespace{
		UserId: enterData.UserId,
		NamespaceId: enterData.NamespaceId,
		Date: time.Now().Format(time.RFC3339),
	}).Error
	if err != nil {
		log.Printf("tg handler: failed to enter namespace, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("tg handler: failed to enter namespace, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}
}

// CreateNamespace
// @Summary Записать пользоватея в неймспейс
// @Produce json
// @Param data body TgEnterNamespaceData true "данные для записи пользователя"
// @Success 200 {object} nil
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/tg/namespace/create [post]
func (h *TgUsersHandler) CreateNamespace(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	tgId := r.PostFormValue("admin_id")
	title := r.PostFormValue("title")

	picture, _, err := r.FormFile("picture")
	if err != nil {
		log.Printf("tg handler: failed to get file, %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("tg handler: failed to get file, %s", err.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}
	pictureBytes, err := io.ReadAll(picture)
	if err != nil {
		log.Printf("tg handler: failed to read picture bytes: %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("tg handler: failed to read picture bytes: %s", err.Error()),
			),
			http.StatusInternalServerError,
		)
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
		AdminId:     tgId,
	}).Error
	if err != nil {
		tx.Rollback()
		log.Println("namespace handler: failed to create namespace!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"namespace handler: failed to create namespace!",
			),
			http.StatusInternalServerError,
		)
		return
	}

	err = tx.Table("namespace_invite").Create(database.NamespaceInvite{
		NamespaceId: namespaceId,
		InviteCode:  rand.Text(),
	}).Error
	if err != nil {
		tx.Rollback()
		log.Println("failed to create namespace invite code!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"failed to create namespace invite code!",
			),
			http.StatusInternalServerError,
		)
		return
	}

	err = tx.Create(&database.UserNamespace{
		NamespaceId: namespaceId, 
		UserId: tgId, 
		Date: time.Now().Format(time.RFC3339),
	}).Error
	if err != nil {
		tx.Rollback()
		log.Printf("namespace admin handler: failed to enter namespace, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace admin handler: failed to enter namespace, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}
	tx.Commit()

	createdNamespace := CreatedNamespace{
		NamespaceId: namespaceId,
	}

	JSONCreatedNamespace, err := json.Marshal(createdNamespace)
	if err != nil {
		log.Printf("namespace admin handler: failed to marshal response, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace admin handler: failed to marshal response, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	w.Write(JSONCreatedNamespace)
}

func (h TgUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/tg" {
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
			http.Error(
				w, 
				database.JSONErr(
					http.StatusMethodNotAllowed, 
					"tg handler: method not allowed",
				), 
				http.StatusMethodNotAllowed,
			)
		}
	} else {
		if r.Method != http.MethodPost {
			http.Error(
				w, 
				database.JSONErr(
					http.StatusMethodNotAllowed, 
					"tg handler: method not allowed",
				), 
				http.StatusMethodNotAllowed,
			)
			return
		}
		switch r.PathValue("method") {
		case "add":
			h.EnterNamespace(w, r)
		case "create":
			h.CreateNamespace(w, r)
		default:
			http.Error(
				w, 
				database.JSONErr(
					http.StatusNotFound, 
					"tg handler: endpoint not found",
				), 
				http.StatusNotFound,
			)
		}
	}
}

