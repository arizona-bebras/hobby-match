package botauth

import (
	"log"
	"net/http"
	"os"
	"strings"

	"shumi/internal/database"

)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authData := strings.Split(r.Header.Get("Authorization"), " ")

		if authData[0] != "Bearer" {
			log.Println("auth: auth data isn`t bearer!")
			http.Error(
				w, 
				database.JSONErr(
					http.StatusBadRequest, 
					"auth: auth data isn`t bearer!",
				), 
				http.StatusBadRequest,
			)
			return
		}

		tokenString := authData[1]

		if tokenString != os.Getenv("BOT_AUTH_TOKEN") {
			log.Println("auth: invalid token")
			http.Error(
				w, 
				database.JSONErr(
					http.StatusForbidden, 
					"auth: invalid token",
				), 
				http.StatusForbidden,
			)
			return
		}
		next.ServeHTTP(w, r)
	})
}
