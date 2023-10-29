package transaction

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	TransactionId uint            `gorm:"not null"`
	AccountId     uint            `gorm:"not null"`
	Date          time.Time       `gorm:"not null"`
	Amount        decimal.Decimal `gorm:"type:decimal(7,6);not null"`
}

func (t *Transaction) Month() time.Month {
	return t.Date.Month()
}
