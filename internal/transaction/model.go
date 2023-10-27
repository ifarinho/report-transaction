package transaction

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type MonthSummary map[time.Month]int

type Transaction struct {
	gorm.Model
	UserId uint            `gorm:"not null"`
	Date   time.Time       `gorm:"not null"`
	Amount decimal.Decimal `gorm:"type:decimal(7,6);not null"`
}

type Report struct {
	gorm.Model
	UserId        uint            `gorm:"not null"`
	TotalBalance  decimal.Decimal `gorm:"type:decimal(7,6);not null"`
	AverageDebit  decimal.Decimal `gorm:"type:decimal(7,6);not null"`
	AverageCredit decimal.Decimal `gorm:"type:decimal(7,6);not null"`
	StartDate     time.Time       `gorm:"not null"`
	EndDate       time.Time       `gorm:"not null"`
	Debit         Balance         `gorm:"-"`
	Credit        Balance         `gorm:"-"`
	MonthSummary  MonthSummary    `gorm:"-"`
}

func (r *Report) AddMonthCount(month time.Month) {
	r.MonthSummary[month]++
}

type Balance struct {
	Value   decimal.Decimal
	Counter int64
}

func (b *Balance) Add(value decimal.Decimal) {
	b.Value.Add(value)
	b.Counter++
}
