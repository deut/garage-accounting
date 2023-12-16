package models

import (
	"fmt"

	"github.com/deut/garage-accounting/db"
	"gorm.io/gorm"
)

type Rate struct {
	gorm.Model
	Year     string `validate:"required" gorm:"index:idx_garage_number,unique"`
	Value    float32
	Payments []Payment `gorm:"foreignKey:RateID"`
}

func (r *Rate) FindByYear(year string) (*Rate, error) {
	err := db.DB.Model(&Rate{Year: year}).Find(r).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find Rate where year = %s: %w", year, err)
	}

	return r, nil
}
