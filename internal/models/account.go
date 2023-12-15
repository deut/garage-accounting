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

type Accounts []Account
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

func (a *Account) GetAll(params ...searchParams) (Accounts, error) {
	accs := Accounts{}
	m := db.DB.Debug().Model(Account{})

	for _, sp := range params {
		m = m.Where(sp())
	}

	fmt.Println(m.ToSQL(func(tx *gorm.DB) *gorm.DB { return db.DB }))
	err := m.Find(&accs).Error

	if err != nil {
		return nil, fmt.Errorf("cannot load account: %w", err)
	}

	return accs, nil
}

func (a *Account) Insert() error {
	err := db.DB.Create(a).Error
	if err != nil {
		return fmt.Errorf("cannot create account record: %w", err)
	}

	return nil
}
