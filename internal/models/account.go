package models

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/deut/garage-accounting/db"
	"github.com/go-playground/validator/v10"
)

type Account struct {
	gorm.Model
	GarageNumber string `validate:"required" gorm:"index:idx_garage_number,unique"`
	FirstName    string `validate:"required" gorm:"not null"`
	LastName     string `validate:"required" gorm:"not null"`
	Credit       float32
	Debit        float32
	PhoneNumber  string
	Address      string
}

type Accounts []Account

func (a *Account) GetAll() (Accounts, error) {
	accs := Accounts{}
	err := db.DB.Find(&accs).Error

	if err != nil {
		return nil, fmt.Errorf("cannot load account: %w", err)
	}

	return accs, nil
}

func (a *Account) Insert() error {
	validate := validator.New()
	err := validate.Struct(a)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	err = db.DB.Create(a).Error
	if err != nil {
		return fmt.Errorf("cannot create account record: %w", err)
	}

	return nil
}

func (a *Account) InitSchema() error {
	err := db.DB.AutoMigrate(a)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}

	return nil
}
