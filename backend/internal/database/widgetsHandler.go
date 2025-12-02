package database

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

const FORM_SIZE_LIMIT int64 = 50000000

func (h *UserDataHandler) CreateWidget(w http.ResponseWriter, r *http.Request) {
	TgID := r.Context().Value(AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	data := r.PostFormValue("data")
	log.Println(data)

	files := [][]byte{}

	var file []byte
	i := 0
	for {
		file = []byte(r.PostFormValue(fmt.Sprintf(`file_%d`, i)))
		if len(file) == 0 {
			break
		}
		files = append(files, file)
	}

	widget := Widget{
		Id:    uuid.NewString(),
		User:  TgID,
		Order: -1,
		Files: files,
		Data:  data,
		Namespace: uuid.Nil.String(),
	}

	result := h.DB.Create(&widget)

	if result.Error != nil {
		log.Printf("failed to update user: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`widget created`))
	w.Write([]byte("\n\n"))
}

func (h *UserDataHandler) UpdateWidget(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	data := r.PostFormValue("data")
	widgetID := r.PostFormValue("widget_id")
	log.Println(data)

	files := [][]byte{}

	var file []byte
	i := 0
	for {
		file = []byte(r.PostFormValue(fmt.Sprintf(`file_%d`, i)))
		if len(file) == 0 {
			break
		}
		files = append(files, file)
	}


	widget := Widget{
		Id:    widgetID,
		Files: files,
		Data:  data,
	}

	result := h.DB.Model(&widget).Select("data").Updates(&widget)
	if result.Error != nil {
		log.Printf("failed to update user: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`widget updated`))
	w.Write([]byte("\n\n"))
}
func (h *UserDataHandler) DelteWidget(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("widget_id")

	result := h.DB.Delete(&Widget{}, "id = ?",id)
	if result.Error != nil {
		log.Printf("failed to update user: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`widget updated`))
	w.Write([]byte("\n\n"))
}

func (h *UserDataHandler) UpdateWidgetOrder(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	widgetID := r.PostFormValue("widget_id")
	order, err := strconv.Atoi(r.PostFormValue("data"))
	if err != nil {
		log.Printf("wrong order: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	posChange, err := strconv.Atoi(r.PostFormValue("pos_change"))
	if err != nil {
		log.Printf("wrong order: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	result := h.DB.Exec(`
		UPDATE widgets
		SET order = $1
		WHERE id = $2

		UPDATE widgets
		SET order = $3
		WHERE order = $1 AND id != $2
	`, order + posChange, widgetID, order)
	if result.Error != nil {
		log.Printf("failed to update user: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`order updated`))
	w.Write([]byte("\n\n"))
}