package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/deut/garage-accounting/internal/services"
)

type AccountsList struct {
	Window              fyne.Window
	accountsService     *services.Account
	contantTableHeaders []tableHeader
	table               *widget.Table
}

type tableHeader struct {
	placeholder  string
	text         string
	isSearchable bool
	searchKey    string
}

func NewAccountsList(w fyne.Window) AccountsList {
	return AccountsList{
		Window:          w,
		accountsService: services.New(),
		contantTableHeaders: []tableHeader{
			{placeholder: "", text: "ID", isSearchable: false, searchKey: "ID"},
			{placeholder: "GarageNumber", text: "", isSearchable: true, searchKey: "GarageNumber"},
			{placeholder: "FullName", text: "", isSearchable: true, searchKey: "FullName"},
			{placeholder: "PhoneNumber", text: "", isSearchable: true, searchKey: "PhoneNumber"},
			{placeholder: "", text: "Address", isSearchable: false},
			{placeholder: "", text: "Debt", isSearchable: false},
			{placeholder: "", text: "ElectricityNumber", isSearchable: false},
			{placeholder: "", text: "paymentsCount", isSearchable: false},
			{placeholder: "          ", text: "", isSearchable: false},
			{placeholder: "          ", text: "", isSearchable: false},
			{placeholder: "          ", text: "", isSearchable: false},
		},
	}
}

func (al *AccountsList) Build() fyne.CanvasObject {
	accsTableContent, err := al.accountsService.All()
	if err != nil {
		dialog.NewError(err, al.Window).Show()
		accsTableContent = [][]string{}
	}

	al.buildContentTable(accsTableContent)
	al.setTableHeader(accsTableContent)

	return al.table
}

func (al *AccountsList) buildContentTable(accsTableContent [][]string) {
	al.table = widget.NewTable(
		func() (int, int) {
			if len(accsTableContent) > 0 {
				return len(accsTableContent), len(al.contantTableHeaders)
			} else {
				return 0, len(al.contantTableHeaders)
			}
		},
		func() fyne.CanvasObject {
			return container.NewStack(
				widget.NewLabel(""),
				widget.NewButton("", func() {}),
			)
		},
		al.buildTableContentFunc(accsTableContent),
	)
}

func (al *AccountsList) buildTableContentFunc(accsTableContent [][]string) func(i widget.TableCellID, o fyne.CanvasObject) {
	return func(i widget.TableCellID, o fyne.CanvasObject) {
		if len(accsTableContent) > 0 {

			// TODO: handle OK
			c := o.(*fyne.Container)
			l := c.Objects[0].(*widget.Label)
			b := c.Objects[1].(*widget.Button)

			if len(accsTableContent[0]) > i.Col {
				b.Hide()
				l.Show()
				l.SetText(accsTableContent[i.Row][i.Col])
			} else {
				l.Hide()
				b.Show()
				switch i.Col - len(accsTableContent[0]) {
				case 0:
					b.SetText("edit")
				case 1:
					b.SetText("receipt")

					b.OnTapped = func() {
						NewReceiptDialog(accsTableContent[i.Row][0], al.Window).Build().Show()
					}
				case 2:
					b.SetText("listReceipts")
				}
			}
		}
	}
}

func (al *AccountsList) setTableHeader(accsTableContent [][]string) {
	al.table.CreateHeader = func() fyne.CanvasObject {
		return container.New(layout.NewStackLayout(), widget.NewLabel(""), widget.NewEntry())
	}

	al.table.UpdateHeader = func(id widget.TableCellID, o fyne.CanvasObject) {
		c := o.(*fyne.Container)
		l := c.Objects[0].(*widget.Label)
		e := c.Objects[1].(*widget.Entry)

		if al.contantTableHeaders[id.Col].isSearchable {
			l.Hide()
			e.Show()
			e.SetPlaceHolder("🔍  " + al.contantTableHeaders[id.Col].placeholder)
			e.OnChanged = func(s string) {
				header := al.contantTableHeaders[id.Col]
				r, err := al.accountsService.Search(header.searchKey, s)

				if err != nil {
					dialog.NewError(err, al.Window).Show()
				} else {
					accsTableContent = r
					al.table.Refresh()
				}
			}
		} else {
			l.Show()
			e.Hide()
			l.SetText(al.contantTableHeaders[id.Col].text)
		}
	}

	al.table.ShowHeaderRow = true

	for i, h := range al.contantTableHeaders {
		var width float32
		if h.isSearchable {
			width = float32(16 * len(h.placeholder))
		} else {
			width = float32(16 * len(h.text))
		}

		if width == 0 {
			width = float32(70)
		}

		al.table.SetColumnWidth(i, width)
	}
}
