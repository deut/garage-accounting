package db

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dbName string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open SQLite connection: %w", err)
	}

	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	return nil
}

func SeedRates() error {
	DB.Debug().Exec(
		`
		INSERT INTO 'rates' ('year', 'value') VALUES
		('2015', 300.00),
		('2016', 400.00),
		('2017', 500.00),
		('2018', 600.00),
		`,
	)
	return nil
}
