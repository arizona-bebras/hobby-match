package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	"shumi/internal/socialapirequests"
	"shumi/internal/database"
)

type GamesHandler struct {
	DB *gorm.DB
}

// GetGames
// @Summary Игры пользователя из его steam профиля для виджета
// @Produce json
// @Param link query string true "Ссылка на профиль"
// @Success 200 {string} string "JSON, содержащий список игр"
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/games [get]
func (h *GamesHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	id, err := socialapirequests.ResolveSteamLink(r.URL.Query().Get("link"))
	if err != nil {
		log.Printf("game handler: failed to get id %v", err)
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				fmt.Sprintf("game handler: failed to get id %v", err),
			), 
			http.StatusInternalServerError,
		)
		return
	}

	gamesReq, err := http.NewRequest("GET", fmt.Sprintf("%s/steam/games?key=%s&id=%s", os.Getenv("CACHING_ENDPOINT"), os.Getenv("STEAM_API_KEY"), id), nil)
	gamesReq.Header.Add("authorization", "Bearer SECRET")
	if err != nil {
		log.Println("game handler: failed to create request")
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				"game handler: failed to create request",
			), 
			http.StatusInternalServerError,
		)
		return
	}
	resp, err := http.DefaultClient.Do(gamesReq)
	if err != nil {
		log.Println("game handler: failed get games")
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				"game handler: failed get games",
			), 
			http.StatusInternalServerError,
		)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("game handler: failed to read body")
		http.Error(
			w, 
			database.JSONErr(
				http.StatusInternalServerError, 
				"game handler: failed get games",
			), 
			http.StatusInternalServerError,
		)
		return
	}
	resp.Body.Close()

	w.Write([]byte(fmt.Sprintf("%s", string(body))))
	w.Write([]byte("\n\n"))
}

func (h GamesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("game handler: method not allowed")
		http.Error(
				w, 
				database.JSONErr(
					http.StatusMethodNotAllowed, 
					"game handler: method not allowed",
				), 
				http.StatusMethodNotAllowed,
			)
		return
	}
	h.GetGames(w, r)
}