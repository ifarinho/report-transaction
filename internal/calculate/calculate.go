package calculate

import (
	"errors"
	"github.com/shopspring/decimal"
	"strconv"
)

var zeroDivisionError = errors.New("cannot divide by 0")

func DecimalDivision(dividend decimal.Decimal, divisor decimal.Decimal) (decimal.Decimal, error) {
	if divisor.IsZero() {
		return decimal.Decimal{}, zeroDivisionError
	}
	return dividend.Div(divisor), nil
}

func DecimalFromInt(i int64) decimal.Decimal {
	return decimal.NewFromInt(i)
}

func ParseDecimal(s string) (decimal.Decimal, error) {
	return decimal.NewFromString(s)
}

func ParseUint(s string) (uint64, error) {
	return strconv.ParseUint(s, 0, 10)
}
