package event

import (
	"fmt"
	"report-transaction/internal/app/env"
	"report-transaction/internal/app/tools/calculate"
)

const (
	cliMode uint = iota + 1
	lambdaMode
)

func Run() error {
	mode, err := calculate.ParseUint(env.RunMode)
	if err != nil {
		return fmt.Errorf("env error: invalid value for RUN_MODE: %v", err)
	}

	if mode == cliMode {
		return Cli()
	}

	if mode == lambdaMode {
		Lambda()
	}

	return fmt.Errorf("invalid run mode, expected cli (1) or lambda (2) but got: %v", mode)
}
