package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func (un *UserNamespace) AfterCreate(tx *gorm.DB) error {
	var namespace Namespace

	err := tx.Model(&namespace).First(&namespace, "id = ?", un.NamespaceId).Error
	if err != nil {
		log.Printf("failed to count member, %v", err)
		return fmt.Errorf("failed to count member, %v", err)
	}

	err = tx.Model(&namespace).Update("members_count", gorm.Expr("members_count + 1")).Error
	if err != nil {
		log.Printf("failed to count member, %v", err)
		return fmt.Errorf("failed to count member, %v", err)
	}

	return nil
}

func (un *UserNamespace) AfterDelete(tx *gorm.DB) error {
	var namespace Namespace

	err := tx.Model(&namespace).First(&namespace, "id = ?", un.NamespaceId).Error
	if err != nil {
		log.Printf("failed to count member, %v", err)
		return fmt.Errorf("failed to count member, %v", err)
	}

	err = tx.Model(&namespace).Update("members_count", gorm.Expr("members_count + 1")).Error
	if err != nil {
		log.Printf("failed to count member, %v", err)
		return fmt.Errorf("failed to count member, %v", err)
	}

	return nil
}