package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/deut/garage-accounting/internal/models"
)

type CreateAccountForm struct {
	Window  fyne.Window
	Account *models.Account
}

func NewCreateAccountForm(w fyne.Window, a *models.Account) CreateAccountForm {

	return CreateAccountForm{Window: w, Account: a}
}

func (caf *CreateAccountForm) Build() fyne.CanvasObject {
	errorPlacehilder := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	errorPlacehilder.Hidden = true
	garageNum := widget.NewEntry()
	firstName := widget.NewEntry()
	lastName := widget.NewEntry()
	phone := widget.NewEntry()
	address := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: errorPlacehilder},
			{Text: "account.garage_number", Widget: garageNum},
			{Text: "account.first_name", Widget: firstName},
			{Text: "account.last_name", Widget: lastName},
			{Text: "account.phone", Widget: phone},
			{Text: "account.address", Widget: address},
		},
		OnSubmit: func() { // optional, handle form submission
			caf.Account = &models.Account{
				GarageNumber: garageNum.Text,
				FirstName:    firstName.Text,
				LastName:     lastName.Text,
				PhoneNumber:  phone.Text,
				Address:      address.Text,
			}

			err := caf.Account.Insert()
			if err != nil {
				dialog.ShowError(err, caf.Window)
			}
		},
	}

	return form
}
