package services

import "github.com/deut/garage-accounting/internal/models"

type Rate struct {
	rate *models.Rate
}

func NewRate() *Rate {
	return &Rate{rate: &models.Rate{}}
}

func (r *Rate) Rates() (map[string]float32, error) {
	rates, err := r.rate.All()
	if err != nil {
		return nil, err
	}

	yearByValue := map[string]float32{}
	for _, rate := range rates {
		yearByValue[rate.Year] = rate.Value
	}

	return yearByValue, nil
}
