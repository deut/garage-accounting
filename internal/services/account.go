package services

import (
	"fmt"
	"strconv"

	"github.com/deut/garage-accounting/internal/models"
	"github.com/go-playground/validator"
)

type Account struct {
	model      models.Account
	collection []models.Account
}

func New() *Account {
	return &Account{
		model:      models.Account{},
		collection: []models.Account{},
	}
}

func (a *Account) All() ([][]string, error) {
	var err error
	if len(a.collection) == 0 {
		a.collection, err = a.model.GetAll()
	}

	return toTable(a.collection), err
}

func (a *Account) Search(field, value string) ([][]string, error) {
	var searchFunc models.SearchQueryFunc
	var err error

	if value == "" {
		a.collection, err = a.model.GetAll()
		return toTable(a.collection), err
	}

	switch field {
	case "ID":
		searchFunc = models.ByID(value)
	case "GarageNumber":
		searchFunc = models.ByGarageNumber(value)
	case "FullName":
		searchFunc = models.ByFullName(value)
	case "PhoneNumber":
		searchFunc = models.ByPhoneNumber(value)
	default:
		return nil, fmt.Errorf("unknown search field: %s", field)
	}

	a.collection, err = a.model.Search(searchFunc)

	if err != nil {
		return nil, err
	}

	return toTable(a.collection), nil
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

func toTable(accs []models.Account) [][]string {
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
