package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	TarrifID  int
	Tarrif    Tarrif `gorm:"references:Tarrif"`
	AccountID int
	Account   Account `gorm:"references:Account"`
	Value     float32
	Quantity  int
}
