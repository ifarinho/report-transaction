package transaction

import (
	"github.com/shopspring/decimal"
	"report-transaction/internal/db"
	"strconv"
	"time"
)

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
		UserId: uint(id),
		Date:   date,
		Amount: amount,
	}, nil
}

func RowWriter(report Report) ([]string, error) {
	return nil, nil
}

func CreateReports(transactions []Transaction) ([]Report, error) {
	userReportMap := make(map[uint][]Transaction)

	for _, transaction := range transactions {
		userReportMap[transaction.UserId] = append(userReportMap[transaction.UserId], transaction)
	}

	reports := make([]Report, 0)

	for userId, userTransactions := range userReportMap {
		report := Report{UserId: userId}

		for _, transaction := range userTransactions {
			if transaction.Date.Before(report.StartDate) {
				report.StartDate = transaction.Date
			}

			if transaction.Date.After(report.EndDate) {
				report.EndDate = transaction.Date
			}

			amount := transaction.Amount

			report.TotalBalance.Add(amount)

			if amount.IsNegative() {
				report.Debit.Add(amount)
			} else {
				report.Credit.Add(amount)
			}

			report.AddMonthCount(transaction.Date.Month())

			reports = append(reports, report)
		}
	}

	return reports, nil
}

func SaveInDatabase[T Transaction | Report](transactions []T) error {
	return db.BatchInsert[T](transactions)
}
