package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/deut/garage-accounting/config/translate"
)

type ToolbarWidget struct {
	widget.BaseWidget
	entryBinding binding.String
	parrent      fyne.Window
}

var _ fyne.Widget = (*ToolbarWidget)(nil)

func NewToolbarWidget(parrent fyne.Window) *ToolbarWidget {
	w := &ToolbarWidget{
		parrent:      parrent,
		entryBinding: binding.NewString(),
	}
	w.ExtendBaseWidget(w)

	return w
}

func (t *ToolbarWidget) CreateRenderer() fyne.WidgetRenderer {
	seachEntry := widget.NewEntryWithData(t.entryBinding)
	seachEntry.SetPlaceHolder("search")

	searchButton := widget.NewButton(translate.T("searchSign"), func() {})

	toltbarContainer := container.NewHBox(
		container.NewCenter(seachEntry),
		searchButton,
		widget.NewSeparator(),
	)

	return widget.NewSimpleRenderer(toltbarContainer)
}
