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

func (a *Account) Search(search string) ([]Account, error) {
	accs := []Account{}
	q := db.DB.Model(&Account{}).Where(
		"garage_number LIKE %?% OR full_name LIKE %?%  OR phone_number LIKE %?%",
		search,
		search,
		search,
	)

	if err := q.Find(&accs).Error; err != nil {
		return nil, fmt.Errorf("cannot find accounts: %w", err)
	}

	return accs, nil
}

func (a *Account) GetAll(orderColumn, orderDirection string) ([]Account, error) {
	accs := []Account{}
	q := db.DB.Model(&Account{}).Order(fmt.Sprintf("%s %s", orderColumn, orderDirection)).Find(&accs)

	if err := q.Error; err != nil {
		return nil, fmt.Errorf("cannot load accounts: %w", err)
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

func (a *Account) FindByGarageNumber(n string) (*Account, error) {
	err := db.DB.Find(a, "garage_number = ?", n).Error
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
