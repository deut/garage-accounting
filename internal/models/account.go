package models

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/deut/garage-accounting/db"
)

type Account struct {
	gorm.Model
	GarageNumber      string `validate:"required" gorm:"index:idx_garage_number,unique"`
	FullName          string `validate:"required" gorm:"not null"`
	PhoneNumber       string
	Address           string
	Debt              float32
	ElectricityNumber int
	Payments          []Payment `gorm:"foreignKey:AccountID"`
}

type searchParams func() (string, string)

func ByID(v string) func() (string, string) {
	return func() (string, string) { return "ID = ?", v }
}

func ByGarageNumber(v string) func() (string, string) {
	return func() (string, string) { return "garage_number LIKE ?", "%" + v + "%" }
}

func ByFullName(v string) func() (string, string) {
	return func() (string, string) { return "full_name LIKE ?", "%" + v + "%" }
}

func ByPhoneNumber(v string) func() (string, string) {
	return func() (string, string) { return "phone_number LIKE ?", "%" + v + "%" }
}

func (a *Account) GetAll(params ...searchParams) ([]Account, error) {
	accs := []Account{}
	q := db.DB.Model(&Account{}).Preload("Payments.Rate").Find(&accs)

	for _, sp := range params {
		q = q.Where(sp())
	}

	if err := q.Error; err != nil {
		return nil, fmt.Errorf("cannot load account: %w", err)
	}

	return accs, nil
}

func (a *Account) FindByID(id int) (*Account, error) {
	err := db.DB.Find(a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Account) Insert() error {
	err := db.DB.Create(a).Error
	if err != nil {
		return fmt.Errorf("cannot create account record: %w", err)
	}

	return nil
}

func (a *Account) LastPayedYear() string {
	payments := a.Payments
	lastPayment := (*Payment)(nil)
	if len(payments) > 0 {
		lastPayment = &a.Payments[len(payments)-1]

	}

	if lastPayment != nil {
		return lastPayment.Rate.Year
	} else {
		return "No payments"
	}
}
