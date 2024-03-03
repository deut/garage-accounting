package ui

import (
	"log"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	App        fyne.App
	MainWindow fyne.Window
	Height     float32
	Width      float32
	Canvas     fyne.CanvasObject
}

func NewUI(appName string, h, w float32) *App {
	a := fyneApp.New()
	return &App{
		App:        a,
		MainWindow: a.NewWindow(appName),
		Height:     h,
		Width:      w,
	}
}

func (a *App) SetContent(c fyne.CanvasObject) {
	a.MainWindow.SetContent(c)
}

func (a *App) ShowMainWindow() {
	a.MainWindow.Resize(fyne.NewSize(a.Width, a.Height))
	a.MainWindow.ShowAndRun()
}

func (a *App) Build() {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	listAccs := NewAccountsList(a.MainWindow)
	accListObj := listAccs.Build()

	content := container.NewBorder(toolbar, nil, nil, nil, accListObj)

	a.SetContent(content)
	a.ShowMainWindow()
}
