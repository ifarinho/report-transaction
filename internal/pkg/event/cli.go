package event

import (
	"fmt"
	"report-transaction/internal/app/args"
)

func Cli() error {
	bucketKey := *args.BucketKey
	accountId := *args.AccountId

	if err := validateArgument(bucketKey, "bucket-key"); err != nil {
		return err
	}

	if err := validateArgument(accountId, "account-id"); err != nil {
		return err
	}

	if err := handler(bucketKey, accountId); err != nil {
		return err
	}

	return nil
}

func validateArgument[T comparable](argument T, description string) error {
	if argument == *new(T) {
		return fmt.Errorf("argument error: got empty value for %s", description)
	}
	return nil
}
