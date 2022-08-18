package service

type Period int

const (
	_ Period = 300 * (iota + 1)
	_
	Morning
	Noon
	_
	Evening
	Night
)

type CheckWeatherOut struct {
	CurrentWeather   `json:"current_weather"`
	WeatherPrevision []WeatherPrevision `json:"weather_prevision"`
}

type CurrentWeather struct {
	FeelsLikeC       string        `json:"feels_like_c"`
	FeelsLikeF       string        `json:"feels_like_f"`
	Cloudcover       string        `json:"cloudcover"`
	Humidity         string        `json:"humidity"`
	LocalObsDateTime string        `json:"local_obs_date_time"`
	ObservationTime  string        `json:"observation_time"`
	PrecipInches     string        `json:"precipInches"`
	PrecipMM         string        `json:"precip_mm"`
	Pressure         string        `json:"pressure"`
	PressureInches   string        `json:"pressure_inches"`
	TempC            string        `json:"temp_C"`
	TempF            string        `json:"temp_F"`
	UvIndex          string        `json:"uv_index"`
	Visibility       string        `json:"visibility"`
	VisibilityMiles  string        `json:"visibility_miles"`
	WeatherCode      string        `json:"weather_code"`
	WeatherDesc      []WeatherDesc `json:"weather_desc"`
	WinddirPoint     string        `json:"wind_dir_point"`
	WinddirDegree    string        `json:"wind_dir_degree"`
	WindspeedKmph    string        `json:"wind_speed_Kkmph"`
	WindspeedMiles   string        `json:"wind_speed_miles"`
}

type WeatherDesc struct {
	Value string `json:"value"`
}

type WeatherPrevision struct {
	DateTime      string          `json:"date"`
	TempMaxC      string          `json:"temp_max_c"`
	TempMinC      string          `json:"temp_min_c"`
	TempMaxF      string          `json:"temp_max_f"`
	TempMinF      string          `json:"temp_min_f"`
	WeatherPeriod []WeatherPeriod `json:"weather_period"`
}

type WeatherPeriod struct {
	TimePeriod       string        `json:"time_period"`
	FeelsLikeC       string        `json:"feels_like_c"`
	FeelsLikeF       string        `json:"feels_like_f"`
	Cloudcover       string        `json:"cloudcover"`
	Humidity         string        `json:"humidity"`
	LocalObsDateTime string        `json:"local_obs_date_time"`
	ObservationTime  string        `json:"observation_time"`
	PrecipInches     string        `json:"precipInches"`
	PrecipMM         string        `json:"precip_mm"`
	Pressure         string        `json:"pressure"`
	PressureInches   string        `json:"pressure_inches"`
	TempC            string        `json:"temp_C"`
	TempF            string        `json:"temp_F"`
	UvIndex          string        `json:"uv_index"`
	Visibility       string        `json:"visibility"`
	VisibilityMiles  string        `json:"visibility_miles"`
	WeatherCode      string        `json:"weather_code"`
	WeatherDesc      []WeatherDesc `json:"weather_desc"`
	WinddirPoint     string        `json:"wind_dir_point"`
	WinddirDegree    string        `json:"wind_dir_degree"`
	WindspeedKmph    string        `json:"wind_speed_Kkmph"`
	WindspeedMiles   string        `json:"wind_speed_miles"`
}
