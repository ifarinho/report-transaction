package main

import (
	"log"
	"report-transaction/internal/app/args"
	"report-transaction/internal/app/awsdk"
	"report-transaction/internal/app/db"
	"report-transaction/internal/pkg/event"
)

func main() {
	if err := args.Parse(); err != nil {
		log.Fatalf("fatal: %v", err)
	}

	if err := db.Init(); err != nil {
		log.Fatalf("fatal: %v", err)
	}

	if err := awsdk.Init(); err != nil {
		log.Fatalf("fatal: %v", err)
	}

	if err := event.Run(); err != nil {
		log.Fatalf("fatal: %v", err)
	}
}
