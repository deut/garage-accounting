package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type App struct {
	App        fyne.App
	MainWindow fyne.Window
	Height     float32
	Width      float32
	Canvas     fyne.CanvasObject
}

func NewUI(appName string, h, w float32) *App {
	a := app.New()
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
	toolbar := NewToolbarWidget(a.MainWindow)
	accountsList := NewAccountsList(a.MainWindow)
	content := container.NewBorder(toolbar, nil, nil, nil, accountsList)

	a.SetContent(content)
	a.ShowMainWindow()
}
