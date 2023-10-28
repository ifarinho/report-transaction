package transaction

import (
	"github.com/shopspring/decimal"
	"report-transaction/internal/calculate"
)

type Balance struct {
	Value   decimal.Decimal
	Counter int64
}

func (b *Balance) Update(value decimal.Decimal) {
	b.Value.Add(value)
	b.Counter++
}

type Movement struct {
	Credit Balance
	Debit  Balance
}

func (m *Movement) UpdateBalance(value decimal.Decimal) {
	if value.IsNegative() {
		m.Debit.Update(value)
		return
	}
	m.Credit.Update(value)
}

func (m *Movement) AverageDebit() (decimal.Decimal, error) {
	return averageBalance(m.Debit)
}

func (m *Movement) AverageCredit() (decimal.Decimal, error) {
	return averageBalance(m.Credit)
}

func (m *Movement) Transactions() int64 {
	return m.Debit.Counter + m.Credit.Counter
}

func averageBalance(balance Balance) (decimal.Decimal, error) {
	return calculate.DecimalDivision(balance.Value, calculate.DecimalFromInt(balance.Counter))
}
