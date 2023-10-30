package stringify

import (
	"github.com/shopspring/decimal"
	"strconv"
)

func Int64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FixedDecimal(amount decimal.Decimal) string {
	return amount.StringFixed(2)
}
