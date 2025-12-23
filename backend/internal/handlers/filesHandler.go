package handlers

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"io"
	"log"
	//"mime/multipart"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"shumi/internal/database"
)

type FilesHandler struct {
	DB *gorm.DB
}

// GetPhoto
// @Summary Получить фото, связанное с объектом
// @Produce  application/octet-stream
// @Param object path string true "тип объекта (users, namespaces, widgets)" 
// @Param id path string true "id объекта" 
// @Param index path string false "индекс файла в списке файлов виджета" 
// @Success 200 {file} binary "Файл"
// @Failure 403 {object} database.Error "Для этого пользователя нет доступа к файлу"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/files/{object}/{id}/{index} [get]
func (h *FilesHandler) GetPhoto(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey).(string)
	object := r.PathValue("object")
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/octet-stream")

	switch object {
	case "users":
		var file []byte
		err := h.DB.Table("users").
			Select("photo").
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

		binaryFile := bytes.NewReader(file)

		io.Copy(w, binaryFile)
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

		binaryFile := bytes.NewReader(file)

		io.Copy(w, binaryFile)
	case "widgets":

		index, err := strconv.Atoi(r.PathValue("index"))
		if err != nil {
			log.Printf("files handler: failed to read index, %v", err)
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					fmt.Sprintf("files handler: failed to read index, %v", err),
				),
				http.StatusInternalServerError,
			)
			return
		}

		var widget database.Widget
		err = h.DB.Table("widgets").First(&widget, "id = ?", id).Error
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

		binaryFile := bytes.NewReader(widget.Files[index])

		io.Copy(w, binaryFile)
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