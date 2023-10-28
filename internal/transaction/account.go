package transaction

import (
	"gorm.io/gorm"
	"report-transaction/internal/db"
)

type Account struct {
	gorm.Model
	Email string `gorm:"not null"`
}

func SelectAccountById(id int) (*Account, error) {
	return db.SelectById[Account](id)
}
