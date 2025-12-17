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
)

type YoutbeUserData struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Subs  int    `json:"subscribers"`
}

func GetYoutubeUserInfo(link string) (YoutbeUserData, error) {
	ok, err := regexp.MatchString(`^https://www.youtube.com/[a-zA-Z0-9_@.-]*$`, link)
	if err != nil || !ok {
		log.Println("wrong link")
		log.Printf("Error: %v", err)
		return YoutbeUserData{}, errors.New("wrong link")
	}
	var qHandle string
	handle := strings.Split(link[8:], "/")[1]
	if string(handle[0]) == "@" {
		qHandle = fmt.Sprintf("handle=%s", handle)
	} else {
		qHandle = fmt.Sprintf("id=%s", handle)
	}
	log.Println(qHandle)
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/youtube/channel?key=%s&%s", os.Getenv("CACHING_ENDPOINT"), os.Getenv("YOUTUBE_API_KEY"), qHandle), nil)
	r.Header.Add("authorization", "Bearer SECRET")
	if err != nil {
		log.Println("failed to create request")
		return YoutbeUserData{}, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("failed get youtube data")
		return YoutbeUserData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body %v", err)
		return YoutbeUserData{}, err
	}
	resp.Body.Close()
	log.Println(string(body))

	var bodyJSON map[string]map[string]string
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		log.Printf("failed to read body %v", err)
		return YoutbeUserData{}, err
	}

	subsNum, err := strconv.Atoi(bodyJSON["info"]["subscribers"])
	if err != nil {
		log.Printf("failed to subs count %v", err)
		return YoutbeUserData{}, err
	}

	return YoutbeUserData{
		Type:  "YouTube",
		Title: bodyJSON["info"]["title"],
		Subs:  subsNum,
	}, nil
}
