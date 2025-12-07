package database

import (
	"github.com/lib/pq"
)

type contextKey string

const AuthContextKey = contextKey("TgID")

type TgUser struct {
	TgID        string `json:"id"`
	TgUsername  string `json:"username"`
	TgFirstname string `json:"firstname"`
}

type User struct {
	TgID      string         `json:"tg_user" gorm:"primaryKey;column:id"`
	Name      string         `json:"miniapp_name" gorm:"column:name"`
	Location  string         `json:"location" gorm:"column:location"`
	Gender    string         `json:"gender" gorm:"column:gender"`
	BirthDate string         `json:"birth_date" gorm:"column:birth_date"`
	Interests pq.StringArray `json:"interests" gorm:"type:text[];column:interests"`
	Photo     []byte         `json:"user_photo" gorm:"column:photo"`
	Info      string         `json:"user_info" gorm:"column:info"`
	Widgets   []Widget       `json:"widgets" gorm:"-"`
}

type Widget struct {
	Id             string        `json:"id" gorm:"primaryKey;column:id"`
	User           string        `json:"user" gorm:"column:user;type:bigint"`
	Order          int           `json:"order" gorm:"column:order"`
	Files          pq.ByteaArray `json:"files" gorm:"type:bytea[];column:files"`
	Data           string        `json:"data" gorm:"column:data"`
	Namespace      string        `json:"namespace" gorm:"column:namespace"`
	AdditionalData string        `json:"additionalData" gorm:"-"`
}

type Vote struct {
	User   string `json:"user" gorm:"column:user"`
	Survey string `json:"survey" gorm:"column:survey"`
	Option int    `json:"option" gorm:"column:option"`
}

type Namespace struct {
	Id          string `json:"id" gorm:"primaryKey;column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Picture     []byte `json:"picture" gorm:"column:picture"`
	Description string `json:"description" gorm:"column:description"`
	Admin       string `json:"admin" gorm:"column:admin"`
}
