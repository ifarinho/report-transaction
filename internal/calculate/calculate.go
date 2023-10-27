package calculate

import (
	"errors"
	"github.com/shopspring/decimal"
)

var zeroDivisionError = errors.New("cannot divide by 0")

func DecimalDivision(dividend decimal.Decimal, divisor decimal.Decimal) (decimal.Decimal, error) {
	if divisor.IsZero() {
		return decimal.Decimal{}, zeroDivisionError
	}
	return dividend.Div(divisor), nil
}
