package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"

	"shumi/internal/database"
)

type FilesHandler struct {
	DB *gorm.DB
}

// GetPhoto
// @Summary Получить фото, связанное с объектом
// @Produce json
// @Param object path string true "тип объекта (users, namespaces, widgets)" 
// @Param id query string false "id объекта (не нужно передавать для users)" 
// @Success 200 {object} Photos
// @Failure 403 {object} database.Error "Для этого пользователя нет доступа к файлу"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/files/{object} [get]
func (h *FilesHandler) GetPhoto(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)
	object := r.PathValue("object")
	id := r.URL.Query().Get("id")

	switch object {
	case "users":
		var file []byte
		err := h.DB.Table("users").
			Select("photo").
			Where("id = ?", tgId).
			Row().
			Scan(&file)
		if err != nil {
			log.Printf("files handler: failed to get photo, %v", err)
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("files handler: failed to get photo, %v", err),
				),
				http.StatusInternalServerError,
			)
			return
		}

		photo := Photos{
			Files: [][]byte{file},
		}

		JSONPhoto, err := json.Marshal(photo)

		if err != nil {
			log.Printf("files handler: failed to marshal photo, %v", err)
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("files handler: failed to marshal photo, %v", err),
				),
				http.StatusInternalServerError,
			)
			return
		}

		w.Write(JSONPhoto)
	case "namespaces":
		var file []byte
		err := h.DB.Table("namespaces").
			Select("picture").
			Where("id = ?", id).
			Row().
			Scan(&file)
		if err != nil {
			log.Printf("files handler: failed to get photo, %v", err)
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("files handler: failed to get photo, %v", err),
				),
				http.StatusInternalServerError,
			)
			return
		}

		picture := Photos{
			Files: [][]byte{file},
		}

		JSONPicture, err := json.Marshal(picture)

		if err != nil {
			log.Printf("files handler: failed to marshal photo, %v", err)
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("files handler: failed to marshal photo, %v", err),
				),
				http.StatusInternalServerError,
			)
			return
		}

		w.Write(JSONPicture)
	case "widgets":
		var widget database.Widget
		err := h.DB.Table("widgets").First(&widget, "id = ?", id).Error
		if err != nil {
			log.Printf("files handler: failed to get photo, %v", err)
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("files handler: failed to get photo, %v", err),
				),
				http.StatusInternalServerError,
			)
			return
		}

		if tgId != widget.UserId {
			log.Println("files handler: access denied, that`s not your widget")
			http.Error(
				w,
				database.JSONErr(
					http.StatusForbidden,
					"files handler: access denied, that`s not your widget",
				),
				http.StatusForbidden,
			)
			return
		}

		photos := Photos{
			Files: widget.Files,
		}

		JSONPhotos, err := json.Marshal(photos)

		if err != nil {
			log.Printf("files handler: failed to marshal photo, %v", err)
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("files handler: failed to marshal photo, %v", err),
				),
				http.StatusInternalServerError,
			)
			return
		}

		w.Write(JSONPhotos)
	}
}

func (h FilesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("files handler: method allowed")
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				"files handler: method allowed",
			),
			http.StatusInternalServerError,
		)
		return
	}
	h.GetPhoto(w, r)
}