package db

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const schemaFileLocation = "sql/schema.sql"

func Connect() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("garage.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open SQLite connection: %w", err)
	}

	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	return nil
}

func InitializeSchema() error {
	sql, err := os.ReadFile(schemaFileLocation)
	if err != nil {
		return fmt.Errorf("cannot read database schema from file: %w", err)
	}

	err = DB.Create(string(sql)).Error
	if err != nil {
		return fmt.Errorf("cannot run initialize database schema: %w", err)
	}

	return nil
}
