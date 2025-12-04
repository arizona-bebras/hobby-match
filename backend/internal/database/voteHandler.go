package database

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"log"
)

type VoteHandler struct {
	DB *gorm.DB
}

func (h *VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {
	tgID := r.Context().Value(AuthContextKey).(string)

	r.ParseMultipartForm(FORM_SIZE_LIMIT)
	survey := r.PostFormValue("survey")
	option, err := strconv.Atoi(r.PostFormValue("option"))
	if err != nil {
		log.Println("failes to vote")
		http.Error(w, "interbal server error", http.StatusInternalServerError)
		return
	}

	vote := Vote{
		User: tgID,
		Survey: survey,
		Option: option,
	}

	result := h.DB.Create(&vote)
	if result.Error != nil {
		log.Println("failed to vote")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`voted`))
	w.Write([]byte("\n\n"))
}

func (h VoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/api/vote/":
		switch r.Method {
		case http.MethodPost:
			h.Vote(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "endpoint not found", http.StatusNotFound)
	}
}