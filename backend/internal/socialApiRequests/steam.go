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
	"strings"
)

type SteamUserData struct {
	Type     string `json:"type"`
	Level    int    `json:"level"`
	Username string `json:"username"`
}

func resolveSteamLink(link string) (string, error) {
	if ok, err := regexp.MatchString(
		`^(?:https?:\/\/)?steamcommunity\.com\/(profiles|id)\/([a-zA-Z0-9_.-]+/?)$`,
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
