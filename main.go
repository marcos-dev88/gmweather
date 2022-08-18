package main

import (
	"context"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/marcos-dev88/gmweather/application"
	redis "github.com/marcos-dev88/gmweather/cache"
)

var (
	searchRegion = make(chan string)
	chErr        = make(chan error)
)

type Input application.Input

func main() {
	a := app.New()

	w := a.NewWindow("weather")

	hello := widget.NewLabel("Weather situation in: ")
	location := widget.NewLabel("--")
	tempWeather := widget.NewLabel("--")
	inputText := widget.NewEntry()
	inputText.SetPlaceHolder("type something")

	w.SetContent(container.NewVBox(
		hello,
		location,
		tempWeather,
		inputText,
		widget.NewButton("Search", func() {
			searchRegion <- inputText.Text
		}),
	))

	in := Input{
		InputSearch:   searchRegion,
		InputError:    chErr,
		TempLabel:     tempWeather,
		LocationLabel: location,
		WeatherImg:    nil,
	}

	ctx := context.Background()
	c := redis.NewClient("localhost:6079", "12345", 0)

	app := application.NewApp(redis.New(c, ctx))

	go app.RunApp(application.Input(in))

	w.ShowAndRun()

}
