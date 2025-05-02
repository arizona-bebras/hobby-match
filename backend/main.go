package main

import (
	"log"
	"os"

	tgAuthPlugin "github.com/iamelevich/pocketbase-plugin-telegram-auth"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
)



func main() {
	app := pocketbase.New()
	godotenv.Load()
  // ignoring for now (useful for docker)
  //   	if err != nil {
  //     	log.Fatal("Error loading .env file")
  //   	}
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
