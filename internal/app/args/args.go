package args

import (
	"flag"
)

var (
	AccountId = flag.Uint("account", 0, "Account id")
	Filename  = flag.String("filename", "", "S3 bucket filename")
)

func Parse() error {
	flag.Parse()
	return nil
}
