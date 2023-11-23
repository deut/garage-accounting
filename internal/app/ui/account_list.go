package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/deut/garage-accounting/internal/models"
)

type AccountsList struct {
	Window fyne.Window
}

func NewAccountsList(w fyne.Window) AccountsList {
	return AccountsList{Window: w}
}

func (al *AccountsList) Build() fyne.CanvasObject {
	acc := models.Account{}
	accs, _ := acc.GetAll()

	accsTableContent := make([][]string, 0, len(accs))
	for _, a := range accs {
		row := make([]string, 0, 6)
		row = append(row, strconv.FormatUint(uint64(a.ID), 10))
		row = append(row, a.GarageNumber)
		row = append(row, a.FirstName)
		row = append(row, a.LastName)
		row = append(row, a.PhoneNumber)
		row = append(row, a.Address)

		accsTableContent = append(accsTableContent, row)

	}

	fmt.Println(accsTableContent)

	list := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(accsTableContent), len(accsTableContent[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			// o.(*widget.).SetText(accsTableContent[i.Row][i.Col])
		})

	return list
}
