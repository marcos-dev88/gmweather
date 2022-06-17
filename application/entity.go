package application

import (
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/marcos-dev88/gmweather/gmweather/adapter"
	"github.com/marcos-dev88/gmweather/gmweather/service"
)

var (
	AdapterConn = adapter.NewConnection
	NewService  = service.NewWeatherService
)

const MinutesReloadWeatherData time.Duration = 30

type WeatherData service.CheckWeatherOut

type app struct {
	service service.WeatherData
	adapter adapter.APIWeather
}

type Input struct {
	InputSearch   chan string
	InputError    chan error
	LocationLabel *widget.Label
	TempLabel     *widget.Label
	WeatherImg    *canvas.Image
}
