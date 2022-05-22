package service

import (
	"reflect"
	"testing"

	"github.com/marcos-dev88/gmweather/gmweather/adapter"
)

// go test -v -failfast -run ^TestNewWeatherService
func TestNewWeatherService(t *testing.T) {

	adapt := mockAdapter{
		out:    mockAdapterSuccessOut,
		outErr: nil,
	}

	tests := []struct {
		name    string
		adapter mockAdapter
		want    WeatherData
	}{
		{
			name:    "success",
			adapter: adapt,
			want:    NewWeatherService(adapt),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeatherService(tt.adapter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeatherService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertPeriod(t *testing.T) {
	type args struct {
		period int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertPeriod(tt.args.period); got != tt.want {
				t.Errorf("convertPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetCurrent(t *testing.T) {
	type fields struct {
		Adapter adapter.APIWeather
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut CurrentWeather
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				Adapter: tt.fields.Adapter,
			}
			gotOut, err := s.GetCurrent()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetCurrent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("service.GetCurrent() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_service_GetPrevision(t *testing.T) {
	type fields struct {
		Adapter adapter.APIWeather
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut []WeatherPrevision
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				Adapter: tt.fields.Adapter,
			}
			gotOut, err := s.GetPrevision()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetPrevision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("service.GetPrevision() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

// go test -count=1 -failfast -v -run ^Test_convert
func Test_convert(t *testing.T) {
	type args struct {
		weatherData adapter.WeatherOut
	}
	tests := []struct {
		name    string
		args    args
		wantOut CheckWeatherOut
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				weatherData: mockAdapterSuccessOut,
			},
			wantOut: CheckWeatherOut{
				CurrentWeather: CurrentWeather{
					TempC: "20",
					TempF: "60",
				},
				WeatherPrevision: []WeatherPrevision{
					{
						DateTime: "01-01-2022",
						WeatherPeriod: []WeatherPeriod{
							{
								TempC:      "10",
								TempF:      "55",
								Humidity:   "40%",
								TimePeriod: "Night",
							},
							{
								TempC:      "10",
								TempF:      "55",
								Humidity:   "40%",
								TimePeriod: "Night",
							},
							{
								TempC:      "10",
								TempF:      "55",
								Humidity:   "40%",
								TimePeriod: "Night",
							},
						},
					},
				},
			},
		},
		{
			name: "success_2",
			args: args{
				weatherData: mockAdapterSuccessOut2,
			},
			wantOut: CheckWeatherOut{
				CurrentWeather: CurrentWeather{
					TempC: "20",
					TempF: "60",
				},
				WeatherPrevision: []WeatherPrevision{
					{
						DateTime: "02-01-2022",
						WeatherPeriod: []WeatherPeriod{
							{
								TempC:      "10",
								TempF:      "55",
								Humidity:   "40%",
								TimePeriod: "Noon",
							},
							{
								TempC:      "10",
								TempF:      "55",
								Humidity:   "40%",
								TimePeriod: "Morning",
							},
						},
					},
					{
						DateTime: "02-02-2022",
						WeatherPeriod: []WeatherPeriod{
							{
								TempC:      "10",
								TempF:      "55",
								Humidity:   "40%",
								TimePeriod: "Night",
							},
							{
								TempC:      "10",
								TempF:      "55",
								Humidity:   "40%",
								TimePeriod: "Noon",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := convert(tt.args.weatherData)
			if (err != nil) != tt.wantErr {
				t.Errorf("convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("convert() = %+v, want %+v", gotOut, tt.wantOut)
			}
		})
	}
}
