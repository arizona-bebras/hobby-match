package handlers

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"reflect"
	"shumi/internal/database"
)

type UserDataHandler struct {
	DB *gorm.DB
}

func getStructFieldNames(s interface{}) []string {
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		log.Printf("Warning: Input is not a struct, but a %s\n", t.Kind())
		return nil
	}

	fieldNames := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.IsExported() {
			fieldNames = append(fieldNames, field.Tag.Get("db"))
		}
	}

	return fieldNames
}

// GetMe
// @Summary Профиль пользователя
// @Produce json
// @Success 200 {object} database.User
// @Failure 401 {object} database.Error "Пользователь не авторизован"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/me [get]
func (h *UserDataHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	tgID, ok := r.Context().Value(database.AuthContextKey).(string)
	if !ok {
		log.Printf("Authentication error, TgID is missing or not int64")
		http.Error(
			w,
			database.JSONErr(
				http.StatusUnauthorized,
				"Authentication error, TgID is missing or not int64",
			),
			http.StatusUnauthorized,
		)
		return
	}
	var user database.User
	result := h.DB.Model(&user).First(&user, "id = ?", tgID)
	if result.Error != nil {
		log.Printf("users handler: failed to get user: %s", result.Error.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to get user: %s", result.Error.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}

	var widgets []database.Widget
	result = h.DB.Table("widgets").Find(&widgets, "widgets.user = ?", tgID)
	if result.Error != nil {
		log.Printf("users handler: failed to get widgets: %s", result.Error.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to get widgets: %s", result.Error.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}
	user.Widgets = widgets

	userJSON, err := json.Marshal(user)
	log.Println(string(userJSON))
	if err != nil {
		log.Printf("users handler: failed to serialize user, %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to serialize user, %s", err.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}

	w.Write(userJSON)
}

// GetMyNamespaces
// @Summary Неймспейсы пользователя
// @Produce json
// @Success 200 {object} handlers.UserNamespaces
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/me/namespaces [get]
func (h *UserDataHandler) GetMyNamespaces(w http.ResponseWriter, r *http.Request) {
	tgId := r.Context().Value(database.AuthContextKey)

	var namespaces []database.Namespace
	err := h.DB.
		Table("user_namespace").
		Select("namespaces.id, namespaces.title, namespaces.members_count").
		Joins("JOIN namespaces ON namespaces.id = user_namespace.namespace_id").
		Find(&namespaces, "user_id = ?", tgId).Error
	if err != nil {
		log.Printf("users handler: failed to get user namespaces, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to get user namespaces, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	userNamespaces := UserNamespaces{
		Namespaces: namespaces,
	}

	namespacesJSON, err := json.Marshal(userNamespaces)
	if err != nil {
		log.Printf("users handler: failed to marshal user namespaces, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to marshal user namespaces, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	w.Write(namespacesJSON)
}

// UpdateMyProfileInfo
// @Summary Обновить информацию профиля пользователя
// @Accept json
// @Param miniapp_name body string false "Имя пользователя"
// @Param location body string false "Локация пользователя"
// @Param gender body string false "Пол пользователя"
// @Param birth_date body string false "Дата рождения пользователя"
// @Param info body string false "Информация о пользователе"
// @Param interests body array false "Интересы о пользователе"
// @Param personality_test body array false "Личностный тест (5 слайдеров)"
// @Success 200 {object} nil "Пользователь успешно обновлен"
// @Failure 401 {object} database.Error "Пользователь не авторизован"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/me [patch]
func (h *UserDataHandler) UpdateMyProfileInfo(w http.ResponseWriter, r *http.Request) {
	tgID := r.Context().Value(database.AuthContextKey)

	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Printf("users handler: failed read data, %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed read data, %s", err.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}

	log.Println(string(body))

	var user database.User
	log.Printf("GOT DATA %v", user)
	err = json.Unmarshal(body, &user)
	user.Id = tgID.(string)

	log.Printf("%v", user)
	log.Printf("%s", getStructFieldNames(user))

	var result *gorm.DB

	result = h.DB.Model(&user).Updates(&user)

	if result.Error != nil {
		log.Printf("users handler: failed to update user, %s", result.Error.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to update user, %s", result.Error.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}
}

// UpdateMyProfilePhoto
// @Summary Обновить фото профиля пользователя
// @Accept json
// @Param user_photo body []byte false "Фото пользователя"
// @Success 200 {object} nil "Пользователь успешно обновлен"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/me/photo [patch]
func (h *UserDataHandler) UpdateMyProfilePhoto(w http.ResponseWriter, r *http.Request) {
	tgID := r.Context().Value(database.AuthContextKey).(string)

	photoFile, _, err := r.FormFile("user_photo")
	if err != nil {
		log.Printf("users handler: failed to get photo file, %s", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to get photo file, %s", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	photoBytes, err := io.ReadAll(photoFile)
	if err != nil {
		log.Printf("users handler: failed to read photo bytes, %s", err.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to read photo bytes, %s", err.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}

	photoFile.Close()

	result := h.DB.Table("users").Where("id = ?", tgID).Select("photo").Updates(&map[string]interface{}{
		"photo": photoBytes,
	})

	if result.Error != nil {
		log.Printf("users handler: failed to update user, %s", result.Error.Error())
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("users handler: failed to update user, %s", result.Error.Error()),
			),
			http.StatusInternalServerError,
		)
		return
	}
}

func (h UserDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	switch path {
	case "/api/me":
		switch r.Method {
		case http.MethodPatch:
			h.UpdateMyProfileInfo(w, r)
		case http.MethodGet:
			h.GetMe(w, r)
		default:
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					"users handler: method not allowed,",
				),
				http.StatusInternalServerError,
			)
			return
		}

	case "/api/me/photo":
		switch r.Method {
		case http.MethodPatch:
			h.UpdateMyProfilePhoto(w, r)
		default:
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					"users handler: method not allowed,",
				),
				http.StatusInternalServerError,
			)
			return
		}

	case "/api/me/widgets":
		switch r.Method {
		case http.MethodPost:
			h.CreateWidget(w, r)
		case http.MethodPatch:
			h.UpdateWidget(w, r)
		case http.MethodDelete:
			h.DeleteWidget(w, r)
		default:
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					"widgets handler: method not allowed,",
				),
				http.StatusInternalServerError,
			)
			return
		}

	case "/api/me/namespaces":
		switch r.Method {
		case http.MethodGet:
			h.GetMyNamespaces(w, r)
		default:
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					"widgets handler: method not allowed,",
				),
				http.StatusInternalServerError,
			)
			return
		}

	case "/api/me/widgets/order":
		switch r.Method {
		case http.MethodPatch:
			h.UpdateWidgetOrder(w, r)
		default:
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					"widgets handler: method not allowed,",
				),
				http.StatusInternalServerError,
			)
			return
		}

	case "/api/me/widgets/photo":
		switch r.Method {
		case http.MethodDelete:
			h.DeleteWidgetPhoto(w, r)
		default:
			http.Error(
				w,
				database.JSONErr(
					http.StatusInternalServerError,
					"widgets handler: method not allowed,",
				),
				http.StatusInternalServerError,
			)
			return
		}
	}
}
