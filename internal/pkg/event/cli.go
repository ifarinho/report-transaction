package event

import (
	"fmt"
	"report-transaction/internal/app/args"
)

func Cli() {
	bucketKey := *args.BucketKey
	accountId := *args.AccountId

	if err := validateArgument(bucketKey, "bucket-key"); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := validateArgument(accountId, "account-id"); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := handler(bucketKey, accountId); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func validateArgument[T comparable](argument T, description string) error {
	if argument == *new(T) {
		return fmt.Errorf("argument error: got empty value for %s", description)
	}
	return nil
}
