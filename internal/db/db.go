package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"report-transaction/internal/env"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Init() *gorm.DB {
	var err error

	db, err = gorm.Open(postgres.Open(env.PostgresDataSourceName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("fatal: failed to open postgres connection: %v", err)
	}

	if err = db.AutoMigrate(); err != nil {
		log.Fatalf("fatal: failed to run migrations: %v", err)
	}

	return db
}
