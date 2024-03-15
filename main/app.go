package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	fyne.Window
}

func NewApp() *App {
	app.NewWithID("github.com/matwachich/dvr-scan-gui")

	a := &App{}
	a.Window = fyne.CurrentApp().NewWindow("DVR-Scan GUI")

	return a
}
