package main

import (
	"log"
	"report-transaction/internal/app/awsdk"
	"report-transaction/internal/app/db"
	"report-transaction/internal/pkg/event"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatalf("fatal: failed postgres database connection: %v", err)
	}

	if err := awsdk.Init(); err != nil {
		log.Fatalf("fatal: failed aws session creation: %v", err)
	}

	event.LambdaStart()
}
