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
	searchKey    string
}

func NewAccountsList(w fyne.Window) AccountsList {
	return AccountsList{Window: w, accountsService: services.New()}
}

func (al *AccountsList) Build() fyne.CanvasObject {
	accsTableContent, err := al.accountsService.All()
	if err != nil {
		dialog.NewError(err, al.Window).Show()
		accsTableContent = [][]string{}
	}

	headers := []tableHeader{
		{placeholder: "", text: "ID", isSearchable: false, searchKey: "ID"},
		{placeholder: "GarageNumber", text: "", isSearchable: true, searchKey: "GarageNumber"},
		{placeholder: "FullName", text: "", isSearchable: true, searchKey: "FullName"},
		{placeholder: "PhoneNumber", text: "", isSearchable: true, searchKey: "PhoneNumber"},
		{placeholder: "", text: "Address", isSearchable: false},
		{placeholder: "", text: "Ballans", isSearchable: false},
		{placeholder: "", text: "", isSearchable: false},
	}

	table := &widget.Table{}
	if len(accsTableContent) > 0 {
		table = widget.NewTable(
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
	} else {
		table = widget.NewTable(
			func() (int, int) { return 0, len(headers) },
			func() fyne.CanvasObject { return widget.NewLabel("") },
			func(i widget.TableCellID, o fyne.CanvasObject) {},
		)
	}

	table.CreateHeader = func() fyne.CanvasObject { return widget.NewEntry() }
	table.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		entry := template.(*widget.Entry)
		if headers[id.Col].isSearchable {
			entry.SetPlaceHolder("🔍  " + headers[id.Col].placeholder)
			entry.OnChanged = func(s string) {
				r, err := al.accountsService.Search(headers[id.Col].searchKey, s)

				if err != nil {
					dialog.NewError(err, al.Window).Show()
				} else {
					accsTableContent = r
					table.Refresh()
				}
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
