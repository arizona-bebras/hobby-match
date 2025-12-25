package database

import (
	"github.com/pgvector/pgvector-go"
	"log"
	"shumi/internal/embeddings"

	"gorm.io/gorm"
)

type UserVector struct {
	Id        string   `json:"id"`
	Interests []string `json:"interest_ids"`
	Text      string   `json:"text"`
}

func upsertUser(user *User) error {
	if user.Info == "" {
		return nil
	}
	embedding, err := embeddings.GenerateEmbedding(user.Info)
	log.Printf("embedding: %s", embedding)

	if err != nil {
		log.Printf("upsert: failed to create user vector: %v", err)
		return err
	}

	user.InfoEmbedding = pgvector.NewVector(embedding)
	return nil
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	return upsertUser(u)
}

func (u *User) BeforeUpdate(_ *gorm.DB) error {
	return upsertUser(u)
}
