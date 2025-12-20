package handlers

import (
	"shumi/internal/database"
)
// @Description Ответ autocomplete эндпоинта воркера.
type AutocompleteResponse struct {
	Q string
	Response []string
}

// @Description Данные регистрации пользователя
type RegisterData struct {
	TgId      string `json:"tg_id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
}

// @Description Данные анкеты (все поля database.User + tg username)
type PageData struct {
	database.User
	Username string `json:"username"`
}

// @Description Неймспейсы, в который состоит пользователь
type UserNamespaces struct {
	Namespaces []database.Namespace `json:"namespaces"`
}