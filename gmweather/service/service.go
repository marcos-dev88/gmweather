package service

import (
	"errors"
	"strconv"

	"github.com/marcos-dev88/gmweather/gmweather/adapter"
)

type WeatherData interface {
	GetCurrent() (CurrentWeather, error)
	GetPrevision() ([]WeatherPrevision, error)
}

type service struct {
	Adapter adapter.APIWeather
}

func NewWeatherService(adapter adapter.APIWeather) WeatherData {
	return &service{Adapter: adapter}
}

func (s *service) GetCurrent() (out CurrentWeather, err error) {
	adaptData, err := s.Adapter.Weather()

	if err != nil {
		return CurrentWeather{}, err
	}

	outCheck, err := convert(adaptData)

	if err != nil {
		return CurrentWeather{}, nil
	}

	if outCheck.CurrentWeather.TempC == "" || outCheck.CurrentWeather.TempF == "" {
		return CurrentWeather{}, errors.New("Error to get prevision")
	}
	return outCheck.CurrentWeather, nil
}

func (s *service) GetPrevision() (out []WeatherPrevision, err error) {

	adaptData, err := s.Adapter.Weather()

	if err != nil {
		return nil, err
	}

	outCheck, err := convert(adaptData)

	if err != nil {
		return nil, err
	}

	for k, v := range outCheck.WeatherPrevision {

		if v.WeatherPeriod[k].TempC == "" || v.WeatherPeriod[k].TempF == "" {
			return nil, errors.New("Error to get prevision")
		}
	}

	return outCheck.WeatherPrevision, nil
}

func convert(weatherData adapter.WeatherOut) (out CheckWeatherOut, err error) {

	outData := make([]WeatherPrevision, 0)
	var outCurrent CurrentWeather

	if err != nil {
		return
	}

	for i := 0; i < len(weatherData.Weather); i++ {
		var outPeriod []WeatherPeriod
		for _, v := range weatherData.Weather[i].Hourly {

			var timePConv int

			if len(v.Time) != 0 {
				timePConv, err = strconv.Atoi(v.Time)

				if err != nil {
					return CheckWeatherOut{}, nil
				}
			}

			if convertPeriod(timePConv) == "Not Defined" {
				continue
			}

			wPeriod := WeatherPeriod{
				TempC:           v.TempC,
				TempF:           v.TempF,
				Humidity:        v.Humidity,
				FeelsLikeC:      v.FeelsLikeC,
				FeelsLikeF:      v.FeelsLikeF,
				Cloudcover:      v.Cloudcover,
				PrecipMM:        v.PrecipMM,
				Pressure:        v.Pressure,
				PrecipInches:    v.PrecipInches,
				UvIndex:         v.UvIndex,
				Visibility:      v.Visibility,
				VisibilityMiles: v.VisibilityMiles,
				WeatherCode:     v.WeatherCode,
				WinddirPoint:    v.Winddir16Point,
				WinddirDegree:   v.WinddirDegree,
				WindspeedKmph:   v.WindGustKmph,
				WindspeedMiles:  v.WindspeedMiles,
				TimePeriod:      convertPeriod(timePConv),
			}

			outPeriod = append(outPeriod, wPeriod)
		}

		w := WeatherPrevision{
			DateTime:      weatherData.Weather[i].Date,
			TempMaxC:      weatherData.Weather[i].MaxtempC,
			TempMinC:      weatherData.Weather[i].MintempC,
			TempMaxF:      weatherData.Weather[i].MaxtempF,
			TempMinF:      weatherData.Weather[i].MintempF,
			WeatherPeriod: outPeriod,
		}

		outData = append(outData, w)
	}

	for j := 0; j < len(weatherData.CurrentCondition); j++ {
		outCurrent.TempC = weatherData.CurrentCondition[j].TempC
		outCurrent.TempF = weatherData.CurrentCondition[j].TempF
	}

	out.CurrentWeather = outCurrent
	out.WeatherPrevision = outData

	return
}

func convertPeriod(period int) string {

	switch Period(period) {
	case Morning:
		return "Morning"
	case Noon:
		return "Noon"
	case Evening:
		return "Evening"
	case Night:
		return "Night"
	default:
		return "Not Defined"
	}
}
