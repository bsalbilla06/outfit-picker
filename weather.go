package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResp struct {
	Name string `json:"name"`
	Main struct {
		Low  float64 `json:"temp_min"`
		High float64 `json:"temp_max"`
	} `json:"main"`
}

func kelvinToFahrenheit(temp float64) float64 {
	return (temp - 273.15) * 9/5 + 32
}

func RetrieveWeatherData(apiKey, longitude, latitude string) (WeatherResp, error) {
	url := "https://api.openweathermap.org/data/2.5/weather"

	url = fmt.Sprintf("%s?lon=%s&lat=%s&appid=%s", url, longitude, latitude, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResp{}, err
	}
	defer resp.Body.Close()

	var weatherData WeatherResp
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return WeatherResp{}, err
	}
	weatherData.Main.High = kelvinToFahrenheit(weatherData.Main.High)
	weatherData.Main.Low = kelvinToFahrenheit(weatherData.Main.Low)
	return weatherData, nil
}
