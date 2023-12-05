package services

import (
	"strconv"

	"github.com/deut/garage-accounting/internal/models"
)

type Account struct {
	model      models.Account
	collection models.Accounts
}

func New() *Account {
	return &Account{
		model:      models.Account{},
		collection: models.Accounts{},
	}
}

func (a *Account) Search(params ...func(string)) ([][]string, error) {
	var err error
	if len(params) == 0 {
		a.collection, err = a.model.GetAll()
		if err != nil {
			return nil, err
		}

		table := [][]string{}
		for _, a := range a.collection {
			t := []string{
				strconv.FormatUint(uint64(a.ID), 10),
				a.GarageNumber,
				a.FirstName,
				a.LastName,
				a.PhoneNumber,
				a.Address,
			}

			table = append(table, t)
		}

		return table, nil

	}

	return nil, nil
}
