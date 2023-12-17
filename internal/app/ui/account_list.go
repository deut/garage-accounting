package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/deut/garage-accounting/internal/services"
)

type AccountsList struct {
	Window          fyne.Window
	accountsService *services.Account
}

type tableHeader struct {
	placeholder  string
	text         string
	isSearchable bool
}

func NewAccountsList(w fyne.Window) AccountsList {
	return AccountsList{Window: w, accountsService: services.New()}
}

func (al *AccountsList) Build() fyne.CanvasObject {
	accsTableContent, err := al.accountsService.Search()
	if err != nil {
		dialog.NewError(err, al.Window).Show()
		accsTableContent = [][]string{}
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
			entry.OnChanged = func(s string) {

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

	table.OnSelected = func(id widget.TableCellID) {
		table.UnselectAll()
		dialog.NewForm(
			"receipt",
			"checkout",
			"cancel",
			[]*widget.FormItem{
				widget.NewFormItem("summ", widget.NewEntry()),
				widget.NewFormItem("type", widget.NewSelectEntry([]string{"rent", "electicity"})),
			},
			func(b bool) {},
			al.Window,
		).Show()
	}

	return table
}
