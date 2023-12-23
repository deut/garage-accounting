package models

import (
	"fmt"

	"github.com/deut/garage-accounting/db"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	AccountID int
	Account   Account `gorm:"foreignKey:AccountID"`
	RateID    int
	Rate      Rate `gorm:"foreignKey:RateID"`
	Value     float32
}

func (p *Payment) Create(a *Account, r *Rate, value float32) (*Payment, error) {
	p.AccountID = int(a.ID)
	p.RateID = int(r.ID)
	p.Value = value

	err := db.DB.Create(p).Error

	return p, fmt.Errorf("cannot create payment: %w", err)
}

func (p *Payment) All(accountID int) ([]Payment, error) {
	payments := []Payment{}
	err := db.DB.Find(&payments, "AccountID = ?", accountID).Preload("Rate").Error
	if err != nil {
		return nil, fmt.Errorf("cannot find payments: %w", err)
	}
	return payments, nil
}
