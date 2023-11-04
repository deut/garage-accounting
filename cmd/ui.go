package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/deut/garage-accounting/internal/types"
)

func main() {
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
		acc := types.Account{}
		acc.GarageNumber, _ = garageNum.Get()
		acc.FirstName, _ = firstName.Get()
		acc.LastName, _ = lastName.Get()
		acc.PhoneNumber, _ = phone.Get()
		acc.Address, _ = address.Get()

		fmt.Println(acc)
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
