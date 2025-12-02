package worker

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type WorkerHandler struct {}

func (_ *WorkerHandler) Autocomplete(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("query")
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/interests/query", os.Getenv("WORKER_ENDPOINT")), strings.NewReader(fmt.Sprintf(`{ "query": "%s" }`, q)))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WORKER_SECRET")))
	if err != nil {
		log.Printf("worker/autocomplete: failed to create request: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("worker/autocomplete: failed to process request: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("worker/autocomplete: failed to read response body: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf(`{ "query": "%s", "response": %s}`, q, string(body))))
	w.Write([]byte("\n\n"))
}

func (h WorkerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	switch path {
	case "/api/worker/autocomplete":
		switch r.Method {
		case http.MethodGet:
			h.Autocomplete(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "unknow endpoint", http.StatusNotFound)
	}
}