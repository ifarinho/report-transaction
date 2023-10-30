package event

import (
	"fmt"
	"os"
	"report-transaction/internal/app/args"
)

const (
	success int = iota
	argumentError
	handlerError
)

func Cli() {
	bucketKey := *args.BucketKey
	accountId := *args.AccountId

	if err := validateArgument(bucketKey, "bucket-key"); err != nil {
		exit(argumentError, err)
	}

	if err := validateArgument(accountId, "account-id"); err != nil {
		exit(argumentError, err)
	}

	if err := handler(bucketKey, accountId); err != nil {
		exit(handlerError, err)
	}

	exit(success, nil)
}

func validateArgument[T comparable](argument T, description string) error {
	if argument == *new(T) {
		return fmt.Errorf("argument error: got empty value for %s", description)
	}
	return nil
}

func exit(code int, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(code)
}
