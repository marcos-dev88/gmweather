package service

import "github.com/marcos-dev88/gmweather/gmweather/adapter"

type mockAdapter struct {
	out     adapter.WeatherOut
	outErr  error
	adapter adapter.APIWeather
}

func (m mockAdapter) Weather() (adapter.WeatherOut, error) {
	return m.out, m.outErr
}

var mockAdapterSuccessOut adapter.WeatherOut = adapter.WeatherOut{

	CurrentCondition: []adapter.CurrentCondition{
		{
			TempC: "20",
			TempF: "60",
		},
	},
	Weather: []adapter.Weather{
		{
			AvgtempC: "21",
			Date:     "01-01-2022",
			Hourly: []adapter.Hourly{
				{
					TempC:    "10",
					TempF:    "55",
					Humidity: "40%",
					Time:     "2100",
				},
				{
					TempC:    "10",
					TempF:    "55",
					Humidity: "40%",
					Time:     "2100",
				},
				{
					TempC:    "10",
					TempF:    "55",
					Humidity: "40%",
					Time:     "2100",
				},
			},
		},
	},
}

var mockAdapterSuccessOut2 adapter.WeatherOut = adapter.WeatherOut{

	CurrentCondition: []adapter.CurrentCondition{
		{
			TempC: "20",
			TempF: "60",
		},
	},
	Weather: []adapter.Weather{
		{
			AvgtempC: "11",
			Date:     "02-01-2022",
			Hourly: []adapter.Hourly{
				{
					TempC:    "10",
					TempF:    "55",
					Humidity: "40%",
					Time:     "1200",
				},
				{
					TempC:    "10",
					TempF:    "55",
					Humidity: "40%",
					Time:     "900",
				},
			},
		},
		{
			AvgtempC: "12",
			Date:     "02-02-2022",
			Hourly: []adapter.Hourly{
				{
					TempC:    "10",
					TempF:    "55",
					Humidity: "40%",
					Time:     "2100",
				},
				{
					TempC:    "10",
					TempF:    "55",
					Humidity: "40%",
					Time:     "1200",
				},
			},
		},
	},
}
