package handlers

import (
	"fmt"
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

// UpdateNamespace
// @Summary Обновить информацию о неймспейсе
// @Accept multipart/form-data
// @Param namespace_id path string false "id неймспейса"
// @Param title formData string false "Название неймспейса"
// @Param photo formData file false "Картинка неймспейса"
// @Param description formData string false "Описание неймспейса"
// @Success 200 {object} nil "Неймспейс успешно обновлен"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/admin/{namespace_id} [patch]
func (h *NamespaceAdminHandler) UpdateNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	id := r.PathValue("namespace_id")

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
// @Param namespace_id path string true "Название неймспейса"
// @Success 200 {object} nil "Неймспейс успешно удален"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Failure 403 {object} database.Error "Этот пользователь не админ неймспейса"
// @Router /api/admin/{namespace_id} [delete]
func (h *NamespaceAdminHandler) DeleteNamespace(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey)

	id := r.PathValue("namespace_id")
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
		log.Println("namespace admin handler: failed to delete namespace!")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"namespace admin handler: failed to delete namespace!",
			),
			http.StatusInternalServerError,
		)
		return
	}
}

// KickNamespaceMember
// @Summary Исключить пользователя из неймспейса
// @Param namespace_id path string true "id неймспейса"
// @Param member_id path string true "id пользователя"
// @Success 200 {object} nil "Участник исключен"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Failure 403 {object} database.Error "Этот пользователь не админ неймспейса"
// @Router /api/admin/{namespace_id}/{member_id} [delete]
func (h *NamespaceAdminHandler) KickNamespaceMember(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)

	namespaceId := r.PathValue("namespace_id")
	memberId := r.PathValue("member_id")

	var namespace database.Namespace
	err := h.DB.Model(&database.Namespace{}).First(&namespace, "id = ?", namespaceId).Error
	if err != nil {
		log.Printf("namespace admin handler: failed to find namespace, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace admin handler: failed to find namespace, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	if namespace.AdminId != tgId {
		log.Println("namespace admin handler: you are not an admin")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"namespace admin handler: you are not an admin",
			),
			http.StatusInternalServerError,
		)
		return
	}

	err = h.DB.Delete(database.UserNamespace{}, "user_id = ? AND namespace_id = ?", memberId, namespaceId).Error
	if err != nil {
		log.Printf("namespace admin handler: failed to kick member, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("namespace admin handler: failed to kick member, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}
}

func (h NamespaceAdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.PathValue("member_id") == "" {
		switch r.Method {
		case http.MethodPatch:
			h.UpdateNamespace(w ,r)
		case http.MethodDelete:
			h.DeleteNamespace(w, r)
		default:
		http.Error(
			w,
			database.JSONErr(
				http.StatusMethodNotAllowed,
				"namespace admin handler: metthod not allowed",
			),
			http.StatusMethodNotAllowed,
		)
		}
	} else {
		switch r.Method {
		case http.MethodDelete:
			h.KickNamespaceMember(w, r)
		default:
		http.Error(
			w,
			database.JSONErr(
				http.StatusMethodNotAllowed,
				"namespace admin handler: metthod not allowed",
			),
			http.StatusMethodNotAllowed,
		)
		}
	}
}
