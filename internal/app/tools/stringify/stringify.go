package stringify

import (
	"github.com/shopspring/decimal"
	"strconv"
)

func Int64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FixedDecimal(d decimal.Decimal) string {
	return d.StringFixed(2)
}
