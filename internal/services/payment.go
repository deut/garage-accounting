package services

import (
	"fmt"
	"strconv"

	"github.com/deut/garage-accounting/internal/models"
)

type Payment struct {
	Account *models.Account
	Rate    *models.Rate
	Payment *models.Payment
}

func InitPayment(accountID, year, paymentValue string) (*Payment, error) {
	id, err := strconv.Atoi(accountID)
	if err != nil {
		return nil, fmt.Errorf("worong accountID (%s) has wrong value: %w", accountID, err)
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

	return &Payment{Account: acc}, nil
}
