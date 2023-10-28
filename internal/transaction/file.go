package transaction

import (
	"bytes"
	"encoding/csv"
	"report-transaction/internal/awsdk"
	"report-transaction/internal/calculate"
	"report-transaction/internal/datetime"
	"report-transaction/internal/env"
	"time"
)

const (
	transactionIdRow int = iota
	accountIdRow
	dateRow
	amountRow
)

func RowParser(record []string) (*Transaction, error) {
	transactionId, err := calculate.ParseUint(record[transactionIdRow])
	if err != nil {
		return nil, err
	}

	accountId, err := calculate.ParseUint(record[accountIdRow])
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
		AccountId:     accountId,
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
