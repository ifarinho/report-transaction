package notification

import (
	"bytes"
	"html/template"
	"report-transaction/internal/app/awsdk"
	"report-transaction/internal/app/tools/stringify"
	"report-transaction/internal/app/transaction"
	"report-transaction/web/templates"
)

const emailSubject = "Account Report Summary"

type EmailContent struct {
	TotalBalance       string
	AverageTotalDebit  string
	AverageTotalCredit string
	MonthSummary       []MonthSummary
}

type MonthSummary struct {
	Month         string
	AverageDebit  string
	AverageCredit string
	Transactions  string
}

func SendEmail(report *transaction.Report, account *transaction.Account) error {
	emailContent, err := createEmailContent(report)
	if err != nil {
		return err
	}

	body, err := createTemplate(emailContent)
	if err != nil {
		return err
	}

	return awsdk.SendEmail(body, emailSubject, account.EmailForSes())
}

func createTemplate(content *EmailContent) (string, error) {
	emailTemplate, err := template.New("").Parse(templates.Email)
	if err != nil {
		return "", err
	}

	buffer := &bytes.Buffer{}

	if err = emailTemplate.Execute(buffer, content); err != nil {
		return "nil", err
	}

	return buffer.String(), nil
}

func createEmailContent(report *transaction.Report) (*EmailContent, error) {
	var monthSummary []MonthSummary

	averageTotalDebit, err := report.AverageTotalDebit()
	if err != nil {
		return nil, err
	}

	averageTotalCredit, err := report.AverageTotalCredit()
	if err != nil {
		return nil, err
	}

	for month, movement := range report.MonthSummary {
		averageMonthDebit, err := movement.AverageDebit()
		if err != nil {
			return nil, err
		}

		averageMonthCredit, err := movement.AverageCredit()
		if err != nil {
			return nil, err
		}

		monthSummary = append(monthSummary, MonthSummary{
			Month:         month.String(),
			AverageDebit:  stringify.FixedDecimal(averageMonthDebit),
			AverageCredit: stringify.FixedDecimal(averageMonthCredit),
			Transactions:  stringify.Int64(movement.Transactions()),
		})
	}

	return &EmailContent{
		TotalBalance:       stringify.FixedDecimal(report.TotalBalance),
		AverageTotalDebit:  stringify.FixedDecimal(averageTotalDebit),
		AverageTotalCredit: stringify.FixedDecimal(averageTotalCredit),
		MonthSummary:       monthSummary,
	}, nil
}
