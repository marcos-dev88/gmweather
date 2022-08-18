package adapter

import (
	"encoding/json"
	"io"
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIWeather interface {
	Weather() (WeatherOut, error)
}

func NewConnection(location string) APIWeather {
	return &Connection{Location: location}
}

func (c *Connection) Weather() (out WeatherOut, err error) {

	req, errReq := http.NewRequest("GET", WeatherAPI+c.Location+APIDataFormat, nil)

	if errReq != nil {
		err = errReq
		return
	}

	var resp *http.Response
	resp, errResp := client.Do(req)

	if errResp != nil {
		err = errResp
		return
	}

	body, errBody := io.ReadAll(resp.Body)

	if errBody != nil {
		err = errBody
		return
	}

	if errMarsh := json.Unmarshal(body, &out); errMarsh != nil {
		err = errMarsh
		return
	}

	return out, nil
}
