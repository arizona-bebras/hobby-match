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


// SetTokens
// @Summary Получить токен авторизации
// @Produce json
// @Param initData body string true "tma initData"
// @Success 200 {string} string "Токен выдан"
// @Failure 401 {string} string "Не валидный токен авторизации"
// @Failure 403 {string} string "Пользователь не записан в базу данных"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /api/auth [post]
func (auth *AuthClient) SetTokens(w http.ResponseWriter, r *http.Request) {
	log.Println("Trying auth")
	token := os.Getenv("BOT_TOKEN")
	jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SERCET")


	authParts := strings.Split(r.Header.Get("authorization"), " ")
	if len(authParts) != 2 {
		log.Println("set tokens: tma token not found")
		http.Error(
			w, 
			database.JSONErr(
				http.StatusUnauthorized, 
				"set tokens: tma token not found",
			), 
			http.StatusUnauthorized,
		)
		return
	}

	authType := authParts[0]
	authData := authParts[1]
	
	switch authType {
		case "tma":
			if err := initdata.Validate(authData, token, time.Hour); err != nil {
				log.Printf("set tokens: invalid tma token, %v", err)
				http.Error(
					w, 
					database.JSONErr(
						http.StatusUnauthorized, 
						fmt.Sprintf("set tokens: invalid tma token, %v", err),
					), 
					http.StatusUnauthorized,
				)
				return
			}

			initData, err := initdata.Parse(authData)
			if err != nil {
				log.Printf("set tokens: failed to parse auth data, %v", err)
				http.Error(
					w, 
					database.JSONErr(
						http.StatusInternalServerError, 
						fmt.Sprintf("set tokens: failed to parse auth data, %v", err),
					), 
					http.StatusInternalServerError,
				)
				return
			}

			log.Println("Got initdata")
			log.Printf("TG ID %d", initData.User.ID)

			var user database.User
			result := auth.DB.Table("users").Take(&user, "id = ?", strconv.FormatInt(initData.User.ID, 10))
			log.Printf("result %v", result.Error)
			if result.Error != nil {
				log.Println(result.Error)
				if result.Error == gorm.ErrRecordNotFound {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("You must register at first\nText /start to bot"))
					w.Write([]byte("\n\n"))
					return
				}
				log.Printf("set tokens: user isn't registred, %v", err)
				http.Error(
					w, 
					database.JSONErr(
						http.StatusInternalServerError, 
						fmt.Sprintf("set tokens: user isn't registred, %v", err),
					), 
					http.StatusInternalServerError,
				)
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
				log.Printf("set tokens: failed to sign claims, %v", err)
				http.Error(
					w, 
					database.JSONErr(
						http.StatusInternalServerError, 
						fmt.Sprintf("set tokens: failed to sign claims, %v", err),
					), 
					http.StatusInternalServerError,
				)
				return
			}

			jwtRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 10 * time.Hour)),
				})

			jwtRefreshString, err := jwtRefreshToken.SignedString([]byte(jwtRefreshSecret))
			if err != nil {
				log.Printf("set tokens: failed to sign claims, %v", err)
				http.Error(
					w, 
					database.JSONErr(
						http.StatusInternalServerError, 
						fmt.Sprintf("set tokens: failed to sign claims, %v", err),
					), 
					http.StatusInternalServerError,
				)
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
				log.Printf("set tokens: failed to marshal user data, %v", err)
				http.Error(
					w, 
					database.JSONErr(
						http.StatusInternalServerError, 
						fmt.Sprintf("set tokens: failed to marshal user data, %v", err),
					), 
					http.StatusInternalServerError,
				)
				return
			}

			log.Println("Tokens set")

			w.Write([]byte(fmt.Sprintf(`{"response": {"access_token": "%s" , "refresh_token": "%s", "user": %s}}`, jwtAccessString, jwtRefreshString, string(userJSON))))
			w.Write([]byte("\n\n"))
		default:
			log.Println("set tokens: auth data isn`t tma!")
			http.Error(
				w, 
				database.JSONErr(
					http.StatusBadRequest, 
					"set tokens: auth data isn`t tma!",
				), 
				http.StatusBadRequest,
			)
			return
		}
}