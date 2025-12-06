package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"shumi/internal/database"
	"shumi/internal/worker"
	"shumi/internal/tgauth"
	"shumi/internal/socialapirequests"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbPassword := os.Getenv("POSTGRESQL_PASSWORD")

	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=shumi port=5432 sslmode=disable TimeZone=Asia/Yekaterinburg", dbPassword)
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	auth := tgauth.AuthClient{
		DB: dbConnection,
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://qh5zm0g8-5173.euw.devtunnels.ms"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          true,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/api/auth", auth.SetTokens)

	userDataHandler := database.UserDataHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/me/", tgauth.AuthMiddleware(userDataHandler))

	voteHandler := database.VoteHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/vote/", tgauth.AuthMiddleware(voteHandler))

	workerHandler := worker.WorkerHandler{}
	mux.Handle("/api/worker/", tgauth.AuthMiddleware(workerHandler))

	gameHandler := socialapirequests.GamesHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/games", tgauth.AuthMiddleware(gameHandler))

	log.Println("Listen started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(mux)))
}