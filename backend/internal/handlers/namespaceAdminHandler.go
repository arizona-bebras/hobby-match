package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"shumi/internal/database"
	// "github.com/xyproto/randomstring"
	"gorm.io/gorm"
)

type NamespaceAdminHandler struct {
	DB *gorm.DB
}

// CreateNamespace
// @Summary Создать неймспейс
// @Accept multipart/form-data
// @Param title formData string true "Название неймспейса"
// @Param photo formData file true "Картинка неймспейса"
// @Param description formData string true "Описание неймспейса"
// @Success 200 {object} CreatedNamespace "Неймспейс успешно создан"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace [post]
func (h *NamespaceAdminHandler) CreateNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	title := r.PostFormValue("title")

	picture, _, err := r.FormFile("photo")
	if err != nil {
		log.Printf("namesapce handler: failed to get file, %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namesapce handler: failed to get file, %s", err.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}
	pictureBytes, err := io.ReadAll(picture)
	if err != nil {
		log.Printf("namespace handler: failed to read picture bytes: %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to read picture bytes: %s", err.Error()),
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

	inviteCode := rand.Text()
	err = tx.Table("namespace_invite").Create(database.NamespaceInvite{
		NamespaceId: namespaceId,
		InviteCode:  inviteCode,
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
	err = tx.Model(&database.Namespace{Id: namespaceId}).Omit("Members.*").Association("Members").Append(&database.User{Id: tgId})
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
		InviteCode:  inviteCode,
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

// UpdateNamespace
// @Summary Обновить информацию о неймспейсе
// @Accept multipart/form-data
// @Param title formData string false "Название неймспейса"
// @Param photo formData file false "Картинка неймспейса"
// @Param description formData string false "Описание неймспейса"
// @Success 200 {object} nil "Неймспейс успешно обновлен"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/namespace [patch]
func (h *NamespaceAdminHandler) UpdateNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	id := r.PostFormValue("id")

	var namespace database.Namespace
	err := h.DB.Table("namespaces").First(&namespace, "id = ?", id).Error
	if err != nil {
		log.Println("namespace handler: failed to get namespace!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"namespace handler: failed to get namespace!",
			),
			http.StatusInternalServerError,
		)
		return
	}
	if namespace.AdminId != tgId {
		log.Println("namespace handler: access denied!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusForbidden,
				"namespace handler: failed to update, you are not an admin!",
			),
			http.StatusForbidden,
		)
		return
	}
	title := r.PostFormValue("title")

	picture, _, err := r.FormFile("photo")
	var pictureBytes []byte
	if err != nil {
		log.Printf("no file file in form: %s", err.Error())
	} else {
		pictureBytes, err = io.ReadAll(picture)
		if err != nil {
			log.Printf("namespace handler: failed to read picture bytes: %s", err.Error())
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("namespace handler: failed to read picture bytes: %s", err.Error()),
				),
				http.StatusInternalServerError,
			)
			return
		}
		picture.Close()
	}

	description := r.PostFormValue("description")

	err = h.DB.Model(&database.Namespace{Id: id}).Updates(database.Namespace{
		Title:       title,
		Picture:     pictureBytes,
		Description: description,
	}).Error

	if err != nil {
		log.Println("namespace handler: failed to update namespace!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"namespace handler: failed to update namespace!",
			),
			http.StatusInternalServerError,
		)
		return
	}
}

// DeleteNamespace
// @Summary Удалить неймспейс
// @Param id query string true "Название неймспейса"
// @Success 200 {object} nil "Неймспейс успешно удален"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Failure 403 {object} database.Error "Этот пользователь не админ неймспейса"
// @Router /api/namespace [delete]
func (h *NamespaceAdminHandler) DeleteNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey)

	id := r.URL.Query().Get("id")
	var namespace database.Namespace
	err := h.DB.Table("namespaces").First(&namespace, "id = ?", id).Error
	if err != nil {
		log.Printf("namespace handler: failed to get namespace! %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace handler: failed to get namespace! %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}
	if namespace.AdminId != tgId {
		log.Println("namespace handler: access denied!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusForbidden,
				"namespace handler: failed to delete, you are not an admin!",
			),
			http.StatusForbidden,
		)
		return
	}

	err = h.DB.Delete(namespace).Error
	if err != nil {
		log.Println("namespace handler: failed to delete namespace!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"namespace handler: failed to delete namespace!",
			),
			http.StatusInternalServerError,
		)
		return
	}
}

func (h NamespaceAdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateNamespace(w, r)
	case http.MethodDelete:
		h.DeleteNamespace(w, r)
	case http.MethodPatch:
		h.UpdateNamespace(w, r)
	default:
		http.Error(
			w,
			database.JSONErr(
				http.StatusMethodNotAllowed,
				"namespace handler: method not allowed",
			),
			http.StatusMethodNotAllowed,
		)
	}
}
