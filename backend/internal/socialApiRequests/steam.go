package socialapirequests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type GamesHandler struct {
	DB *gorm.DB
}

type SteamUserData struct {
	Type     string `json:"type"`
	Level    int    `json:"level"`
	Username string `json:"username"`
}

type GameData struct {
	Type        string `json:"type"`
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	HoursPlayed int    `json:"hours_played"`
}

func resolveSteamLink(link string) (string, error) {
	if ok, err := regexp.MatchString(
		`^(?:https?:\/\/)?steamcommunity\.com\/(profiles|id)\/([a-zA-Z0-9_.-]+\/?)$`,
		link); err != nil || !ok {
		log.Println("wrong link")
		log.Printf("Error: %v", err)
		return "", errors.New("wrong link")
	}
	parts := strings.Split(link[8:], "/")
	profileType := parts[1]
	if profileType == "profiles" {
		return parts[2], nil
	}

	r, err := http.NewRequest("GET", fmt.Sprintf("%s/steam/vanity?key=%s&vanityurl=%s", os.Getenv("CACHING_ENDPOINT"), os.Getenv("STEAM_API_KEY"), parts[2]), nil)
	if err != nil {
		log.Println("failed to create request")
		return "", err
	}
	r.Header.Add("authorization", "Bearer SECRET")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("failed to resolve url")
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body")
		return "", err
	}
	resp.Body.Close()
	log.Println(string(body))
	var bodyJSON map[string]interface{}
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		log.Printf("failed to read body")
		return "", err
	}
	id, ok := bodyJSON["id"].(string)
	if !ok {
		log.Printf("failed to read body")
		return "", err
	}
	return id, nil
}

func GetSteamUserInfo(link string) (SteamUserData, error) {
	id, err := resolveSteamLink(link)
	if err != nil {
		return SteamUserData{}, err
	}

	r, err := http.NewRequest("GET", fmt.Sprintf("%s/steam/level?key=%s&id=%s", os.Getenv("CACHING_ENDPOINT"), os.Getenv("STEAM_API_KEY"), id), nil)
	r.Header.Add("authorization", "Bearer SECRET")
	if err != nil {
		log.Println("failed to create request")
		return SteamUserData{}, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("failed get lvl")
		return SteamUserData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body")
		return SteamUserData{}, err
	}
	resp.Body.Close()
	log.Println(string(body))
	var levelBodyJSON map[string]int
	err = json.Unmarshal(body, &levelBodyJSON)
	if err != nil {
		log.Printf("failed to read body %v", err)
		return SteamUserData{}, err
	}
	level, ok := levelBodyJSON["level"]
	if !ok {
		log.Printf("failed to read body %v", err)
		return SteamUserData{}, err
	}

	r, err = http.NewRequest("GET", fmt.Sprintf("%s/steam/username?key=%s&id=%s", os.Getenv("CACHING_ENDPOINT"), os.Getenv("STEAM_API_KEY"), id), nil)
	r.Header.Add("authorization", "Bearer SECRET")
	if err != nil {
		log.Println("failed to create request")
		return SteamUserData{}, err
	}
	resp, err = http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("failed get lvl")
		return SteamUserData{}, err
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body")
		return SteamUserData{}, err
	}
	resp.Body.Close()
	log.Println(string(body))
	var usernameBodyJSON map[string]string
	err = json.Unmarshal(body, &usernameBodyJSON)
	if err != nil {
		log.Printf("failed to read body %v", err)
		return SteamUserData{}, err
	}
	username, ok := usernameBodyJSON["username"]
	if !ok {
		log.Printf("failed to read body")
		return SteamUserData{}, err
	}

	return SteamUserData{
		Type:     "Steam",
		Level:    level,
		Username: username,
	}, nil
}

func GetGameInfo(link string, appid string) (GameData, error) {
	id, err := resolveSteamLink(link)
	if err != nil {
		return GameData{}, err
	}

	appidNum, err := strconv.Atoi(appid)
	if err != nil {
		return GameData{}, err
	}

	r, err := http.NewRequest("GET", fmt.Sprintf("%s/steam/game?key=%s&id=%s&appid=%d", os.Getenv("CACHING_ENDPOINT"), os.Getenv("STEAM_API_KEY"), id, appidNum), nil)
	r.Header.Add("authorization", "Bearer SECRET")
	if err != nil {
		log.Println("failed to create request")
		return GameData{}, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("failed get lvl")
		return GameData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body")
		return GameData{}, err
	}
	resp.Body.Close()
	log.Println(string(body))
	var bodyJSON map[string]map[string]interface{}
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		log.Printf("failed to read body %v", err)
		return GameData{}, err
	}

	return GameData{
		Type: "steam_game",
		Icon: fmt.Sprintf(`https://media.steampowered.com/steamcommunity/public/images/apps/%s/%s.jpg`, appid, bodyJSON["game"]["icon"].(string)),
		Title: bodyJSON["game"]["game_name"].(string),
		HoursPlayed: int(bodyJSON["game"]["hours"].(float64)),
	}, nil
}

func (h *GamesHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	id, err := resolveSteamLink(r.URL.Query().Get("link"))
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
