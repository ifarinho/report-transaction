package args

import (
	"flag"
	"fmt"
)

const (
	Lambda uint = iota + 1
	Cli
)

var (
	Mode      = flag.Uint("mode", 0, "Program run mode")
	AccountId = flag.Uint("account-id", 0, "Account id")
	BucketKey = flag.String("bucket-key", "", "S3 bucket object key")
)

func Parse() error {
	flag.Parse()

	if !isValidMode(*Mode) {
		return fmt.Errorf("args error: expected lambda (%d) or cli (%d) but got: %v", Lambda, Cli, *Mode)
	}

	return nil
}

func isValidMode(value uint) bool {
	return value <= Cli
}
