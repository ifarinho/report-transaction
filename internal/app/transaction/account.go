package transaction

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email string `gorm:"not null"`
}

func (a *Account) EmailForSes() []string {
	return []string{a.Email}
}
