package tgauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"shumi/internal/database"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"gorm.io/gorm"
)

type AuthClient struct{
	DB *gorm.DB
}

type TgAuthClaims struct {
	TgID string 
	jwt.RegisteredClaims
}

func (auth *AuthClient) SetTokens(w http.ResponseWriter, r *http.Request) {
	log.Println("Trying auth")
	token := os.Getenv("BOT_TOKEN")
	jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SERCET")


	authParts := strings.Split(r.Header.Get("authorization"), " ")
	if len(authParts) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	authType := authParts[0]
	authData := authParts[1]
	
	switch authType {
		case "tma":
			if err := initdata.Validate(authData, token, time.Hour); err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			initData, err := initdata.Parse(authData)
			if err != nil {
				http.Error(w, "internal server error", 500)
				return
			}

			log.Println("Got initdata")
			log.Printf("TG ID %d", initData.User.ID)

			var user database.User
			result := auth.DB.Table("users").Take(&user, "tg_id = ?", strconv.FormatInt(initData.User.ID, 10))
			log.Printf("result %v", result.Error)
			if result.Error != nil {
				log.Println(result.Error)
				if result.Error == gorm.ErrRecordNotFound {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("You must register at first\nText /start to bot"))
					w.Write([]byte("\n\n"))
					return
				}
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			tgAuthClaims := TgAuthClaims{
				TgID: strconv.Itoa(int(initData.User.ID)),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				},
			}

			jwtAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tgAuthClaims)
			jwtAccessString, err := jwtAccessToken.SignedString([]byte(jwtAccessSecret))
			if err != nil {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			jwtRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 10 * time.Hour)),
				})

			jwtRefreshString, err := jwtRefreshToken.SignedString([]byte(jwtRefreshSecret))
			if err != nil {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name: "access_token",
				Value: jwtAccessString,
				MaxAge: 60 * 60,
				SameSite: http.SameSiteNoneMode,
			})

			http.SetCookie(w, &http.Cookie{
				Name: "refresh_token",
				Value: jwtRefreshString,
				MaxAge: 60 * 60 * 24 * 7,
				SameSite: http.SameSiteNoneMode,
			})

			userJSON, err := json.Marshal(user)
			if err != nil {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			log.Println("Tokens set")

			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"response": {"access_token": "%s" , "refresh_token": "%s", "user": %s}}`, jwtAccessString, jwtRefreshString, string(userJSON))))
			w.Write([]byte("\n\n"))
		default:
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
}