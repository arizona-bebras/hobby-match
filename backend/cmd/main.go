package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"shumi/internal/database"
	"shumi/internal/handlers"
	"shumi/internal/tgauth"
	"shumi/internal/botauth"

	_ "shumi/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Shumi API
// @version 1.0
// @description Shumi API

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введите: Bearer {token}

// @securityDefinitions.apikey TmaAuth
// @in header
// @name Authorization
// @description Введите: tma {data}

// @security BearerAuth
func main() {
	var prod bool
	flag.BoolVar(&prod, "prod", false, "dev/prod")
	flag.Parse()

	if !prod {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Println(err)
		}
		log.Println("env vars loaded!")
	}

	log.Printf("prod: %v", prod)

	dbPassword := os.Getenv("POSTGRESQL_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")

	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=shumi port=5432 sslmode=disable", dbAddress, dbPassword)
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = dbConnection.Exec("CREATE EXTENSION IF NOT EXISTS vector").Error
	if err != nil {
		log.Fatalf("failed to create extension 'vector', %v", err)
	}

	err = dbConnection.AutoMigrate(
		&database.User{},
		&database.TgUser{},
		&database.Widget{},
		&database.Namespace{},
		&database.Vote{},
		&database.Interest{},
		&database.View{},
		&database.NamespaceInvite{},
		&database.UserNamespace{},
	)
	if err != nil {
		log.Fatalf("failed to migrate %v", err)
	}

	auth := tgauth.AuthClient{
		DB: dbConnection,
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://qh5zm0g8-5173.euw.devtunnels.ms", "http://192.168.1.156:5173", "http://127.0.0.1:8080", "https://*.shumi.space"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          true,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/api/auth", auth.SetTokens)

	registerHandler := handlers.TgUsersHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/tg", botauth.AuthMiddleware(registerHandler))

	userDataHandler := handlers.UserDataHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/me", tgauth.AuthMiddleware(userDataHandler))
	mux.Handle("/api/me/", tgauth.AuthMiddleware(userDataHandler))

	voteHandler := handlers.VoteHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/vote", tgauth.AuthMiddleware(voteHandler))

	autocompleteHandler := handlers.AutocompleteHandler{DB: dbConnection}
	mux.Handle("/api/autocomplete", tgauth.AuthMiddleware(autocompleteHandler))

	gameHandler := handlers.GamesHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/games", tgauth.AuthMiddleware(gameHandler))

	namespaceAdminHandler := handlers.NamespaceAdminHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/namespace", tgauth.AuthMiddleware(namespaceAdminHandler))

	namespaceHandler := handlers.NamespaceHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/namespace/{namespace_id}", tgauth.AuthMiddleware(namespaceHandler))
	mux.Handle("/api/namespace/{namespace_id}/", tgauth.AuthMiddleware(namespaceHandler))

	pagesHandler := handlers.PagesHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/pages/{page_id}", tgauth.AuthMiddleware(pagesHandler))

	filesHandler := handlers.FilesHandler{
		DB: dbConnection,
	}
	mux.Handle("/api/files/{object}/{id}", tgauth.AuthMiddleware(filesHandler))
	mux.Handle("/api/files/{object}/{id}/{index}", tgauth.AuthMiddleware(filesHandler))

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	log.Println("Listen started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.JSONMiddleware(c.Handler(mux))))
}
