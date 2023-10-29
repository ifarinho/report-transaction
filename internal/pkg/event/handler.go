package event

import (
	"fmt"
	"report-transaction/internal/app/args"
	"report-transaction/internal/app/db"
	"report-transaction/internal/app/file"
	"report-transaction/internal/app/notification"
	"report-transaction/internal/app/transaction"
)

func Run() {
	if *args.Mode == args.Lambda {
		Lambda()
	} else {
		Cli()
	}
}

func handler(key string, accountId uint) error {
	account, err := db.SelectById[transaction.Account](accountId)
	if err != nil {
		return err
	}

	fullPath := transaction.FileFullPath(key, accountId)

	fmt.Printf("fullPath: %v\n", fullPath)

	fileContent, err := file.CsvReader(fullPath, transaction.GetFileFromBucket, transaction.RowParser)
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
