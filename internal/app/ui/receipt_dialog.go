package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/deut/garage-accounting/config/translate"
	"github.com/deut/garage-accounting/internal/services"
)

type ReceiptDialog struct {
	garageNumber   string
	rateService    *services.Rate
	paymentService *services.Payment
	window         fyne.Window
	refresh        func(...bool)
}

func NewReceiptDialog(garageNumber string, window fyne.Window, refresh func(...bool)) *ReceiptDialog {
	return &ReceiptDialog{
		garageNumber:   garageNumber,
		rateService:    services.NewRate(),
		paymentService: services.NewPayment(),
		window:         window,
		refresh:        refresh,
	}
}

func (rd *ReceiptDialog) Build() dialog.Dialog {
	ratesByYears, err := rd.rateService.Rates()
	if err != nil {
		return dialog.NewError(err, rd.window)
	}

	valueBind := binding.NewString()
	valueW := widget.NewEntryWithData(valueBind)
	valueW.SetPlaceHolder(translate.T("amount"))

	years := []string{}
	for y := range ratesByYears {
		years = append(years, y)
	}

	yearStr := ""
	yearW := widget.NewSelect(years, func(s string) {
		yearStr = s
		valueW.SetText(fmt.Sprintf("%.2f", ratesByYears[s]))
	})

	yearW.PlaceHolder = translate.T("selectYearPromt")

	formItems := []*widget.FormItem{
		widget.NewFormItem("", yearW),
		widget.NewFormItem("", valueW),
	}

	return dialog.NewForm(
		translate.T("paymentFormName"),
		translate.T("create"),
		translate.T("cancel"),
		formItems,
		rd.receiptHandlerFunc(rd.garageNumber, yearStr, valueBind),
		rd.window,
	)
}

func (rd *ReceiptDialog) receiptHandlerFunc(garageNumber string, year string, value binding.String) func(bool) {
	return func(isCreate bool) {
		if isCreate {
			paymentValue, err := value.Get()
			if err != nil {
				dialog.NewError(err, rd.window).Show()
			}

			_, err = rd.paymentService.Pay(garageNumber, year, paymentValue)
			if err != nil {
				dialog.NewError(err, rd.window).Show()
			}
		}

		rd.refresh()
	}
}
