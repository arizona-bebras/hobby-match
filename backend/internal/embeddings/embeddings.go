package embeddings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type OllamaEmbeddingRequest struct {
	Model     string   `json:"model"`
	Input     []string `json:"input"`
	KeepAlive int      `json:"keep_alive"`
}

type OllamaEmbeddingResponse struct {
	Embeddings [][]float32 `json:"embeddings"`
}

func GenerateEmbedding(text string) ([]float32, error) {
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
