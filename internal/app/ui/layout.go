package ui

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
)

type Layout struct {
	App        fyne.App
	MainWindow fyne.Window
	Canvas     fyne.CanvasObject
}

func NewLayout(appName string) *Layout {
	a := fyneApp.New()
	return &Layout{
		App:        a,
		MainWindow: a.NewWindow(appName),
	}
}

func (l *Layout) SetContent(c fyne.CanvasObject) {
	l.MainWindow.SetContent(c)
}

func (l *Layout) ShowMainWindow() {
	l.MainWindow.ShowAndRun()
}
