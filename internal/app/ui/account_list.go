package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/deut/garage-accounting/config/translate"
	"github.com/deut/garage-accounting/internal/services"
)

type AccountsList struct {
	Window              fyne.Window
	accountsService     *services.Account
	contantTableHeaders []tableHeader
	table               *widget.Table
	accsTableContent    [][]string
	tableColumnCount    int
	editRow             int
}

type tableHeader struct {
	placeholder string
	text        string
	searchKey   string
}

const (
	editRowNumber           = 4
	editAccountButtonColumn = iota
)

func NewAccountsList(w fyne.Window) AccountsList {
	return AccountsList{
		Window:          w,
		accountsService: services.New(),
		contantTableHeaders: []tableHeader{
			{

				placeholder: translate.T["garageNumber"],
				text:        "",
				searchKey:   services.GarageNumber,
			},
			{
				placeholder: translate.T["fullName"],
				text:        "",
				searchKey:   services.FullName,
			},
			{
				placeholder: translate.T["phoneNumber"],
				text:        "",
				searchKey:   services.PhoneNumber,
			},
			{
				placeholder: "",
				text:        translate.T["address"],
			},
			{
				placeholder: "          ",
				text:        "",
			},
			{
				placeholder: "          ",
				text:        "",
			},
			{
				placeholder: "          ",
				text:        "",
			},
		},
	}
}

func (al *AccountsList) Build() fyne.CanvasObject {
	accs, err := al.accountsService.All()
	if err != nil {
		dialog.NewError(err, al.Window).Show()
		al.accsTableContent = [][]string{}
	}

	al.tableColumnCount = len(accs[0])

	// Add first empty element to have first row "sticky"
	al.accsTableContent = [][]string{{}}
	al.accsTableContent = append(al.accsTableContent, accs...)

	al.buildContentTable()
	al.setTableHeader()

	return al.table
}

func (al *AccountsList) buildContentTable() {
	al.editRow = -1
	al.table = widget.NewTable(
		func() (int, int) {
			if len(al.accsTableContent) > 0 {
				return len(al.accsTableContent), len(al.contantTableHeaders)
			} else {
				return 0, len(al.contantTableHeaders)
			}
		},
		func() fyne.CanvasObject {
			return container.NewStack(
				widget.NewLabel(""),
				widget.NewButton("", func() {}),
				widget.NewEntry(),
			)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			// TODO: handle OK
			c := o.(*fyne.Container)
			l := c.Objects[0].(*widget.Label)
			b := c.Objects[1].(*widget.Button)
			e := c.Objects[2].(*widget.Entry)

			if i.Row == 0 {
				e.Hide()
				l.Hide()
				b.Hide()

				return
			}

			if len(al.accsTableContent) == 0 {
				return
			}

			if al.tableColumnCount > i.Col && i.Row != al.editRow {
				b.Hide()
				e.Hide()
				l.Show()
				controlText := ""

				// The row can be empty
				if len(al.accsTableContent[i.Row]) == al.tableColumnCount {
					controlText = al.accsTableContent[i.Row][i.Col]
				}

				l.SetText(controlText)
				e.SetText(controlText)

			} else if i.Row < al.tableColumnCount && i.Row != al.editRow {
				e.Hide()
				l.Hide()
				b.Show()
				switch i.Col - al.tableColumnCount {
				case 0:
					b.SetText(translate.T["edit"])
					editStarted := false
					b.OnTapped = func() {
						if !editStarted {
							al.editRow = i.Row
							b.SetText(translate.T["done"])
							editStarted = true
							al.Refresh()
						} else {
							al.editRow = -1
							b.SetText(translate.T["edit"])
							editStarted = false
							al.Refresh()
						}
					}
				case 1:
					b.SetText(translate.T["paymentButton"])
					b.OnTapped = func() {
						d := NewReceiptDialog(al.accsTableContent[i.Row][0], al.Window, al.Refresh)
						d.Build().Show()
					}
				case 2:
					b.SetText(translate.T["showPayments"])
					b.OnTapped = func() {

					}
				}
			} else if al.editRow == i.Row {
				if i.Row == al.editRow {
					b.Show()
					if i.Col < editRowNumber {
						e.Show()
						l.Hide()
						b.Hide()
					} else if i.Col > editRowNumber {
						e.Hide()
						l.Hide()
						b.Hide()
					}
				}
			} else {
				e.Hide()
				l.Hide()
				b.Hide()
			}
		},
	)

	al.table.StickyRowCount = 1
}

func (al *AccountsList) Refresh() {
	al.table.Refresh()
}

func (al *AccountsList) setTableHeader() {
	al.table.CreateHeader = func() fyne.CanvasObject {
		return container.New(layout.NewStackLayout(), widget.NewLabel(""), widget.NewEntry())
	}

	al.table.UpdateHeader = func(id widget.TableCellID, o fyne.CanvasObject) {
		c := o.(*fyne.Container)
		l := c.Objects[0].(*widget.Label)
		e := c.Objects[1].(*widget.Entry)

		if al.contantTableHeaders[id.Col].searchKey != "" {
			l.Hide()
			e.Show()
			e.SetPlaceHolder("🔍  " + al.contantTableHeaders[id.Col].placeholder)
			e.OnChanged = func(s string) {
				header := al.contantTableHeaders[id.Col]
				r, err := al.accountsService.Search(header.searchKey, s)

				if err != nil {
					dialog.NewError(err, al.Window).Show()
				} else {
					al.accsTableContent = r
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
		if h.searchKey != "" {
			width = float32(8 * len(h.placeholder))
		} else {
			width = float32(8 * len(h.text))
		}

		if width == 0 {
			width = float32(92)
		}

		al.table.SetColumnWidth(i, width)
	}
}
