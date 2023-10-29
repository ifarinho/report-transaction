package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"report-transaction/internal/app/env"
	"report-transaction/internal/app/transaction"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Init() error {
	var err error

	db, err = gorm.Open(postgres.Open(env.PostgresDataSourceName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		CreateBatchSize:                          1000,
	})
	if err != nil {
		return fmt.Errorf("failed postgres database connection: %v", err)
	}

	err = db.AutoMigrate(transaction.Transaction{}, transaction.Account{})
	if err != nil {
		return fmt.Errorf("failed postgres migration: %v", err)
	}

	return nil
}
