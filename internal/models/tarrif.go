package models

import "gorm.io/gorm"

type Tarrif struct {
	gorm.Model
	IsCurrent bool
	Price     float32
}
