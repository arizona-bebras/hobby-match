package database

import (
	"encoding/json"
	"log"

	// "gorm.io/gorm"
	"github.com/lib/pq"
	"github.com/pgvector/pgvector-go"
)

type contextKey string

const AuthContextKey = contextKey("TgID")

type TgUser struct {
	UserId      string `json:"id" gorm:"primaryKey;column:id"`
	TgUsername  string `json:"username" gorm:"column:username"`
	TgFirstname string `json:"firstname" gorm:"column:firstname"`
}

// @Description Основные данные профиля и связанные виджеты.
type User struct {
	Id              string          `json:"tg_user" gorm:"primaryKey;type:text;column:id"`
	Name            string          `json:"miniapp_name" gorm:"column:name"`
	Location        string          `json:"location" gorm:"column:location"`
	Gender          string          `json:"gender" gorm:"column:gender"`
	BirthDate       string          `json:"birth_date" gorm:"column:birth_date"`
	Interests       pq.StringArray  `json:"interests" gorm:"type:text[];column:interests"`
	PersonalityTest pgvector.Vector `json:"personality_test" gorm:"type:vector(5);default:{}"`
	Photo           []byte          `json:"-" gorm:"type:bytea;column:photo"`
	Info            string          `json:"user_info" gorm:"column:info"`
	Hide            bool            `json:"hide" gorm:"column:hide"`
	Widgets         []Widget        `json:"widgets" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TgUser          TgUser          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserNamespace   UserNamespace   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Vote            Vote            `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// @Description Виджет
type Widget struct {
	Id             string        `json:"id" gorm:"primaryKey;column:id"`
	UserId         string        `json:"user" gorm:"column:user;type:text"`
	User           User          `gorm:"foreignKey:UserId;references:Id"`
	Order          int           `json:"order" gorm:"column:order"`
	Files          pq.ByteaArray `json:"-" gorm:"type:bytea[];column:files"` // TODO: возвращать отдельным endpoint...
	Data           string        `json:"data" gorm:"column:data"`
	Namespace      string        `json:"namespace" gorm:"column:namespace"`
	AdditionalData string        `json:"additionalData" gorm:"-"`
	Vote           Vote          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// @Description Голос в опросе
type Vote struct {
	UserId   string `json:"user" gorm:"column:user;"`
	WidgetId string `json:"survey" gorm:"column:survey;"`
	Option   int    `json:"option" gorm:"column:option"`
}

// @Description Неймспейс
type Namespace struct {
	Id              string          `json:"id" gorm:"primaryKey;column:id"`
	Title           string          `json:"title" gorm:"column:title"`
	Picture         []byte          `json:"-" gorm:"type:bytea;column:picture"`
	Description     string          `json:"description" gorm:"column:description"`
	MembersCount    int64           `json:"members_count" gorm:"column:members_count"`
	AdminId         string          `json:"admin" gorm:"column:admin_id"`
	Admin           User            `json:"-" gorm:"foreignKey:AdminId;references:Id"`
	NamespaceInvite NamespaceInvite `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserNamespace   UserNamespace   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// @Description Инвайт код для неймспейса
type NamespaceInvite struct {
	NamespaceId string `json:"namespace_id" gorm:"column:namespace_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	InviteCode  string `json:"invite_code" gorm:"column:invite_code"`
}

func (NamespaceInvite) TableName() string {
	return "namespace_invite"
}

// @Description many-to-many Пользователь - Неймспейс + уникальня для неймспейса инфа
type UserNamespace struct {
	UserId      string `json:"user" gorm:"column:user_id"`
	NamespaceId string `json:"namespace" gorm:"column:namespace_id"`
	Date        string `json:"date" gorm:"column:date;type:Date"`
}

func (UserNamespace) TableName() string {
	return "user_namespace"
}

// @Description Интерес
type Interest struct {
	Id        string          `json:"id" gorm:"primaryKey;column:id;type:uuid;default:gen_random_uuid()"`
	Tag       string          `json:"tag" gorm:"column:tag"`
	Embedding pgvector.Vector `json:"-" gorm:"type:vector(2048)"`
}

// @Description Просмотр анкеты
type View struct {
	ViewerId    string `json:"viewer"`
	Viewer      User   `gorm:"foreignKey:ViewerId;references:Id"`
	PageOwnerId string `json:"page_owner"`
	PageOwner   User   `gorm:"foreignKey:PageOwnerId;references:Id"`
	Date        string `json:"date"`
}

// @Description Ошибка
type Error struct {
	StatusCode int    `json:"status_code"`
	Cause      string `json:"cause"`
}

func JSONErr(status int, cause string) string {
	err := Error{
		StatusCode: status,
		Cause:      cause,
	}

	JSONErr, marshalErr := json.Marshal(err)
	if marshalErr != nil {
		log.Println("failed to marshal error")
		return ""
	}

	return string(JSONErr)
}
