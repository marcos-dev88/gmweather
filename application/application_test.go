package application

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/marcos-dev88/gmweather/internal/adapter"
	"github.com/marcos-dev88/gmweather/internal/service"
)

func TestNewApp(t *testing.T) {
	tests := []struct {
		name string
		want Application
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_app_RunApp(t *testing.T) {
	type fields struct {
		service service.WeatherData
		adapter adapter.APIWeather
	}
	type args struct {
		in Input
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &app{
				service: tt.fields.service,
				adapter: tt.fields.adapter,
			}
			a.RunApp(tt.args.in)
		})
	}
}

// go test -v -failfast -count=1 -run ^Test_app_GetWeather$
func Test_app_GetWeather(t *testing.T) {
	tests := []struct {
		name    string
		mock    mockApp
		want    WeatherData
		wantErr bool
	}{
		{
			name: "success",
			mock: mockApp{
				adapter: mockAdapter{
					out:    adapter.WeatherOut{},
					outErr: nil,
				},
				service: mockServiceSuccess,
				out:     WeatherData{},
				outErr:  nil,
			},
			want: WeatherData{
				CurrentWeather: service.CurrentWeather{
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
				WeatherPrevision: []service.WeatherPrevision{
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
			},
			wantErr: false,
		},
		{
			name: "error current",
			mock: mockApp{
				adapter: mockAdapter{
					out:    adapter.WeatherOut{},
					outErr: nil,
				},
				service: mockServiceErrorCurrent,
				out:     WeatherData{},
				outErr:  nil,
			},
			want:    WeatherData{},
			wantErr: true,
		},
		{
			name: "error prevision",
			mock: mockApp{
				adapter: mockAdapter{
					out:    adapter.WeatherOut{},
					outErr: nil,
				},
				service: mockServiceErrorPrevision,
				out:     WeatherData{},
				outErr:  nil,
			},
			want:    WeatherData{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &app{
				service: tt.mock.service,
				adapter: tt.mock.adapter,
			}
			got, err := a.GetWeather()
			if (err != nil) != tt.wantErr {
				t.Errorf("app.GetWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			msg := cmp.Diff(got, tt.want)

			if msg != "" {
				t.Errorf("\n%s\n", msg)
			}
		})
	}
}

// go test -v -failfast -count=1 -run ^Test_app_UpdateData$
func Test_app_UpdateData(t *testing.T) {
	type args struct {
		d *WeatherData
	}
	tests := []struct {
		name    string
		mock    mockApp
		args    args
		wantErr bool
		want    *WeatherData
	}{
		{
			name: "success",
			args: args{
				d: &WeatherData{
					CurrentWeather: service.CurrentWeather{
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
					WeatherPrevision: []service.WeatherPrevision(nil),
				},
			},
			mock: mockApp{
				adapter: mockAdapter{
					out:    adapter.WeatherOut{},
					outErr: nil,
				},
				service: mockServiceSuccess,
				out:     WeatherData{},
				outErr:  nil,
			},
			want: &WeatherData{
				CurrentWeather: service.CurrentWeather{
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
				WeatherPrevision: []service.WeatherPrevision(nil),
			},
			wantErr: false,
		},
		{
			name: "error current",
			mock: mockApp{
				adapter: mockAdapter{
					out:    adapter.WeatherOut{},
					outErr: nil,
				},
				service: mockServiceErrorCurrent,
				out:     WeatherData{},
				outErr:  nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error prevision",
			mock: mockApp{
				adapter: mockAdapter{
					out:    adapter.WeatherOut{},
					outErr: nil,
				},
				service: mockServiceErrorPrevision,
				out:     WeatherData{},
				outErr:  nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &app{
				service: tt.mock.service,
				adapter: tt.mock.adapter,
			}
			if err := a.UpdateData(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("app.UpdateData() error = %v, wantErr %v", err, tt.wantErr)
			}

			msg := cmp.Diff(tt.args.d, tt.want)

			if msg != "" {
				t.Errorf("\n%s\n", msg)
			}
		})
	}
}
