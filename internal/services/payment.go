package services

import (
	"fmt"
	"strconv"

	"github.com/deut/garage-accounting/internal/models"
)

type Payment struct {
	payment *models.Payment
}

func NewPayment() *Payment {
	return &Payment{payment: &models.Payment{}}
}

func (p *Payment) Pay(accountID, year, paymentValue string) (*models.Payment, error) {
	id, err := strconv.Atoi(accountID)
	if err != nil {
		return nil, fmt.Errorf("wrong accountID (%s) has wrong value: %w", accountID, err)
	}

	acc := &models.Account{}
	acc, err = acc.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("account with ID = %d is not found: %w", id, err)
	}

	rate := &models.Rate{}
	rate, err = rate.FindByYear(year)
	if err != nil {
		return nil, fmt.Errorf("rate with year = %s is not found: %w", year, err)
	}

	value, err := strconv.ParseFloat(paymentValue, 64)
	if err != nil {
		return nil, fmt.Errorf("wrong float value %s", paymentValue)
	}

	payment := &models.Payment{Account: *acc, Rate: *rate, Value: float32(value)}
	payment.Create(acc, rate, float32(value))

	return payment, nil
}
