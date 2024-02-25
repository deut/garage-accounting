package services

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
	"github.com/deut/garage-accounting/internal/models"
	"github.com/go-playground/validator"
)

const (
	GarageNumber = "garageNumber"
	FullName     = "fullName"
	PhoneNumber  = "phoneNumber"
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
	a.collection, err = a.model.GetAll()

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
	case GarageNumber:
		searchFunc = models.ByGarageNumber(value)
	case FullName:
		searchFunc = models.ByFullName(value)
	case PhoneNumber:
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

func (a *Account) CreateFromBindings(bindings ...binding.String) error {
	if len(bindings) < 4 {
		return fmt.Errorf("wrong argument set")
	}

	garageNum, err := bindings[0].Get()
	if err != nil {
		return fmt.Errorf("cannot read garage number from binding 0")
	}

	FullName, err := bindings[1].Get()
	if err != nil {
		return fmt.Errorf("cannot read full name from binding 1")
	}

	phone, err := bindings[2].Get()
	if err != nil {
		return fmt.Errorf("cannot read phone from  binding 2")
	}

	address, err := bindings[3].Get()
	if err != nil {
		return fmt.Errorf("cannot read address from binding 3")
	}

	return a.Create(garageNum, FullName, phone, address)
}

func (a *Account) Create(garageNum, FullName, phone, address string) error {
	a.model = models.Account{
		GarageNumber: garageNum,
		FullName:     FullName,
		PhoneNumber:  phone,
		Address:      address,
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
			a.GarageNumber,
			a.FullName,
			a.PhoneNumber,
			a.Address,
		}

		table = append(table, t)
	}

	return table
}
