package ui

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
)

type Layout struct {
	App        fyne.App
	MainWindow fyne.Window
	Height     float32
	Width      float32
	Canvas     fyne.CanvasObject
}

func NewUI(appName string, h, w float32) *Layout {
	a := fyneApp.New()
	return &Layout{
		App:        a,
		MainWindow: a.NewWindow(appName),
		Height:     h,
		Width:      w,
	}
}

func (l *Layout) SetContent(c fyne.CanvasObject) {
	l.MainWindow.SetContent(c)
}

func (l *Layout) ShowMainWindow() {
	l.MainWindow.Resize(fyne.NewSize(l.Width, l.Height))
	l.MainWindow.ShowAndRun()
}
