package handlers

import (
	"shumi/internal/database"
)

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

// @Description Данные неймспеса и его участники
type NamespaceMembers struct {
	Namespace database.Namespace `json:"namespace"`
	UserIds   []string           `json:"user_ids"`
}

// @Description Неймспейсы, в который состоит пользователь
type UserNamespaces struct {
	Namespaces []database.Namespace `json:"namespaces"`
}
