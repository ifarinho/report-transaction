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
	Amount        decimal.Decimal `gorm:"type:decimal(20,8);not null" sql:"type:decimal(20,8);"`
}

func (t *Transaction) Month() time.Month {
	return t.Date.Month()
}
