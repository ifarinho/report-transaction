package db

import (
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
	})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(transaction.Transaction{}, transaction.Account{})
	if err != nil {
		return err
	}

	return nil
}
