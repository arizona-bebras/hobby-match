package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func (un *UserNamespace) AfterCreate(tx *gorm.DB) error {
	var namespace Namespace

	err := tx.First(&namespace, "id = ?", un.NamespaceId).Error
	if err != nil {
		log.Printf("failed to count member, %v", err)
		return fmt.Errorf("failed to count member, %v", err)
	}

	namespace.MembersCount += 1
	err = tx.Save(namespace).Error
	if err != nil {
		log.Printf("failed to count member, %v", err)
		return fmt.Errorf("failed to count member, %v", err)
	}

	return nil
}

func (un *UserNamespace) AfterDelete(tx *gorm.DB) error {
	var namespace Namespace

	err := tx.First(&namespace, "id = ?", un.NamespaceId).Error
	if err != nil {
		log.Printf("failed to count member, %v", err)
		return fmt.Errorf("failed to count member, %v", err)
	}

	namespace.MembersCount -= 1
	err = tx.Save(namespace).Error
	if err != nil {
		log.Printf("failed to count member, %v", err)
		return fmt.Errorf("failed to count member, %v", err)
	}

	return nil
}