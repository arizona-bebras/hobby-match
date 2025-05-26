package main

import (
	"encoding/json"
	"log"
	"os"
	"fmt"

	tgAuthPlugin "github.com/iamelevich/pocketbase-plugin-telegram-auth"

	"github.com/go-zoox/fetch"
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)



func main() {
	app := pocketbase.New()
	godotenv.Load()
  // ignoring for now (useful for docker)
  //   	if err != nil {
  //     	log.Fatal("Error loading .env file")
  //   	}
	if _, err := os.Stat("./pb_hooks/main.pb.js"); err == nil {
        log.Println("JS hooks file found")
    } else {
        log.Println("JS hooks file NOT found")
    }
	registerHooks(app)

	token := os.Getenv("BOT_TOKEN")

	// Setup tg auth for users collection
	tgAuthPlugin.MustRegister(app, &tgAuthPlugin.Options{
		BotToken: token,
		CollectionKey: "users",
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

type UserData struct {
	Type     string `json: "type"`
	Platform string `json: "platform"`
	Username string `json: "username"`
	Link     string `json: "link"`
}

func registerHooks(app *pocketbase.PocketBase) {
	app.OnRecordViewRequest().Add(func(e *core.RecordViewEvent) error {
		collectionName := e.Record.Collection().Name
		log.Printf("[VIEW HOOK] Просмотр записи %s из коллекции %s", e.Record.Id, collectionName)

		if collectionName == "widgets" {
			e.Record.WithUnknownData(true);
			fmt.Println(e.Record)
			data := UserData{}
			json.Unmarshal([]byte(e.Record.GetString("data")), &data);
			username := data.Username
			e.Record.Set("additional_data", youtubeRequest(username))
			fmt.Println(youtubeRequest(username))
			fmt.Println(e.Record)
		}

		return nil
	})
}

func youtubeRequest(username string) string {
	response, err := fetch.Post("http://localhost:5173/api/youtube", &fetch.Config{
		Body: map[string]interface{}{
			"username":     username,
		},
	})
	if err != nil {
		fmt.Println("request error")
		panic(err)
	}

	additionalData, err := response.JSON();

	if err != nil {
		panic(err)
	}
  return additionalData;
}
