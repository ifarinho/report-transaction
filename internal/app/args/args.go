package args

import (
	"flag"
)

var (
	AccountId = flag.Uint("account", 0, "Account id")
	Filename  = flag.String("filename", "", "S3 bucket object key")
)

func Parse() error {
	flag.Parse()
	return nil
}
