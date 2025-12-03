package database

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	// "gorm.io/gorm"
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
		Order: 0,
		Files: files,
		Data:  data,
		Namespace: uuid.Nil.String(),
	}

	result := h.DB.Create(&widget)

	if result.Error != nil {
		log.Printf("failed to create widget: %s", result.Error.Error())
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
		log.Printf("failed to update widget: %s", result.Error.Error())
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
		log.Printf("failed to delete widget: %s", result.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`widget updated`))
	w.Write([]byte("\n\n"))
}

func (h *UserDataHandler) UpdateWidgetOrder(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	widgetID := r.PostFormValue("widget_id")
	order, err := strconv.Atoi(r.PostFormValue("order"))
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

	if order == 0 && posChange == -1 {
		log.Println("cant move first widget up")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	oldOrderA := order
	oldOrderB := order + posChange

	tx := h.DB.Begin()
	if tx.Error != nil {
		log.Printf("failed to update widget order: %s", tx.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	defer func() {
        if tx.Error != nil {
            tx.Rollback() 
        }
    }()
	log.Println("transaction began")
	
	txResult := tx.Exec(`
		UPDATE widgets
		SET "order" = $2
		WHERE id = $1`, widgetID, oldOrderB)
	if txResult.Error != nil {
		log.Printf("failed to update widget order: %s", txResult.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("transaction 1st part")

	txResult = tx.Debug().Exec(`
		UPDATE widgets
		SET "order" = $2
		WHERE "order" = $3 AND id != $1
	`, widgetID, oldOrderA, oldOrderB)
	if txResult.RowsAffected == 0 {
		log.Printf("cant move last widget down")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if txResult.Error != nil {
		log.Printf("failed to update widget order: %s", txResult.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("transaction 2nd part")

	txResult = tx.Commit()
	if txResult.Error != nil {
		log.Printf("failed to update widget: %s", txResult.Error.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("transaction end")

	w.Write([]byte(`order updated`))
	w.Write([]byte("\n\n"))
}