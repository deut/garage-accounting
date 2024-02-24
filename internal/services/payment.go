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

func (p *Payment) Pay(garageNumber, year, paymentValue string) (*models.Payment, error) {
	acc := &models.Account{}
	acc, err := acc.FindByGarageNumber(garageNumber)
	if err != nil {
		return nil, fmt.Errorf("account with GarageNumber = %s is not found: %w", garageNumber, err)
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
	_, err = payment.Create(acc, rate, float32(value))
	if err != nil {
		return nil, fmt.Errorf("payment failed: %w", err)
	}

	return payment, nil
}

func (p *Payment) ListPayments(accountID string) ([]float32, error) {
	id, err := strconv.Atoi(accountID)
	if err != nil {
		return nil, fmt.Errorf("accountID is incorrect: %w", err)
	}

	paymentsRecords, err := p.payment.All(id)
	if err != nil {
		return nil, fmt.Errorf("error loading payments: %w", err)
	}

	payments := []float32{}
	for _, pr := range paymentsRecords {
		payments = append(payments, pr.Value)
	}

	return payments, nil
}
