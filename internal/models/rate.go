package models

import "gorm.io/gorm"

type Rate struct {
	gorm.Model
	Year     string
	Value    float32
	Payments []Payment `gorm:"foreignKey:RateID"`
}
