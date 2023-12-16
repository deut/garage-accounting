package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ReceiptForm struct {
	accountID string
}

func NewReceiptForm(accountID string) *ReceiptForm {
	return &ReceiptForm{accountID: accountID}
}

func (rf *ReceiptForm) Build() fyne.CanvasObject {
	return container.New(layout.NewCenterLayout(), widget.NewLabel("Receipt form"))
}