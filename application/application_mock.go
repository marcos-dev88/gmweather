package application

import (
	"errors"

	"github.com/marcos-dev88/gmweather/internal/adapter"
	"github.com/marcos-dev88/gmweather/internal/service"
)

/*====== MOCK ADAPTER ======*/
type mockAdapter struct {
	out     adapter.WeatherOut
	outErr  error
	adapter adapter.APIWeather
}

func (m mockAdapter) Weather() (adapter.WeatherOut, error) {
	return m.out, m.outErr
}

/*====== MOCK SERVICE ======*/
type mockService struct {
	outC    service.CurrentWeather
	outP    []service.WeatherPrevision
	outErrC error
	outErrP error
}

func (m mockService) GetCurrent() (service.CurrentWeather, error) {
	return m.outC, m.outErrC
}

func (m mockService) GetPrevision() ([]service.WeatherPrevision, error) {
	return m.outP, m.outErrP
}

/*====== MOCK APP ======*/
type mockApp struct {
	adapter adapter.APIWeather
	service service.WeatherData
	out     WeatherData
	outErr  error
}

func (a *mockApp) RunApp(in Input) {

}

func (a *mockApp) GetWeather() (WeatherData, error) {
	return a.out, a.outErr
}

func (a *mockApp) UpdateData(d *WeatherData) error {
	return a.outErr
}

func (a *mockApp) GetCurrent() (service.CurrentWeather, error) {
	return a.service.GetCurrent()
}

func (a *mockApp) GetPrevision() ([]service.WeatherPrevision, error) {
	return a.service.GetPrevision()
}

func (a *mockApp) Weather() (adapter.WeatherOut, error) {
	return a.adapter.Weather()
}

/* ======== MOCK VARS ======== */

var mockServiceSuccess = mockService{
	outC: service.CurrentWeather{
		FeelsLikeC:       "",
		FeelsLikeF:       "",
		Cloudcover:       "",
		Humidity:         "",
		LocalObsDateTime: "",
		ObservationTime:  "",
		PrecipInches:     "",
		PrecipMM:         "",
		Pressure:         "",
		PressureInches:   "",
		TempC:            "7",
		TempF:            "",
		UvIndex:          "",
		Visibility:       "",
		VisibilityMiles:  "",
		WeatherCode:      "",
		WeatherDesc:      nil,
		WinddirPoint:     "",
		WinddirDegree:    "",
		WindspeedKmph:    "",
		WindspeedMiles:   "",
	},
	outP: []service.WeatherPrevision{
		{
			DateTime: "",
			TempMaxC: "",
			TempMinC: "",
			TempMaxF: "",
			TempMinF: "",
			WeatherPeriod: []service.WeatherPeriod{
				{
					TimePeriod:       "",
					FeelsLikeC:       "",
					FeelsLikeF:       "",
					Cloudcover:       "",
					Humidity:         "",
					LocalObsDateTime: "",
					ObservationTime:  "",
					PrecipInches:     "",
					PrecipMM:         "",
					Pressure:         "",
					PressureInches:   "",
					TempC:            "",
					TempF:            "",
					UvIndex:          "",
					Visibility:       "",
					VisibilityMiles:  "",
					WeatherCode:      "",
					WeatherDesc: []service.WeatherDesc{
						{
							Value: "",
						},
					},
					WinddirPoint:   "",
					WinddirDegree:  "",
					WindspeedKmph:  "",
					WindspeedMiles: "",
				},
			},
		},
	},
	outErrC: nil,
	outErrP: nil,
}

var mockServiceErrorCurrent = mockService{
	outC:    service.CurrentWeather{},
	outP:    []service.WeatherPrevision{},
	outErrC: errors.New("someone stole the almanac sports"),
	outErrP: nil,
}

var mockServiceErrorPrevision = mockService{
	outC:    service.CurrentWeather{},
	outP:    []service.WeatherPrevision{},
	outErrC: nil,
	outErrP: errors.New("someone stole the almanac sports"),
}
