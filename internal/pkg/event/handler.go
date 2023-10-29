package event

import (
	"fmt"
	"report-transaction/internal/app/db"
	"report-transaction/internal/app/env"
	"report-transaction/internal/app/file"
	"report-transaction/internal/app/notification"
	"report-transaction/internal/app/transaction"
)

func handler(key string, accountId uint) error {
	bucketKey := fmt.Sprintf("%s/%s", env.AwsFullPath, key)

	// 1. select account
	account, err := db.SelectById[transaction.Account](accountId)
	if err != nil {
		return err
	}

	// 2. read csv file
	fileContent, err := file.CsvReader(bucketKey, transaction.GetFileFromBucket, transaction.RowParser)
	if err != nil {
		return err
	}

	// 3. insert in postgres the file content as transactions
	err = db.BatchInsert(fileContent)
	if err != nil {
		return err
	}

	// 4. generate report from csv file content
	report, err := transaction.CreateReport(fileContent)
	if err != nil {
		return err
	}

	// 5. send email
	err = notification.SendEmail(report, account)
	if err != nil {
		return err
	}

	return nil
}
