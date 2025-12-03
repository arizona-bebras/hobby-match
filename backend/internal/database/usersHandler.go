package database

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"

	// "strconv"

	"gorm.io/gorm"
)

type contextKey string

const AuthContextKey = contextKey("TgID")

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

func (h *UserDataHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	tgID, ok := r.Context().Value(AuthContextKey).(string)
	if !ok {
        log.Printf("Authentication error: TgID is missing or not int64")
        http.Error(w, "Authentication required", http.StatusUnauthorized)
        return
    }
	var user User
	result := h.DB.Table("users").First(&user, "tg_id = ?", tgID)
	if result.Error != nil {
		log.Printf("failed to get user: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var widgets []Widget
	result = h.DB.Table("widgets").Find(&widgets, "widgets.user = ?", tgID)
	if result.Error != nil {
		log.Printf("failed to get widgets: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	user.Widgets = widgets

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Printf("failed to serialize user: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(userJSON)
	w.Write([]byte("\n\n"))
}

func (h *UserDataHandler) UpdateMyProfileInfo(w http.ResponseWriter, r *http.Request) {
	tgID := r.Context().Value(AuthContextKey)

	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Printf("failed read data: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	log.Println(string(body))

	var user User
	log.Printf("GOT DATA %v", user)
	err = json.Unmarshal(body, &user)
	user.TgID = tgID.(string)

	log.Printf("%v", user)
	log.Printf("%s", getStructFieldNames(user))

	var result *gorm.DB
	if len(user.Interests) == 0 {
		result = h.DB.Model(&user).Select("name", "location", "gender", "birth_date", "info").Updates(&user)
	} else {
		result = h.DB.Model(&user).Select("interests").Updates(&user)
	}

	if result.Error != nil {
		log.Printf("failed to update user: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`user updated`))
	w.Write([]byte("\n\n"))
}

func (h *UserDataHandler) UpdateMyProfilePhoto(w http.ResponseWriter, r *http.Request) {
	tgID := r.Context().Value(AuthContextKey).(string)

	photoFile, _, err := r.FormFile("user_photo")
	if err != nil {
		log.Printf("failed to get photo file: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	photoBytes, err := io.ReadAll(photoFile)
	if err != nil {
		log.Printf("failed to read photo bytes: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	photoFile.Close()

	result := h.DB.Table("users").Where("tg_id = ?", tgID).Select("photo").Updates(&map[string]interface{}{
		"photo": photoBytes,
	})

	if result.Error != nil {
		log.Printf("failed to update user: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}


	w.Write([]byte(`user updated`))
	w.Write([]byte("\n\n"))
}

func (h UserDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	switch path {
	case "/api/me/":
		switch r.Method {
		case http.MethodPost:
			h.UpdateMyProfileInfo(w, r)
		case http.MethodGet:
			h.GetMe(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	case "/api/me/photo":
		switch r.Method {
		case http.MethodPost:
			h.UpdateMyProfilePhoto(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

	case "/api/me/widgets":
		switch r.Method {
		case http.MethodPost:
			h.CreateWidget(w, r)
		case http.MethodPut:
			h.UpdateWidget(w, r)
		case http.MethodDelete:
			h.DelteWidget(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

	case "/api/me/widgets/order":
		switch r.Method {
		case http.MethodPut:
			h.UpdateWidgetOrder(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

