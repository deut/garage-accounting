package models

import (
	"fmt"

	"github.com/deut/garage-accounting/db"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	AccountID int
	Account   Account `gorm:"foreignKey:AccountID;not null"`
	RateID    int
	Rate      Rate `gorm:"foreignKey:RateID;not null"`
	Value     float32
}

func (p *Payment) Create(a *Account, r *Rate, value float32) (*Payment, error) {
	p.AccountID = int(a.ID)
	p.RateID = int(r.ID)
	p.Value = value

	if err := db.DB.Create(p).Error; err != nil {
		return nil, fmt.Errorf("cannot create payment: %w", err)
	}

	return p, nil
}

func (p *Payment) All(accountID int) ([]Payment, error) {
	payments := []Payment{}
	err := db.DB.Find(&payments, "AccountID = ?", accountID).Preload("Rate").Error
	if err != nil {
		return nil, fmt.Errorf("cannot find payments: %w", err)
	}
	return payments, nil
}
