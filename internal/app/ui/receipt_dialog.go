package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/deut/garage-accounting/internal/services"
)

type ReceiptDialog struct {
	accountID   string
	rateService *services.Rate
	window      fyne.Window
}

func NewReceiptDialog(accountID string, window fyne.Window) *ReceiptDialog {
	return &ReceiptDialog{accountID: accountID, rateService: services.NewRate(), window: window}
}

func (rd *ReceiptDialog) Build() dialog.Dialog {
	ratesByYears, err := rd.rateService.Rates()
	if err != nil {
		return dialog.NewError(err, rd.window)
	}

	valueW := widget.NewEntry()
	valueW.SetPlaceHolder("amount")

	years := []string{}
	for y, _ := range ratesByYears {
		years = append(years, y)
	}
	yearW := widget.NewSelect(years, func(s string) {
		valueW.SetText(fmt.Sprintf("%.2f", ratesByYears[s]))
	})

	accountIDw := widget.NewEntry()
	accountIDw.SetText(rd.accountID)
	accountIDw.Hide()
	formItems := []*widget.FormItem{
		widget.NewFormItem("", yearW),
		widget.NewFormItem("", valueW),
		widget.NewFormItem("", accountIDw),
	}

	return dialog.NewForm(
		"Receipt",
		"create",
		"cancel",
		formItems,
		func(bool) {},
		rd.window,
	)
}
