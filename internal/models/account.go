package models

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/deut/garage-accounting/db"
)

type Account struct {
	gorm.Model
	GarageNumber string
	FirstName    string
	LastName     string
	PhoneNumber  string
	Address      string
}

func (a *Account) Insert() error {
	err := db.DB.Create(a).Error
	if err != nil {
		return fmt.Errorf("cannot create account record: %w", err)
	}

	return nil
}
