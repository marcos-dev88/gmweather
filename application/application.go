package application

import (
	"github.com/marcos-dev88/gmweather/gmweather/adapter"
	"github.com/marcos-dev88/gmweather/gmweather/service"
)

var (
	AdapterConn = adapter.NewConnection
	NewService  = service.NewWeatherService
)

type WeatherData service.CheckWeatherOut

type Application interface {
	GetWeather() (WeatherData, error)
	UpdateData(*WeatherData) error
}

type app struct {
	service service.WeatherData
	adapter adapter.APIWeather
}

func NewApp(adapter adapter.APIWeather, service service.WeatherData) Application {
	return &app{adapter: adapter, service: service}
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
