package transaction

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Surname string `gorm:"not null"`
	Email   string `gorm:"not null"`
}

func (a *Account) EmailForSes() []string {
	return []string{a.Email}
}
