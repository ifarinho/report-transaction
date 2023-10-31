package event

import (
	"report-transaction/internal/app/db"
	"report-transaction/internal/app/file"
	"report-transaction/internal/app/notification"
	"report-transaction/internal/app/transaction"
)

func handler(filename string, accountId uint) error {
	account, err := db.SelectById[transaction.Account](accountId)
	if err != nil {
		return err
	}

	fullPath := transaction.FileFullPath(filename, accountId)

	fileContent, err := file.CsvReader(fullPath, accountId, transaction.GetFileFromBucket, transaction.RowParser)
	if err != nil {
		return err
	}

	err = db.BatchInsert(fileContent)
	if err != nil {
		return err
	}

	report, err := transaction.CreateReport(fileContent)
	if err != nil {
		return err
	}

	err = notification.SendEmail(report, account)
	if err != nil {
		return err
	}

	return nil
}
