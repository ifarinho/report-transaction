package main

import (
	"fmt"
	"report-transaction/internal/db"
	"report-transaction/internal/env"
	"report-transaction/internal/file"
	"report-transaction/internal/notification"
	"report-transaction/internal/transaction"
)

func main() {

}

func handler(key string, accountId int) error {
	bucketKey := fmt.Sprintf("%s/%s", env.AwsFullPath, key)

	// 1. select account
	account, err := transaction.SelectAccountById(accountId)
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
