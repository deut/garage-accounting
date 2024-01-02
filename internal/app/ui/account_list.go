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
}

type tableHeader struct {
	placeholder string
	text        string
	searchKey   string
}

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
				placeholder: "",
				text:        translate.T["lastPayedYear"],
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
	var err error
	al.accsTableContent, err = al.accountsService.All()
	if err != nil {
		dialog.NewError(err, al.Window).Show()
		al.accsTableContent = [][]string{}
	}

	al.buildContentTable()
	al.setTableHeader()

	return al.table
}

func (al *AccountsList) buildContentTable() {
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
			)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if len(al.accsTableContent) > 0 {

				// TODO: handle OK
				c := o.(*fyne.Container)
				l := c.Objects[0].(*widget.Label)
				b := c.Objects[1].(*widget.Button)

				if len(al.accsTableContent[0]) > i.Col {
					b.Hide()
					l.Show()
					l.SetText(al.accsTableContent[i.Row][i.Col])
				} else {
					l.Hide()
					b.Show()
					switch i.Col - len(al.accsTableContent[0]) {
					case 0:
						b.SetText("edit")
					case 1:
						b.SetText("receipt")
						b.OnTapped = func() {
							NewReceiptDialog(al.accsTableContent[i.Row][0], al.Window).Build().Show()
						}
					}
				}
			}
		},
	)
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
