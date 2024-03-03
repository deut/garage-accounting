package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
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
	primaryControl   int
	secondaryControl int
	isControlColumnt bool
	orderKey         string
}

const (
	editRowNumber = 4
	labelControl  = iota
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
		tableHeaders: []tableHeader{
			{
				orderKey:         "garageNumber",
				text:             translate.T["garageNumber"],
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				orderKey:         "fullName",
				text:             translate.T["fullName"],
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				orderKey:         "phoneNumber",
				text:             translate.T["phoneNumber"],
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				orderKey:         "address",
				placeholder:      "",
				text:             translate.T["address"],
				primaryControl:   labelControl,
				secondaryControl: entryControl,
			},
			{
				text:             translate.T["edit"],
				primaryControl:   buttonControl,
				isControlColumnt: true,
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

		b.OnTapped = func() {
			currentOrderColumn = columnConfig.orderKey
			isOrderDESC = !isOrderDESC
			al.table.Refresh()
		}

		if columnConfig.orderKey == currentOrderColumn {
			if isOrderDESC {
				b.SetIcon(theme.MenuDropDownIcon())
			} else {
				b.SetIcon(theme.MenuDropUpIcon())
			}
		} else {
			b.SetIcon(nil)
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
