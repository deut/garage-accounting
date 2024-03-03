package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/deut/garage-accounting/config/translate"
	"github.com/deut/garage-accounting/internal/services"
)

type AccountsList struct {
	Window           fyne.Window
	accountsService  *services.Account
	tableHeaders     []tableHeader
	table            *widget.Table
	accsTableContent [][]string
	// editRow          int
}

type tableHeader struct {
	placeholder      string
	text             string
	searchKey        string
	primaryControl   int
	secondaryControl int
	// callbacks        *callbacks
}

// type callbacks struct {
// 	onTappend  func()
// 	onSelected func()
// }

const (
	editRowNumber = 4
	labelControl  = iota
	entryControl
	buttonControl
)

func NewAccountsList(w fyne.Window) AccountsList {
	return AccountsList{
		Window:          w,
		accountsService: services.New(),
		tableHeaders: []tableHeader{
			{

				placeholder:      translate.T["garageNumber"],
				text:             "",
				searchKey:        services.GarageNumber,
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				placeholder:      translate.T["fullName"],
				text:             "",
				searchKey:        services.FullName,
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				placeholder:      translate.T["phoneNumber"],
				text:             "",
				searchKey:        services.PhoneNumber,
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				placeholder:      "",
				text:             translate.T["address"],
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				text:           translate.T["edit"],
				primaryControl: buttonControl,
			},
		},
	}
}

func (al *AccountsList) Build() fyne.CanvasObject {
	al.buildData()
	al.buildContentTable()
	al.setTableHeader()

	return al.table
}

func (al *AccountsList) buildData() {
	accs, err := al.accountsService.All()
	if err != nil {
		dialog.NewError(err, al.Window).Show()
		al.accsTableContent = [][]string{}
	}

	// Add first empty element to have first row "sticky"
	al.accsTableContent = [][]string{{}}
	al.accsTableContent = append(al.accsTableContent, accs...)
}

func (al *AccountsList) buildContentTable() {
	// Bind all values to edit/delete records
	bindings := make([][]binding.String, len(al.accsTableContent))
	for i := range bindings {
		bindings[i] = make([]binding.String, len(al.tableHeaders))
	}

	al.table = widget.NewTable(
		func() (int, int) {
			if len(al.accsTableContent) > 0 {
				return len(al.accsTableContent), len(al.tableHeaders)
			} else {
				return 0, len(al.tableHeaders)
			}
		},
		func() fyne.CanvasObject {
			return container.NewStack(
				widget.NewLabel(""),
				widget.NewButton("", func() {}),
				widget.NewEntry(),
			)
		},
		func(cellID widget.TableCellID, cellObject fyne.CanvasObject) {
			if len(al.accsTableContent) == 0 {
				return
			}

			if cellID.Col > len(al.tableHeaders)-1 {
				dialog.NewError(fmt.Errorf("wrong table header configuration"), al.Window).Show()
				return
			}

			cellConfig := al.tableHeaders[cellID.Col]
			cellContainer := cellObject.(*fyne.Container)
			cellLabel := cellContainer.Objects[0].(*widget.Label)
			cellEntry := cellContainer.Objects[2].(*widget.Entry)
			cellButton := cellContainer.Objects[1].(*widget.Button)

			bind := binding.NewString()
			cellEntry.Bind(bind)
			bindings[cellID.Row][cellID.Col] = bind

			// Create account form
			if cellID.Row == 0 {
				if cellConfig.primaryControl == entryControl {
					cellLabel.Hide()
					cellEntry.Show()
					cellButton.Hide()
				} else if cellConfig.primaryControl == buttonControl {
					cellLabel.Hide()
					cellEntry.Hide()
					cellButton.Show()

					cellButton.SetText(translate.T["create"])
					cellButton.OnTapped = func() {
						err := al.accountsService.CreateFromBindings(bindings[cellID.Row]...)

						if err != nil {
							// TODO: Translate error
							dialog.NewError(err, al.Window).Show()
						} else {
							al.refreshTableAndContent()
						}
					}
				}

				return
			}

			switch cellConfig.primaryControl {
			case labelControl:
				cellContentText := al.accsTableContent[cellID.Row][cellID.Col]
				cellLabel.SetText(cellContentText)
				cellEntry.SetText(cellContentText)

				cellLabel.Show()
				cellEntry.Hide()
				cellButton.Hide()
			case entryControl:
				cellContentText := al.accsTableContent[cellID.Row][cellID.Col]
				cellLabel.SetText(cellContentText)
				cellEntry.SetText(cellContentText)

				cellLabel.Hide()
				cellEntry.Show()
				cellButton.Hide()
			case buttonControl:
				cellButton.SetText(cellConfig.text)

				cellLabel.Hide()
				cellEntry.Hide()
				cellButton.Show()
			default:
				cellLabel.Hide()
				cellEntry.Hide()
				cellButton.Hide()
			}
		},
	)
}

func (al *AccountsList) refreshTableAndContent() {
	r, _ := al.accountsService.All()
	al.accsTableContent = r

	al.refreshTable()
}
func (al *AccountsList) refreshTable() {
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

		if al.tableHeaders[id.Col].searchKey != "" {
			l.Hide()
			e.Show()
			e.SetPlaceHolder("🔍  " + al.tableHeaders[id.Col].placeholder)
			e.OnChanged = func(s string) {
				header := al.tableHeaders[id.Col]
				r, err := al.accountsService.Search(header.searchKey, s)

				if err != nil {
					dialog.NewError(err, al.Window).Show()
				} else {
					al.accsTableContent = r
					al.refreshTable()
				}
			}
		} else {
			l.Show()
			e.Hide()
			l.SetText(al.tableHeaders[id.Col].text)
		}
	}

	al.table.ShowHeaderRow = true

	for i, h := range al.tableHeaders {
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
