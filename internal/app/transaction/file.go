package transaction

import (
	"bytes"
	"encoding/csv"
	"report-transaction/internal/app/awsdk"
	"report-transaction/internal/app/env"
	"report-transaction/internal/app/tools/calculate"
	"report-transaction/internal/app/tools/datetime"
	"time"
)

const (
	transactionIdRow int = iota
	dateRow
	amountRow
)

func RowParser(record []string) (*Transaction, error) {
	transactionId, err := calculate.ParseUint(record[transactionIdRow])
	if err != nil {
		return nil, err
	}

	date, err := time.Parse(time.RFC3339, record[dateRow])
	if err != nil {
		return nil, err
	}

	amount, err := calculate.ParseDecimal(record[amountRow])
	if err != nil {
		return nil, err
	}

	return &Transaction{
		TransactionId: transactionId,
		Date:          datetime.TimeInUtc(date),
		Amount:        amount,
	}, nil
}

func GetFileFromBucket(key string) (*csv.Reader, error) {
	content, err := awsdk.GetObject(env.AwsFullPath + key)
	if err != nil {
		return nil, err
	}
	return csv.NewReader(bytes.NewBuffer(content)), nil
}
