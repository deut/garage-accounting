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
		row = append(row, "")

		accsTableContent = append(accsTableContent, row)

	}

	table := widget.NewTable(
		func() (int, int) {
			return len(accsTableContent), len(accsTableContent[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(accsTableContent[i.Row][i.Col])
		})

	header := []string{"ID", "GarageNumber", "FirstName", "LastName", "PhoneNumber", "Address", ""}
	table.CreateHeader = func() fyne.CanvasObject { return widget.NewEntry() }
	table.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		entry := template.(*widget.Entry)
		entry.SetPlaceHolder(header[id.Col])
		entry.OnChanged = func(s string) {
			// TODO: Search here
			fmt.Println(s, fmt.Sprintf(", Changed: %v", id))
		}
	}
	table.ShowHeaderRow = true

	// table.Cl = func(id widget.TableCellID) {
	// 	for i := range header {
	// 		table.Select(widget.TableCellID{Col: id.Col, Row: i})
	// 	}
	// }

	for i, h := range header {
		width := float32(18 * len(h))
		table.SetColumnWidth(i, width)
	}

	return table
}
