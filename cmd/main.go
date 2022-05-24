package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/marcos-dev88/gmweather/application"
)

var (
	searchRegion = make(chan string)
	chErr        = make(chan error)
)

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

	go getWeather(searchRegion, hello)
	w.ShowAndRun()

}

func getWeather(input chan string, at *widget.Label) {
	var updatedData string
	var weatherData *application.WeatherData

	for {
		select {
		case <-time.After(5 * time.Second):
			adapat := application.AdapterConn(updatedData)
			s := application.NewService(adapat)
			appS := application.NewApp(adapat, s)

			err := appS.UpdateData(weatherData)

			if err != nil {
				chErr <- err
			}

			if weatherData != nil {
				at.SetText(weatherData.TempC)
			}

			log.Printf("\n\ndataloop -> %+v\n\n", weatherData)

		case data := <-input:
			updatedData = data
			adapat := application.AdapterConn(updatedData)
			s := application.NewService(adapat)
			appS := application.NewApp(adapat, s)

			d, err := appS.GetWeather()

			if err != nil {
				chErr <- err
			}

			weatherData = &d

			at.SetText(weatherData.TempC)

			log.Printf("data -> %v", weatherData)

		case errCh := <-chErr:
			log.Fatalf("error: %v", errCh)
		}
	}
}
