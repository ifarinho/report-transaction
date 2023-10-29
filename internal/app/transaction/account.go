package transaction

import (
	"gorm.io/gorm"
	"report-transaction/internal/app/db"
)

type Account struct {
	gorm.Model
	Email string `gorm:"not null"`
}

func (a *Account) EmailForSes() []string {
	return []string{a.Email}
}

func SelectAccountById(id uint) (*Account, error) {
	return db.SelectById[Account](id)
}
