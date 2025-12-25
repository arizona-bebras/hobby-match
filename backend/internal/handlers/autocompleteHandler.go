package handlers

import (
	"encoding/json"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
	"log"
	"net/http"
	"shumi/internal/database"
	"shumi/internal/embeddings"
)

type AutocompleteHandler struct {
	DB *gorm.DB
}

// InterestSimilarity represents a single suggested tag with its similarity score.
type InterestSimilarity struct {
	Id         string  `json:"id"`
	Tag        string  `json:"tag"`
	Similarity float32 `json:"similarity"`
}

// AutocompleteResponse response for autocomplete
type AutocompleteResponse struct {
	Completions []InterestSimilarity `json:"completions"`
}

// Autocomplete
// @Summary Предложенные интересы
// @Description Возвращает список интересов, семантически похожих на запрос 'q'.
// @Produce json
// @Param q query string false "Строка для поиска и автодополнения интересов"
// @Success 200 {object} handlers.AutocompleteResponse "Успешное получение списка интересов"
// @Failure 401 {object} database.Error "Ошибка авторизации"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/autocomplete [get]
func (h *AutocompleteHandler) Autocomplete(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}

	embedding, err := embeddings.GenerateEmbedding(q)
	if err != nil {
		log.Printf("Error generating embedding: %v", err)
		http.Error(w, database.JSONErr(http.StatusInternalServerError, "error generating embedding"), http.StatusInternalServerError)
		return
	}

	var completions []InterestSimilarity
	result := h.DB.Raw(
		"SELECT id, tag, 1 - (embedding <-> ?) AS similarity FROM interests ORDER BY embedding <-> ? LIMIT 10",
		pgvector.NewVector(embedding),
		pgvector.NewVector(embedding)).Scan(&completions)
	if result.Error != nil {
		log.Printf("Database error: %v", result.Error)
		http.Error(w, database.JSONErr(http.StatusInternalServerError, "database error"), http.StatusInternalServerError)
		return
	}

	response := AutocompleteResponse{Completions: completions}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, database.JSONErr(http.StatusInternalServerError, "error marshalling response"), http.StatusInternalServerError)
		return
	}

	w.Write(responseBytes)
}

func (h AutocompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	switch path {
	case "/api/autocomplete":
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
