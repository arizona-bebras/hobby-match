package main

import (
	"log"
	"os"

	tgAuthPlugin "github.com/iamelevich/pocketbase-plugin-telegram-auth"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
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

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
