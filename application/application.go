package application

import (
	"encoding/json"
	"log"
	"time"

	"fyne.io/fyne/v2/canvas"
	redis "github.com/marcos-dev88/gmweather/core/cache"
	"github.com/marcos-dev88/gmweather/internal/adapter"
	"github.com/marcos-dev88/gmweather/internal/service"
)

type Application interface {
	RunApp(in Input)
}

func NewApp(c redis.Cache) Application {
	return &app{cache: c}
}

func (a *app) RunApp(in Input) {
	var updatedSearch string
	var weatherData *WeatherData

	in.WeatherImg = canvas.NewImageFromFile("") //TODO: create an away to populate this
	a.adapter = AdapterConn(updatedSearch)
	a.service = NewService(a.adapter)

	for {
		select {
		case <-time.After(MinutesReloadWeatherData * time.Minute):
			err := a.UpdateData(weatherData)

			if err != nil {
				in.InputError <- err
			}

			if weatherData != nil {
				in.TempLabel.SetText(weatherData.CurrentWeather.TempC)
			}

			log.Printf("\n\ndataloop -> %+v\n\n", weatherData)

		case data := <-in.InputSearch:
			updatedSearch = data
			hasCache := true
			var bData []byte
			defer a.cache.Close()
			bData, err := a.Get(data)
			if err != nil {
				if err.Error() == "no_cache_data" {
					hasCache = false
				} else {
					in.InputError <- err
				}
			}
			if len(bData) == 0 {
				hasCache = false
			}
			if hasCache {
				if err := json.Unmarshal(bData, &weatherData); err != nil {
					in.InputError <- err
				}
			} else {
				a.adapter = AdapterConn(data)
				a.service = NewService(a.adapter)

				d, err := a.GetWeather()

				if err != nil {
					in.InputError <- err
				}

				weatherData = &d
			}
			if !hasCache {
				if weatherData != nil && len(updatedSearch) > 0 {
					// Defining redis cache
					b, _ := json.Marshal(weatherData)
					if err := a.Set(updatedSearch, b, MinutesReloadWeatherData*time.Minute); err != nil {
						in.InputError <- err
					}
				}
			}

			in.TempLabel.SetText(weatherData.CurrentWeather.TempC)
			in.LocationLabel.SetText(data)

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

func (a *app) Set(key string, data interface{}, ttl time.Duration) error {
	return a.cache.Set(key, data, ttl)
}

func (a *app) Get(key string) ([]byte, error) {
	return a.cache.Get(key)
}
