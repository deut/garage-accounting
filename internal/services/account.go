package services

import (
	"fmt"
	"strconv"

	"github.com/deut/garage-accounting/internal/models"
	"github.com/go-playground/validator"
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

func (a *Account) All() ([][]string, error) {
	accs, err := a.model.GetAll()
	if err != nil {
		return nil, err
	}

	return toTable(accs), nil
}

func (a *Account) Search(field, value string) ([][]string, error) {
	var accs models.Accounts
	var err error

	switch field {
	case "ID":
		accs, err = a.model.GetAll(models.ByID(value))
	case "GarageNumber":
		accs, err = a.model.GetAll(models.ByGarageNumber(value))
	case "FullName":
		accs, err = a.model.GetAll(models.ByFullName(value))
	case "PhoneNumber":
		accs, err = a.model.GetAll(models.ByPhoneNumber(value))
	default:
		return nil, fmt.Errorf("unknown search field: %s", field)
	}

	if err != nil {
		return nil, err
	}

	return toTable(accs), nil
}

func (a *Account) Create(garageNum, FullName, phone, address string, debt float32, electricityNumber int) error {
	a.model = models.Account{
		GarageNumber:      garageNum,
		FullName:          FullName,
		PhoneNumber:       phone,
		Address:           address,
		Debt:              debt,
		ElectricityNumber: electricityNumber,
	}

	validate := validator.New()
	err := validate.Struct(a)

	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	err = a.model.Insert()
	if err != nil {
		return fmt.Errorf("cannot create account record: %w", err)
	}

	return nil
}

func toTable(accs models.Accounts) [][]string {
	table := [][]string{}
	for _, a := range accs {
		t := []string{
			strconv.FormatUint(uint64(a.ID), 10),
			a.GarageNumber,
			a.FullName,
			a.PhoneNumber,
			a.Address,
			fmt.Sprintf("%.2f", a.Debt),
			fmt.Sprintf("%d", a.ID),
			a.LastPayedYear(),
		}

		table = append(table, t)
	}

	return table
}
