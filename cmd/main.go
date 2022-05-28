package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/marcos-dev88/gmweather/application"
)

var (
	searchRegion = make(chan string)
	chErr        = make(chan error)
)

type Input struct {
	InputSearch chan string
	InputError  chan error
	Label       *widget.Label
}

func main() {
	a := app.New()

	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello fyne!")
	inputText := widget.NewEntry()
	inputText.SetPlaceHolder("type something")

	w.SetContent(container.NewVBox(
		hello,
		inputText,
		widget.NewButton("Search", func() {
			searchRegion <- inputText.Text
		}),
	))

	in := Input{
		InputSearch: searchRegion,
		InputError:  chErr,
		Label:       hello,
	}

	app := application.NewApp()

	go app.RunApp(application.Input(in))

	w.ShowAndRun()

}
