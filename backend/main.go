package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tgAuthPlugin "github.com/iamelevich/pocketbase-plugin-telegram-auth"

	"github.com/go-zoox/fetch"
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
)

func main() {
	app := pocketbase.New()
	godotenv.Load()
	// ignoring for now (useful for docker)
	//   	if err != nil {
	//     	log.Fatal("Error loading .env file")
	//   	}

	token := os.Getenv("BOT_TOKEN")

	var hooksDir string
	app.RootCmd.PersistentFlags().StringVar(
		&hooksDir,
		"hooksDir",
		"pb_hooks",
		"the directory with the JS app hooks",
	)

	var hooksWatch bool
	app.RootCmd.PersistentFlags().BoolVar(
		&hooksWatch,
		"hooksWatch",
		true,
		"auto restart the app on pb_hooks file change",
	)

	var hooksPool int
	app.RootCmd.PersistentFlags().IntVar(
		&hooksPool,
		"hooksPool",
		25,
		"the total prewarm goja.Runtime instances for the JS app hooks execution",
	)

	var migrationsDir string
	app.RootCmd.PersistentFlags().StringVar(
		&migrationsDir,
		"migrationsDir",
		"pb_migrations",
		"the directory with the user defined migrations",
	)

	var automigrate bool
	app.RootCmd.PersistentFlags().BoolVar(
		&automigrate,
		"automigrate",
		true,
		"enable/disable auto migrations",
	)

	app.RootCmd.ParseFlags(os.Args[1:])

	jsvm.MustRegister(app, jsvm.Config{
		MigrationsDir: migrationsDir,
		HooksDir:      hooksDir,
		HooksWatch:    hooksWatch,
		HooksPoolSize: hooksPool,
	})

	// Setup tg auth for users collection
	tgAuthPlugin.MustRegister(app, &tgAuthPlugin.Options{
		BotToken:      token,
		CollectionKey: "users",
	})

	registerHooks(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

type UserData struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Username string `json:"username"`
	Link     string `json:"link"`
}

func registerHooks(app *pocketbase.PocketBase) {
	app.OnRecordEnrich("widgets").BindFunc(func(e *core.RecordEnrichEvent) error {
		collectionName := e.Record.Collection().Name
		log.Printf("[VIEW HOOK] Просмотр записи %s из коллекции %s", e.Record.Id, collectionName)

		if collectionName == "widgets" {
			e.Record.WithCustomData(true)
			fmt.Println(e.Record)
			data := UserData{}
			json.Unmarshal([]byte(e.Record.GetString("data")), &data)
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
			"username": username,
		},
	})
	if err != nil {
		fmt.Println("request error")
		panic(err)
	}

	additionalData, err := response.JSON()

	if err != nil {
		panic(err)
	}
	return additionalData
}
