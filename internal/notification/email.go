package notification

import (
	"report-transaction/internal/awsdk"
	"report-transaction/internal/transaction"
)

func SendEmail(report *transaction.Report, account *transaction.Account) error {
	template, err := CreateTemplate(report)
	if err != nil {
		return err
	}
	return awsdk.SendEmail(template, account.EmailForSes())
}

func CreateTemplate(report *transaction.Report) ([]byte, error) {
	return nil, nil
}
