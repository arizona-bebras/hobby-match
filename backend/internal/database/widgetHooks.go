package database

import (
	"encoding/json"
	"errors"
	"log"
	"gorm.io/gorm"

	"github.com/lib/pq"

	"shumi/internal/socialapirequests"
)

type SurveyAdditionalData struct {
	Type   string `json:"type"`
	Stats  []int  `json:"stats"`
	MyVote *int   `json:"myVote"`
}

type SurveyOptions struct {
	Options []map[string]string `json:"options"`
}

type PhotoAdditionalData struct {
	Type string        `json:"type"`
	Urls pq.ByteaArray `json:"urls"`
}

type SocialData struct {
	Platform string `json:"platform"`
	Link     string `json:"link"`
}

type GameData struct {
	Type        string `json:"type"`
	GameId      string  `json:"gameId"`
	AccountLink string `json:"accountLink"`
}

type WidgetDataType struct {
	Type string `json:"type"`
}

func (w *Widget) BeforeCreate(db *gorm.DB) error {
	result := db.Table("widgets").
		Where(`"user" = ?`, w.User).
		Update("order", gorm.Expr(`"order" + ?`, 1))

	if result.Error != nil {
		log.Printf("failed to increase widget order: %s", result.Error.Error())
		return errors.New("failed to increase widget order")
	}
	return nil
}

func (w *Widget) AfterFind(db *gorm.DB) error {
	var widgetType WidgetDataType
	log.Println(w.Data)
	err := json.Unmarshal([]byte(w.Data), &widgetType)
	if err != nil {
		log.Printf("failed to get widget type %v", err)
		return errors.New("failed to get widget type")
	}

	switch widgetType.Type {
	case "survey":
		var surveyOptions SurveyOptions
		err := json.Unmarshal([]byte(w.Data), &surveyOptions)
		if err != nil {
			log.Printf("failed to get widget type %v", err)
			return errors.New("failed to get widget type")
		}

		type Result struct {
			Option int `gorm:"column:option"`
			Count  int `gorm:"column:count"`
			MyVote int `gorm:"column:my_vote"`
		}

		additionalData := SurveyAdditionalData{
			Type:   "survey",
			Stats:  make([]int, len(surveyOptions.Options)),
			MyVote: nil,
		}

		var results []Result
		err = db.Raw(`
			WITH vote_stats AS (
				SELECT 
					option, 
					COUNT(*) as count,
					(SELECT option FROM votes WHERE survey = $1 AND "user" = $2) as my_vote
				FROM votes
				WHERE survey = $1
				GROUP BY option
				ORDER BY option ASC
			)
			SELECT * FROM vote_stats
		`, w.Id, w.User).Scan(&results).Error

		if err != nil {
			log.Printf("failed to get votes %v", err)
			return err
		}

		for _, result := range results {
			additionalData.Stats[result.Option] = result.Count
		}

		var myVote int
		if len(results) > 0 {
			myVote = results[0].MyVote
			additionalData.MyVote = &myVote
		}

		additionalDataJSON, err := json.Marshal(additionalData)
		if err != nil {
			log.Printf("failed to get votes %v", err)
			return err
		}
		w.AdditionalData = string(additionalDataJSON)
	case "social_media":
		var platform SocialData
		err := json.Unmarshal([]byte(w.Data), &platform)
		if err != nil {
			log.Printf("failed to get platform type %v", err)
			return errors.New("failed to get widget type")
		}
		switch platform.Platform {
			
		case "Steam":
			additionalData, err := socialapirequests.GetSteamUserInfo(platform.Link)
			if err != nil {
				log.Println(err)
				return nil
			}
			additionalDataJSON, err := json.Marshal(additionalData)
			if err != nil {
				log.Printf("failed to get votes %v", err)
				return err
			}
			w.AdditionalData = string(additionalDataJSON)

		case "YouTube":
			additionalData, err := socialapirequests.GetYoutubeUserInfo(platform.Link)
			if err != nil {
				log.Println(err)
				return nil
			}
			additionalDataJSON, err := json.Marshal(additionalData)
			if err != nil {
				log.Printf("failed to get votes %v", err)
				return err
			}
			w.AdditionalData = string(additionalDataJSON)
		case "Twitch":
			additionalData, err := socialapirequests.GetTwitchUserInfo(platform.Link)
			if err != nil {
				log.Println(err)
				return nil
			}
			additionalDataJSON, err := json.Marshal(additionalData)
			if err != nil {
				log.Printf("failed to get votes %v", err)
				return err
			}
			w.AdditionalData = string(additionalDataJSON)
		}
	case "steam_game":
		var gameData GameData
		err := json.Unmarshal([]byte(w.Data), &gameData)
		if err != nil {
			log.Println(err)
			return nil
		}

		additionalData, err := socialapirequests.GetGameInfo(gameData.AccountLink, gameData.GameId)
		if err != nil {
			log.Println(err)
			return nil
		}
		additionalDataJSON, err := json.Marshal(additionalData)
		if err != nil {
			log.Printf("failed to get votes %v", err)
			return err
		}
		w.AdditionalData = string(additionalDataJSON)
	}

	return nil
}

func (w *Widget) BeforeDelete(db *gorm.DB) error {
	log.Printf("widget: %v", w)
	result := db.Debug().Table("widgets").
		Where(`"user" = ? AND "order" > ?`, w.User, w.Order).
		Update("order", gorm.Expr(`"order" - ?`, 1))

	if result.Error != nil {
		log.Printf("failed to normalize widget order: %s", result.Error.Error())
		return errors.New("failed to normalize widget order")
	}
	return nil
}