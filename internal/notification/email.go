package notification

import (
	"bytes"
	"html/template"
	"report-transaction/internal/awsdk"
	"report-transaction/internal/stringify"
	"report-transaction/internal/transaction"
)

const (
	emailTemplatePath = "templates/email.gohtml"
	emailSubject      = "Account Report Summary"
)

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
	templateContent, err := emailContent(report)
	if err != nil {
		return err
	}

	body, err := createTemplate(templateContent)
	if err != nil {
		return err
	}

	return awsdk.SendEmail(body, emailSubject, account.EmailForSes())
}

func createTemplate(content *EmailContent) (string, error) {
	emailTemplate, err := template.ParseFiles(emailTemplatePath)
	if err != nil {
		return "", err
	}

	buffer := &bytes.Buffer{}

	if err = emailTemplate.Execute(buffer, content); err != nil {
		return "nil", err
	}

	return buffer.String(), nil
}

func emailContent(report *transaction.Report) (*EmailContent, error) {
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
			AverageDebit:  averageMonthDebit.String(),
			AverageCredit: averageMonthCredit.String(),
			Transactions:  stringify.Int64(movement.Transactions()),
		})
	}

	return &EmailContent{
		TotalBalance:       report.TotalBalance.String(),
		AverageTotalDebit:  averageTotalDebit.String(),
		AverageTotalCredit: averageTotalCredit.String(),
		MonthSummary:       monthSummary,
	}, nil
}
