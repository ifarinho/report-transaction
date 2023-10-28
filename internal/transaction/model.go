package transaction

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"report-transaction/internal/datetime"
	"time"
)

type MonthSummary map[time.Month]int

type CustomerReportMap map[uint]Report

type Transaction struct {
	gorm.Model
	CustomerId uint            `gorm:"not null"`
	Date       time.Time       `gorm:"not null"`
	Amount     decimal.Decimal `gorm:"type:decimal(7,6);not null"`
}

type Report struct {
	gorm.Model
	CustomerId    uint            `gorm:"not null"`
	BucketKey     string          `gorm:"not null"`
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

func (r *Report) AddTotalBalance(amount decimal.Decimal) {
	r.TotalBalance.Add(amount)
}

func (r *Report) UpdateTimePeriod(startDate time.Time, endDate time.Time) {
	if startDate.Before(r.StartDate) {
		r.StartDate = startDate
	}

	if endDate.After(r.EndDate) {
		r.EndDate = endDate
	}
}

func (r *Report) UpdateBalance(amount decimal.Decimal) {
	if amount.IsNegative() {
		r.Debit.Add(amount)
	} else {
		r.Credit.Add(amount)
	}
}

func (r *Report) CalculateAverageDebit() error {
	average, err := AverageBalance(r.Debit)
	if err != nil {
		return err
	}
	r.AverageDebit = average
	return nil
}

func (r *Report) CalculateAverageCredit() error {
	average, err := AverageBalance(r.Credit)
	if err != nil {
		return err
	}
	r.AverageCredit = average
	return nil
}

func (r *Report) GenerateBucketKey() {
	r.BucketKey = fmt.Sprintf(
		"%d_%s_%s.%s",
		r.CustomerId,
		datetime.ReportFormat(r.StartDate),
		datetime.ReportFormat(r.EndDate),
		csvExtension,
	)
}

type Balance struct {
	Value   decimal.Decimal
	Counter int64
}

func (b *Balance) Add(value decimal.Decimal) {
	b.Value.Add(value)
	b.Counter++
}
