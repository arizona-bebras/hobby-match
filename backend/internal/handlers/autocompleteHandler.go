package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"shumi/internal/database"
	"strings"
)

type AutocompleteHandler struct {
	DB *gorm.DB
}

// Autocompletion represents a single suggested tag with its similarity score.
type Autocompletion struct {
	Id         string  `json:"id"`
	Tag        string  `json:"tag"`
	Similarity float32 `json:"similarity"`
}

// AutocompleteResponse response for autocomplete
type AutocompleteResponse struct {
	Completions []Autocompletion `json:"completions"`
}

type OllamaEmbeddingRequest struct {
	Model     string   `json:"model"`
	Input     []string `json:"input"`
	KeepAlive int      `json:"keep_alive"`
}

type OllamaEmbeddingResponse struct {
	Embeddings [][]float32 `json:"embeddings"`
}

func generateEmbedding(text string) ([]float32, error) {
	text = strings.ToLower(text)
	ollamaApiUrl := os.Getenv("OLLAMA_API_URL")
	ollamaApiKey := os.Getenv("OLLAMA_API_KEY")
	ollamaModel := os.Getenv("OLLAMA_MODEL")

	reqBody := OllamaEmbeddingRequest{
		Model:     ollamaModel,
		KeepAlive: -1,
		Input:     []string{text},
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ollama request: %w", err)
	}

	req, err := http.NewRequest("POST", ollamaApiUrl, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("error making ollama request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+ollamaApiKey)
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making ollama request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ollama request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var ollamaResp OllamaEmbeddingResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("error decoding ollama response: %w", err)
	}

	if len(ollamaResp.Embeddings) > 0 {
		return ollamaResp.Embeddings[0], nil
	}

	return nil, fmt.Errorf("received empty embedding for text '%s'", text)
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

	embedding, err := generateEmbedding(q)
	if err != nil {
		log.Printf("Error generating embedding: %v", err)
		http.Error(w, database.JSONErr(http.StatusInternalServerError, "error generating embedding"), http.StatusInternalServerError)
		return
	}

	var completions []Autocompletion
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
