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

	// 3. generate reports from csv file content
	reports, err := transaction.CreateReports(fileContent)
	if err != nil {
		panic(err)
	}

	// 4. for each report

	// a. generate new csv file from report

	// b. upload csv file to s3 bucket

	// c. add some s3 reference to report

	// d. generate email template

	// e. send email to customer with amazon ses

	// 5. save reports data to postgres with some s3 reference
	err = transaction.SaveInDatabase[transaction.Report](reports)
	if err != nil {
		panic(err)
	}
}
