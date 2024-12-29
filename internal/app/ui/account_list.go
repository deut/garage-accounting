package ui

import (
	"fmt"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
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
	orderColumn    string
	orderDirection string
}

type tableHeader struct {
	placeholder      string
	text             string
	primaryControl   int
	secondaryControl int
	isControlColumnt bool
	orderKey         string
}

const (
	defaultOrderColumn    = "created_at"
	defaultOrderDirection = "DESC"
	editRowNumber         = 4
	labelControl          = iota
	entryControl
	buttonControl
)

var (
	currentOrderColumn string
	isOrderDESC        bool
)

func NewAccountsList(w fyne.Window) AccountsList {
	return AccountsList{
		Window:          w,
		accountsService: services.New(),
		orderColumn:     defaultOrderColumn,
		orderDirection:  defaultOrderDirection,
		tableHeaders: []tableHeader{
			{
				orderKey:         "garage_number",
				text:             translate.T("garageNumber"),
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				orderKey:         "full_name",
				text:             translate.T("fullName"),
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				orderKey:         "phone_number",
				text:             translate.T("phoneNumber"),
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				orderKey:         "address",
				placeholder:      "",
				text:             translate.T("address"),
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				orderKey:         "created_at",
				placeholder:      "",
				text:             translate.T("createdAt"),
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				text:             translate.T("edit"),
				primaryControl:   buttonControl,
				isControlColumnt: true,
			},
		},
	}
}

func (al *AccountsList) Build() fyne.CanvasObject {
	al.buildData(al.orderColumn, al.orderDirection)
	al.buildContentTable()
	al.setTableHeader()
	return al.table
}

func (al *AccountsList) buildData(orderColumn, orderDirection string) {
	accs, err := al.accountsService.All(orderColumn, orderDirection)
	if err != nil {
		dialog.NewError(err, al.Window).Show()
		al.accsTableContent = [][]string{}
	}

	al.accsTableContent = accs
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

func (al *AccountsList) setTableHeader() {
	al.table.CreateHeader = func() fyne.CanvasObject {
		return widget.NewButton("", func() {})

	}

	al.table.UpdateHeader = func(cellID widget.TableCellID, o fyne.CanvasObject) {
		b := o.(*widget.Button)
		columnConfig := al.tableHeaders[cellID.Col]

		if columnConfig.orderKey != currentOrderColumn {
			b.SetIcon(nil)
		}

		b.OnTapped = func() {
			currentOrderColumn = columnConfig.orderKey
			isOrderDESC = !isOrderDESC

			if columnConfig.orderKey == currentOrderColumn {
				al.orderColumn = columnConfig.orderKey
				if isOrderDESC {
					b.SetIcon(theme.MenuDropDownIcon())
					al.orderDirection = "DESC"
					al.buildData(al.orderColumn, al.orderDirection)
					al.table.Refresh()
				} else {
					b.SetIcon(theme.MenuDropUpIcon())
					al.orderDirection = "ASC"
					al.buildData(al.orderColumn, al.orderDirection)
					al.table.Refresh()
				}
			}

			al.table.Refresh()
		}

		b.SetText(columnConfig.text)
	}

	al.table.ShowHeaderRow = true

	for i, h := range al.tableHeaders {
		width := float32(8 * len(h.text))

		if width == 0 {
			width = float32(92)
		}

		al.table.SetColumnWidth(i, width)
	}
}
