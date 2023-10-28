package transaction

import (
	"github.com/shopspring/decimal"
	"report-transaction/internal/datetime"
	"strconv"
	"time"
)

const csvExtension = "csv"

const (
	userIdRow int = iota
	dateRow
	amountRow
)

func RowParser(record []string) (*Transaction, error) {
	id, err := strconv.ParseUint(record[userIdRow], 0, 10)
	if err != nil {
		return nil, err
	}

	date, err := time.Parse(time.RFC3339, record[dateRow])
	if err != nil {
		return nil, err
	}

	amount, err := decimal.NewFromString(record[amountRow])
	if err != nil {
		return nil, err
	}

	return &Transaction{
		CustomerId: uint(id),
		Date:       datetime.TimeInUtc(date),
		Amount:     amount,
	}, nil
}

func RowWriter(report Report) ([]string, error) {
	return nil, nil
}
