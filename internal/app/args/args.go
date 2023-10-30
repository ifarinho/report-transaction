package args

import (
	"flag"
)

var (
	AccountId = flag.Uint("account-id", 0, "Account id")
	BucketKey = flag.String("bucket-key", "", "S3 bucket object key")
)

func Parse() error {
	flag.Parse()
	return nil
}
