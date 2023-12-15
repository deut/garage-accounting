package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	YearID    int
	Year      Rate
	AccountID int
	Account   Account `gorm:"foreignKey:AccountID"`
	RateID    int
	Rate      Rate `gorm:"foreignKey:RateID"`
	Value     float32
}
