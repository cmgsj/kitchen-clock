package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()

	window := app.NewWindow(" ")
	window.SetContent(newClock())
	window.Show()

	app.Run()
}

func newClock() fyne.CanvasObject {
	clock := widget.NewLabelWithStyle(
		time.Now().Format(time.Kitchen),
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Monospace: true,
		},
	)

	go func() {
		for range time.Tick(time.Second) {
			clock.SetText(time.Now().Format(time.Kitchen))
		}
	}()

	return clock
}
