package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	"shumi/internal/socialapirequests"
)

type GamesHandler struct {
	DB *gorm.DB
}

// GetGames
// @Summary Игры пользователя из его steam профиля для виджета
// @Produce json
// @Param link string query true Ссылка на профиль
// @Success 200 {string} "JSON, содержащий список игр"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/games [get]

func (h *GamesHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	id, err := socialapirequests.ResolveSteamLink(r.URL.Query().Get("link"))
	if err != nil {
		log.Printf("failed to get id %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	gamesReq, err := http.NewRequest("GET", fmt.Sprintf("%s/steam/games?key=%s&id=%s", os.Getenv("CACHING_ENDPOINT"), os.Getenv("STEAM_API_KEY"), id), nil)
	gamesReq.Header.Add("authorization", "Bearer SECRET")
	if err != nil {
		log.Println("failed to create request")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(gamesReq)
	if err != nil {
		log.Printf("failed get games")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	resp.Body.Close()

	w.Write([]byte(fmt.Sprintf(`%s`, string(body))))
	w.Write([]byte("\n\n"))
}

func (h GamesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.GetGames(w, r)
}