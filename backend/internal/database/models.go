package database

import "github.com/lib/pq"

type TgUser struct {
	TgID        string `json:"id"`
	TgUsername  string `json:"username"`
	TgFirstname string `json:"firstname"`
}

type User struct {
	TgID      string         `json:"tg_user" gorm:"primaryKey" db:"tgid"`
	Name      string         `json:"miniapp_name" db:"name"`
	Location  string         `json:"location" db:"location"`
	Gender    string         `json:"gender" db:"gender"`
	BirthDate string         `json:"birth_date" db:"birth_date"`
	Interests pq.StringArray `json:"interests" gorm:"type:text[]" db:"interests"`
	Photo     []byte         `json:"user_photo" db:"photo"`
	Info      string         `json:"user_info" db:"info"`
	Widgets   []Widget       `json:"widgets" gorm:"-"`
}

type Widget struct {
	Id        string        `json:"id" gorm:"primaryKey" db:"id"`
	User      string        `json:"user" gorm:"column:user;type:bigint" db:"user"`
	Order     int           `json:"order" db:"order"`
	Files     pq.ByteaArray `json:"files" gorm:"type:bytea[]" db:"files"`
	Data      string        `json:"data" db:"data"`
	Namespace string        `json:"namespace" db:"namespace"`
}
