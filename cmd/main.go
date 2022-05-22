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
	chAaaa = make(chan string)
	chErr  = make(chan error)
	chOut  = make(chan *application.WeatherData)
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
			chAaaa <- inputText.Text
		}),
	))

	go getWeather(chAaaa, chOut)

	w.ShowAndRun()
}

func getWeather(input chan string, out chan *application.WeatherData) {
	var updatedData string

	for {
		select {
		case <-time.After(5 * time.Second):
			adapat := application.AdapterConn(updatedData)
			s := application.NewService(adapat)
			appS := application.NewApp(adapat, s)
			err := appS.UpdateData(<-out)
			if err != nil {
				chErr <- err
			}

			log.Printf("\n\ndataloop -> %+v\n\n", out)

		case data := <-input:
			updatedData = data
			adapat := application.AdapterConn(updatedData)
			s := application.NewService(adapat)
			appS := application.NewApp(adapat, s)

			d, err := appS.GetWeather()

			if err != nil {
				chErr <- err
			}

			out <- &d

		case errCh := <-chErr:
			log.Fatalf("error: %v", errCh)
		}
	}
}
