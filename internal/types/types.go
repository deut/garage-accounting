package types

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID           int
	GarageNumber string
	FirstName    string
	LastName     string
	PhoneNumber  string
	Address      string
}

type Bill struct {
	gorm.Model
	Account Account    `gorm:"foreignKey:AccountRefer"`
	Year    YearTariff `gorm:"foreignKey:YearTariffRefer"`
	Payed   float64
}

type YearTariff struct {
	gorm.Model
	ID    int
	Price int
}
