package database

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

type UserVector struct {
	Id        string   `json:"id"`
	Interests []string `json:"interest_ids"`
	Text      string   `json:"text"`
}

func upsertUser(user User) error {
	userVector := UserVector {
		Id: user.TgID,
		Interests: user.Interests,
		Text: user.Info,
	}

	userVectorJSON, err := json.Marshal(userVector)
	if err != nil {
		log.Printf("worker/upsert: failed to create user vector: %v", err)
		return err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/feed/upsert", os.Getenv("WORKER_ENDPOINT")), strings.NewReader(string(userVectorJSON)))
	if err != nil {
		log.Printf("worker/upsert: failed to create request: %v", err)
		return err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WORKER_SECRET")))

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("worker/upsert: failed to process request: %v", err)
		return err
	}

	if resp.Status != string(http.StatusOK) {
		log.Println("worker/upsert: failed create vector")
		return errors.New("failed create vector")
	}

	return nil
}

func (u *User) AfterCreate(_ *gorm.DB) error {
	err := upsertUser(*u)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (u *User) AfterUpdate(_ *gorm.DB) error {
	err := upsertUser(*u)
	if err != nil {
		log.Println(err)
	}
	return nil
}