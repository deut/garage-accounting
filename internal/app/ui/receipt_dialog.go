package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/deut/garage-accounting/internal/services"
)

type ReceiptDialog struct {
	accountID      string
	rateService    *services.Rate
	paymentService *services.Payment
	window         fyne.Window
}

func NewReceiptDialog(accountID string, window fyne.Window) *ReceiptDialog {
	return &ReceiptDialog{
		accountID:      accountID,
		rateService:    services.NewRate(),
		paymentService: services.NewPayment(),
		window:         window,
	}
}

func (rd *ReceiptDialog) Build() dialog.Dialog {
	ratesByYears, err := rd.rateService.Rates()
	if err != nil {
		return dialog.NewError(err, rd.window)
	}

	valueBind := binding.NewString()
	valueW := widget.NewEntryWithData(valueBind)
	valueW.SetPlaceHolder("amount")

	years := []string{}
	for y, _ := range ratesByYears {
		years = append(years, y)
	}

	yearStr := ""
	yearW := widget.NewSelect(years, func(s string) {
		yearStr = s
		valueW.SetText(fmt.Sprintf("%.2f", ratesByYears[s]))
	})

	formItems := []*widget.FormItem{
		widget.NewFormItem("", yearW),
		widget.NewFormItem("", valueW),
	}

	return dialog.NewForm(
		"Receipt",
		"create",
		"cancel",
		formItems,
		rd.receiptHandlerFunc(rd.accountID, yearStr, valueBind),
		rd.window,
	)
}

func (rd *ReceiptDialog) receiptHandlerFunc(accountID string, year string, value binding.String) func(bool) {
	return func(isCreate bool) {
		if isCreate {
			paymentValue, err := value.Get()
			if err != nil {
				dialog.NewError(err, rd.window).Show()
			}

			_, err = rd.paymentService.Pay(accountID, year, paymentValue)
			if err != nil {
				dialog.NewError(err, rd.window).Show()
			}
		}
	}
}
