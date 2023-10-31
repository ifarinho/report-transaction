package transaction

import (
	"github.com/shopspring/decimal"
	"time"
)

type MonthSummary map[time.Month]*Movement

type Report struct {
	TotalBalance decimal.Decimal
	MonthSummary MonthSummary
}

func (r *Report) AddTotalBalance(amount decimal.Decimal) {
	r.TotalBalance = r.TotalBalance.Add(amount)
}

func (r *Report) AverageTotalDebit() (decimal.Decimal, error) {
	balance := Balance{}

	for _, movement := range r.MonthSummary {
		balance.Value = balance.Value.Add(movement.Debit.Value)
		balance.Counter += movement.Debit.Counter
	}

	return averageBalance(balance)
}

func (r *Report) AverageTotalCredit() (decimal.Decimal, error) {
	balance := Balance{}

	for _, movement := range r.MonthSummary {
		balance.Value = balance.Value.Add(movement.Credit.Value)
		balance.Counter += movement.Credit.Counter
	}

	return averageBalance(balance)
}

func CreateReport(transactions []Transaction) (*Report, error) {
	report := &Report{}

	monthSummary := make(MonthSummary)

	for _, transaction := range transactions {
		report.AddTotalBalance(transaction.Amount)

		if movement, ok := monthSummary[transaction.Month()]; ok {
			movement.UpdateBalance(transaction.Amount)
			continue
		}

		movement := &Movement{}
		movement.UpdateBalance(transaction.Amount)
		monthSummary[transaction.Month()] = movement
	}

	report.MonthSummary = monthSummary

	return report, nil
}
