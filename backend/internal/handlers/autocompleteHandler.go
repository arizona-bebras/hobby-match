package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"gorm.io/gorm"
	"shumi/internal/database"
)

type AutocompleteHandler struct {
	DB *gorm.DB
}

// Autocomplete
// @Summary Предложенные пользователю интересы
// @Produce json
// @Param q query string false "Ввод пользователя"
// @Success 200 {object} handlers.AutocompleteResponse
// @Failure 401 {object} database.Error "Ошибка авторизации"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/worker/autocomplete [get]
func (_ *AutocompleteHandler) Autocomplete(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("query")
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/interests/query", os.Getenv("WORKER_ENDPOINT")), strings.NewReader(fmt.Sprintf(`{ "query": "%s" }`, q)))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WORKER_SECRET")))
	if err != nil {
		log.Printf("autocomplete handler: failed to create request, %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("autocomplete handler: failed to create request, %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("autocomplete handler: failed to process request, %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("autocomplete handler: failed to process request, %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("autocomplete handler: failed to read response body, %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("autocomplete handler: failed to read response body, %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	w.Write([]byte(fmt.Sprintf(`{ "query": "%s", "response": %s}`, q, string(body))))
}

func (h AutocompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	switch path {
	case "/api/worker/autocomplete":
		switch r.Method {
		case http.MethodGet:
			h.Autocomplete(w, r)
		default:
			http.Error(
				w, 
				database.JSONErr(
					http.StatusMethodNotAllowed, 
					"autocomplete handler: method not allowed",
				), 
				http.StatusMethodNotAllowed,
			)
		}
	default:
		http.Error(
			w, 
			database.JSONErr(
				http.StatusNotFound, 
				"autocomplete handler: unknown endpoint",
			), 
			http.StatusNotFound,
		)
	}
}