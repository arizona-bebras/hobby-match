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

type NamespaceMember struct {
	Id   string `json:"tg_user"`
	Name string `json:"miniapp_name"`
	Date string `json:"date"`
}

// @Description Данные неймспеса и его участники
type NamespaceMembers struct {
	Namespace        database.Namespace `json:"namespace"`
	NamespaceMembers []NamespaceMember  `json:"users"`
}

// @Description Неймспейсы, в который состоит пользователь
type UserNamespaces struct {
	Namespaces []database.Namespace `json:"namespaces"`
}

// @Description Токены для авторизации и данные пользователя
type JWTTokens struct {
	AuthToken    string        `json:"auth_token"`
	RefreshToken string        `json:"refresh_token"`
	User         database.User `json:"user"`
}

type SteamGame struct {
	AppId                 uint   `json:"appid"`
	Name                  string `json:"name"`
	Playtime              uint   `json:"playtime_forever"`
	ImgIconUrl            string `json:"img_icon_url"`
	CommunityVisibleStats bool   `json:"has_community_visible_stats"`
	ContentDescriptorids  []int  `json:"content_descriptorids"`
}

// @Description Игры из steam профиля пользователя
type Games struct {
	Games []SteamGame `json:"games"`
}

// @Description Изображение
type Photos struct {
	Files [][]byte `json:"files"`
}

// @Description Ответ на успешное создание неймспейса
type CreatedNamespace struct {
	NamespaceId string `json:"namespace_id"`
}

type TgEnterNamespaceData struct {
	UserId      string `json:"user_id"`
	NamespaceId string `json:"namepsace_id"`
}
