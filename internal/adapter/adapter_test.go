package adapter

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"testing"
)

// go test -count=1 -failfast -v -run ^TestSome
func TestSome(t *testing.T) {
	var out WeatherOut
	req, errReq := http.NewRequest("GET", WeatherAPI+"sao_paulo"+APIDataFormat, nil)

	if errReq != nil {
		t.Errorf("error -> %v", errReq)
	}

	var resp *http.Response
	resp, errResp := client.Do(req)

	if errResp != nil {
		t.Errorf("error -> %v", errResp)
	}

	body, errBody := io.ReadAll(resp.Body)

	if errBody != nil {
		t.Errorf("error -> %v", errBody)
	}

	if errMarsh := json.Unmarshal(body, &out); errMarsh != nil {
		t.Errorf("error -> %v", errMarsh)
	}

	log.Printf("out -> %+v", out)
}

func TestNewConnection(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name string
		args args
		want APIWeather
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConnection(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_Weather(t *testing.T) {
	type fields struct {
		Location string
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut []Weather
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connection{
				Location: tt.fields.Location,
			}
			gotOut, err := c.Weather()
			if (err != nil) != tt.wantErr {
				t.Errorf("Connection.Weather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Connection.Weather() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
