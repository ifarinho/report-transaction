package main

import (
	"report-transaction/internal/env"
	"report-transaction/internal/file"
	"report-transaction/internal/transaction"
)

func main() {

	// 1. read csv file
	fileContent, err := file.CsvReader[transaction.Transaction](env.FileTargetPath, true, transaction.RowParser)
	if err != nil {
		panic(err)
	}

	// 2. insert in postgres the file content as transactions
	err = transaction.SaveInDatabase[transaction.Transaction](fileContent)
	if err != nil {
		panic(err)
	}

	// 3. generate report from csv file content
	reports, err := transaction.CreateReports(fileContent)
	if err != nil {
		panic(err)
	}

	// 4. generate new csv files from reports

	// 5. upload csv files to s3 bucket

	// 6. save to postgres with some s3 reference
	err = transaction.SaveInDatabase[transaction.Report](reports)
	if err != nil {
		panic(err)
	}

	// 7. generate email template

	// 8. send email with amazon ses
}
