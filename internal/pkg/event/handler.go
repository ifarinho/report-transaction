package event

import (
	"report-transaction/internal/app/db"
	"report-transaction/internal/app/file"
	"report-transaction/internal/app/notification"
	"report-transaction/internal/app/tools/decode"
	"report-transaction/internal/app/transaction"
)

func handler(eventBody string) error {

	// 1. parse event request
	request, err := decode.DeserializeJson[Request]([]byte(eventBody))
	if err != nil {
		return err
	}

	// 2. select account
	account, err := db.SelectById[transaction.Account](request.AccountId)
	if err != nil {
		return err
	}

	// 3. read csv file
	fileContent, err := file.CsvReader(request.Key, transaction.GetFileFromBucket, transaction.RowParser)
	if err != nil {
		return err
	}

	// 4. insert in postgres the file content as transactions
	err = db.BatchInsert(fileContent)
	if err != nil {
		return err
	}

	// 5. generate report from csv file content
	report, err := transaction.CreateReport(fileContent)
	if err != nil {
		return err
	}

	// 6. send email
	err = notification.SendEmail(report, account)
	if err != nil {
		return err
	}

	return nil
}
