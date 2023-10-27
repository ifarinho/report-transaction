package transaction

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"report-transaction/internal/calculate"
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
	UserId       uint            `gorm:"not null"`
	TotalBalance decimal.Decimal `gorm:"type:decimal(7,6);not null"`
	Debit        Balance         `gorm:"embedded;embeddedPrefix:debit_;not null"`
	Credit       Balance         `gorm:"embedded;embeddedPrefix:credit_;not null"`
	StartDate    time.Time       `gorm:"not null"`
	EndDate      time.Time       `gorm:"not null"`
	MonthSummary MonthSummary    `gorm:"-"`
}

func (r *Report) AverageDebit() (decimal.Decimal, error) {
	return calculate.DecimalDivision(r.Debit.Value, decimal.NewFromInt(r.Debit.Counter))
}

func (r *Report) AverageCredit() (decimal.Decimal, error) {
	return calculate.DecimalDivision(r.Credit.Value, decimal.NewFromInt(r.Credit.Counter))
}

func (r *Report) AddMonthCount(month time.Month) {
	r.MonthSummary[month]++
}

type Balance struct {
	Value   decimal.Decimal `gorm:"type:decimal(7,6);not null"`
	Counter int64           `gorm:"-"`
}

func (b *Balance) Add(value decimal.Decimal) {
	b.Value.Add(value)
	b.Counter++
}
