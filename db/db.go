package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *sql.DB

const schemaFileLocation = "sql/schema.sql"

func Connect() error {
	gormDB, err := gorm.Open(sqlite.Open("garage.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open SQLite connection: %w", err)
	}

	DB, err = gormDB.DB()

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

	r, err := DB.Exec(string(sql))
	fmt.Println(r)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("cannot run initialize database schema: %w", err)
	}

	return nil
}
