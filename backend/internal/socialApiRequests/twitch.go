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
	// "strconv"
	"strings"
)

type TwitchUserData struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Followers int    `json:"followers"`
}

func GetTwitchUserInfo(link string) (TwitchUserData, error) {
	ok, err := regexp.MatchString(`^https://www.twitch.tv/[a-zA-Z0-9_.-]*$`, link)
	if err != nil || !ok {
		log.Println("wrong link")
		log.Printf("Error: %v", err)
		return TwitchUserData{}, errors.New("wrong link")
	}
	handle := strings.Split(link[8:], "/")[1]
	log.Println(handle)
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/twitch/channel?key=%s&clientId=%s&handle=%s", os.Getenv("CACHING_ENDPOINT"), os.Getenv("TWITCH_TOKEN"),  os.Getenv("TWITCH_CLIENT_ID"), handle), nil)
	r.Header.Add("authorization", "Bearer SECRET")
	if err != nil {
		log.Println("failed to create request")
		return TwitchUserData{}, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("failed get youtube data")
		return TwitchUserData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body %v", err)
		return TwitchUserData{}, err
	}
	resp.Body.Close()
	log.Println(string(body))

	var bodyJSON map[string]int
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		log.Printf("failed to read body %v", err)
		return TwitchUserData{}, err
	}

	followers := bodyJSON["followers"]
	if err != nil {
		log.Printf("failed to subs count %v", err)
		return TwitchUserData{}, err
	}

	return TwitchUserData{
		Type:  "Twitch",
		Title: handle,
		Followers:  followers,
	}, nil
}
