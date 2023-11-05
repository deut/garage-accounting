package main

import (
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/zap"

	"github.com/deut/garage-accounting/db"
	"github.com/deut/garage-accounting/internal/models"
)

const DBName = "garage.db"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	err := db.Connect("garage.db")
	if err != nil {
		sugar.Errorf("db connection error: %v", err)
		os.Exit(1)
	}

	sugar.Info("initializing DB schema")
	a := models.Account{}
	err = a.InitSchema()
	if err != nil {
		sugar.Errorf("intet account schema error: %v", err)
		os.Exit(1)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("app.name")

	garageNum := binding.NewString()
	lGarageNum := widget.NewLabel("account.garage_number")
	eGarageNum := widget.NewEntryWithData(garageNum)

	firstName := binding.NewString()
	lFirstName := widget.NewLabel("account.first_name")
	eFirstName := widget.NewEntryWithData(firstName)

	lastName := binding.NewString()
	lLastName := widget.NewLabel("account.last_name")
	eLastName := widget.NewEntryWithData(lastName)

	phone := binding.NewString()
	lPhone := widget.NewLabel("account.phone")
	ePhone := widget.NewEntryWithData(phone)

	address := binding.NewString()
	lAddress := widget.NewLabel("account.address")
	eAddress := widget.NewEntryWithData(address)

	lCreateAcc := widget.NewLabel("account.create_account")
	bCreateAcc := widget.NewButton("account.create_account", func() {
		acc := models.Account{}
		acc.GarageNumber, _ = garageNum.Get()
		acc.FirstName, _ = firstName.Get()
		acc.LastName, _ = lastName.Get()
		acc.PhoneNumber, _ = phone.Get()
		acc.Address, _ = address.Get()

		err = acc.Insert()
		if err != nil {
			sugar.Error("db inserttion error:", zap.Error(err))
		}
	})

	grid := container.New(
		layout.NewFormLayout(),
		lGarageNum, eGarageNum,
		lFirstName, eFirstName,
		lLastName, eLastName,
		lPhone, ePhone,
		lAddress, eAddress,
		lCreateAcc, bCreateAcc,
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("account.form", grid),
		container.NewTabItem("account.list", widget.NewLabel("TODO")),
	)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
