package event

import (
	"fmt"
	"report-transaction/internal/app/args"
)

func Cli() error {
	filename := *args.Filename
	accountId := *args.AccountId

	if err := validateArgument(filename, "filename"); err != nil {
		return err
	}

	if err := validateArgument(accountId, "account"); err != nil {
		return err
	}

	if err := handler(filename, accountId); err != nil {
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
