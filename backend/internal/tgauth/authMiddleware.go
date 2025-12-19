package tgauth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"shumi/internal/database"

	"github.com/golang-jwt/jwt/v5"
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

		accessTokenString := authData[1]

		authClaims := TgAuthClaims{}

		jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")

		log.Println(accessTokenString)

		accessToken, err := jwt.ParseWithClaims(accessTokenString, &authClaims, func(t *jwt.Token) (any, error) {
			if expTime, err := t.Claims.GetExpirationTime(); err == nil {
				if expTime.Before(time.Now()) {
					return nil, errors.New("access token expired")
				}
			} else {
				return nil, err
			}
			return []byte(jwtAccessSecret), nil
		})

		if err != nil || !accessToken.Valid {
			if err.Error() == "access token expired" {
				log.Println("Not Authorized! Accession token expired")
				http.Error(
					w, 
					database.JSONErr(
						498, 
						"Not Authorized! Accession token expired",
					), 
					498,
				)
				return
			}
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

		ctx := context.WithValue(r.Context(), database.AuthContextKey, authClaims.TgID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
