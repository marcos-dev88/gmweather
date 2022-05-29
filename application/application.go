package application

import (
	"log"
	"time"

	"github.com/marcos-dev88/gmweather/gmweather/adapter"
	"github.com/marcos-dev88/gmweather/gmweather/service"
)

type Application interface {
	RunApp(in Input)
}

func NewApp() Application {
	return &app{}
}

func (a *app) RunApp(in Input) {
	var updatedSearch string
	var weatherData *WeatherData

	a.adapter = AdapterConn(updatedSearch)
	a.service = NewService(a.adapter)

	for {
		select {
		case <-time.After(10 * time.Second):
			err := a.UpdateData(weatherData)

			if err != nil {
				in.InputError <- err
			}

			if weatherData != nil {
				in.Label.SetText(weatherData.CurrentWeather.TempC)
			}

			log.Printf("\n\ndataloop -> %+v\n\n", weatherData)

		case data := <-in.InputSearch:
			updatedSearch = data
			a.adapter = AdapterConn(data)
			a.service = NewService(a.adapter)

			d, err := a.GetWeather()

			if err != nil {
				in.InputError <- err
			}

			weatherData = &d

			in.Label.SetText(weatherData.CurrentWeather.TempC)

			log.Printf("data -> %v", weatherData)

		case errCh := <-in.InputError:
			log.Fatalf("error: %v", errCh)
		}
	}

}

func (a *app) GetWeather() (WeatherData, error) {
	currentWeather, errC := a.GetCurrent()

	if errC != nil {
		return WeatherData{}, errC
	}

	previsionWeather, errP := a.GetPrevision()

	if errP != nil {
		return WeatherData{}, errP
	}

	return WeatherData{currentWeather, previsionWeather}, nil
}

func (a *app) UpdateData(d *WeatherData) error {

	currentWeather, errC := a.GetCurrent()

	if errC != nil {
		return errC
	}

	previsionWeather, errP := a.GetPrevision()

	if errP != nil {
		return errP
	}

	d = &WeatherData{
		CurrentWeather:   currentWeather,
		WeatherPrevision: previsionWeather,
	}

	return nil
}

func (a *app) GetCurrent() (service.CurrentWeather, error) {
	return a.service.GetCurrent()
}

func (a *app) GetPrevision() ([]service.WeatherPrevision, error) {
	return a.service.GetPrevision()
}

func (a *app) Weather() (adapter.WeatherOut, error) {
	return a.adapter.Weather()
}
