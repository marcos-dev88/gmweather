package application

import (
	"fyne.io/fyne/v2/widget"
	"github.com/marcos-dev88/gmweather/gmweather/adapter"
	"github.com/marcos-dev88/gmweather/gmweather/service"
)

var (
	AdapterConn = adapter.NewConnection
	NewService  = service.NewWeatherService
)

type WeatherData service.CheckWeatherOut

type app struct {
	service service.WeatherData
	adapter adapter.APIWeather
}

type Input struct {
	InputSearch chan string
	InputError  chan error
	Label       *widget.Label
}
