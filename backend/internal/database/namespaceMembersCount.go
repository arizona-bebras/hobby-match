package database

import (
	"gorm.io/gorm"
)

func (n *Namespace) AfterFind(tx *gorm.DB) (err error) {
	n.MembersCount = tx.Model(n).Association("Members").Count()
	return
}
