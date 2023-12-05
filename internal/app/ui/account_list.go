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

type tableHeader struct {
	placeholder  string
	text         string
	isSearchable bool
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
			l := widget.NewLabel("")

			return l
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(accsTableContent[i.Row][i.Col])
		})

	headers := []tableHeader{
		{placeholder: "", text: "ID", isSearchable: false},
		{placeholder: "GarageNumber", text: "", isSearchable: true},
		{placeholder: "FirstName", text: "", isSearchable: true},
		{placeholder: "LastName", text: "", isSearchable: true},
		{placeholder: "PhoneNumber", text: "", isSearchable: true},
		{placeholder: "", text: "Address", isSearchable: false},
		{placeholder: "", text: "", isSearchable: false},
	}

	table.CreateHeader = func() fyne.CanvasObject { return widget.NewEntry() }
	table.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		entry := template.(*widget.Entry)
		if headers[id.Col].isSearchable {
			entry.SetPlaceHolder("🔍  " + headers[id.Col].placeholder)
			// entry.ActionItem = canvas.NewImageFromResource(theme.QuestionIcon())
			// entry.Refresh()
			// entry.ActionItem.Show()
			entry.OnChanged = func(s string) {
				// TODO: Search here
				fmt.Println(s, fmt.Sprintf(", Changed: %v", id))
			}
		} else {
			entry.SetText(headers[id.Col].text)
			entry.Disable()
		}
	}

	table.ShowHeaderRow = true

	for i, h := range headers {
		var width float32
		if h.isSearchable {
			width = float32(18 * len(h.placeholder))
		} else {
			width = float32(18 * len(h.text))
		}
		table.SetColumnWidth(i, width)
	}

	return table
}
