package database

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	// "strings"

	"gorm.io/gorm"
)

type contextKey string

const AuthContextKey = contextKey("TgID")

type UserDataHandler struct {
	DB *gorm.DB
}

func getWidgetsByUserID(DB *gorm.DB, userID int64) ([]Widget, error) {
	var widgets []Widget
	result := DB.Table("widgets").Find(&widgets, "user = ?", userID)
	log.Printf("%v", widgets)
	if result.Error != nil {
		return []Widget{}, result.Error
	}
	return widgets, nil
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
	tgID := r.Context().Value(AuthContextKey)
	var user User
	result := h.DB.Table("users").First(&user, "tg_id = ?", tgID)
	if result.Error != nil {
		log.Printf("failed to get user: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	widgets, err := getWidgetsByUserID(h.DB, user.TgID)
	if err != nil {
		log.Printf("failed to get widgets: %s", err.Error())
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
	if err != nil {
		log.Printf("failed read data: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	log.Println(string(body))

	var payload User
	log.Printf("GOT DATA %v", payload)
	err = json.Unmarshal(body, &payload)
	payload.TgID = tgID.(int64)

	log.Printf("%v", payload)
	log.Printf("%s", getStructFieldNames(payload))

	var result *gorm.DB
	if len(payload.Interests) == 0 {
		result = h.DB.Model(&payload).Select("name", "location", "gender", "birth_date", "info").Updates(&payload)
	} else {
		result = h.DB.Model(&payload).Select("interests").Updates(&payload)
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
	tgID, ok := r.Context().Value(AuthContextKey).(int64)
	if !ok {
		log.Printf("failed to get tg id")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

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
	}
}
