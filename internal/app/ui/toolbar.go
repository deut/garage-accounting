package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Toolbar struct {
	entryBinding binding.String
}

func NewToolbar() *Toolbar {
	return &Toolbar{
		entryBinding: binding.NewString(),
	}
}

func (t *Toolbar) Build() fyne.CanvasObject {
	seachEntry := widget.NewEntryWithData(t.entryBinding)
	seachEntry.SetPlaceHolder("search")

	searchButton := widget.NewButton("🔍", func() {})

	toltbarContainer := container.NewHBox(
		container.NewCenter(seachEntry),
		searchButton,
		widget.NewSeparator(),
	)

	return toltbarContainer
}

// func (t *Toolbar) searchToobarObject() *widget.ToolbarObject {
// 	return widget.ToolbarObject{}
// }
