package transaction

import (
	"github.com/shopspring/decimal"
	"report-transaction/internal/calculate"
)

type Balance struct {
	Value   decimal.Decimal
	Counter int64
}

type Movement struct {
	Credit Balance
	Debit  Balance
}

func (m *Movement) UpdateBalance(value decimal.Decimal) {
	if value.IsNegative() {
		m.Debit.Value.Add(value)
		m.Debit.Counter++
		return
	}
	m.Credit.Value.Add(value)
	m.Credit.Counter++
}

func (m *Movement) AverageDebit() (decimal.Decimal, error) {
	return AverageBalance(m.Debit)
}

func (m *Movement) AverageCredit() (decimal.Decimal, error) {
	return AverageBalance(m.Credit)
}

func AverageBalance(balance Balance) (decimal.Decimal, error) {
	return calculate.DecimalDivision(balance.Value, calculate.DecimalFromInt(balance.Counter))
}
