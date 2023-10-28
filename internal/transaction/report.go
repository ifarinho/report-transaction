package transaction

import (
	"github.com/shopspring/decimal"
	"time"
)

type Report struct {
	TotalBalance decimal.Decimal
	MonthSummary map[time.Month]Movement
}

func (r *Report) AddMonthSummary(transaction *Transaction) {
	if movement, ok := r.MonthSummary[transaction.Month()]; ok {
		movement.UpdateBalance(transaction.Amount)
		return
	}

	movement := Movement{}
	movement.UpdateBalance(transaction.Amount)
	r.MonthSummary[transaction.Month()] = movement
}

func (r *Report) AddTotalBalance(amount decimal.Decimal) {
	r.TotalBalance.Add(amount)
}

func (r *Report) AverageTotalDebit() (decimal.Decimal, error) {
	balance := Balance{}

	for _, movement := range r.MonthSummary {
		balance.Update(movement.Debit.Value)
	}

	return averageBalance(balance)
}

func (r *Report) AverageTotalCredit() (decimal.Decimal, error) {
	balance := Balance{}

	for _, movement := range r.MonthSummary {
		balance.Update(movement.Credit.Value)
	}

	return averageBalance(balance)
}

func NewReport() *Report {
	return &Report{
		TotalBalance: decimal.Decimal{},
		MonthSummary: make(map[time.Month]Movement),
	}
}

func CreateReport(transactions []Transaction) (*Report, error) {
	report := NewReport()

	for _, transaction := range transactions {
		report.AddTotalBalance(transaction.Amount)
		report.AddMonthSummary(&transaction)
	}

	return report, nil
}
