package transaction

import (
	"github.com/shopspring/decimal"
	"report-transaction/internal/calculate"
	"report-transaction/internal/db"
)

func CreateReports(transactions []Transaction) ([]Report, error) {
	reportMap := CreateCustomerReportMap(transactions)

	reports := make([]Report, 0)

	for _, report := range reportMap {
		if err := report.CalculateAverageDebit(); err != nil {
			return nil, err
		}

		if err := report.CalculateAverageCredit(); err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func CreateCustomerReportMap(transactions []Transaction) CustomerReportMap {
	customerReportMap := make(CustomerReportMap)

	for _, transaction := range transactions {
		amount := transaction.Amount

		if report, ok := customerReportMap[transaction.CustomerId]; ok {
			report.UpdateTimePeriod(report.StartDate, report.EndDate)
			report.AddTotalBalance(amount)
			report.UpdateBalance(amount)
			report.AddMonthCount(transaction.Date.Month())

			continue
		}

		report := Report{
			CustomerId: transaction.CustomerId,
			StartDate:  transaction.Date,
			EndDate:    transaction.Date,
		}

		report.AddTotalBalance(amount)
		report.UpdateBalance(amount)
		report.AddMonthCount(transaction.Date.Month())

		customerReportMap[transaction.CustomerId] = report
	}

	return customerReportMap
}

func SaveInDatabase[T Transaction | Report](records []T) error {
	return db.BatchInsert[T](records)
}

func AverageBalance(balance Balance) (decimal.Decimal, error) {
	return calculate.DecimalDivision(balance.Value, calculate.DecimalFromInt(balance.Counter))
}
