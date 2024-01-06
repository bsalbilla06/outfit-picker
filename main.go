package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type WeatherResp struct {
	Name string `json:"name"`
	Main struct {
		Low  float64 `json:"temp_min"`
		High float64 `json:"temp_max"`
	} `json:"main"`
}

func main() {
	apiKey := os.Getenv("OPENWEATHERMAP_KEY")
	url := "https://api.openweathermap.org/data/2.5/weather"

	// City latitude and longitude
	longitude := os.Getenv("LONGITUDE")
	latitude := os.Getenv("LATITUDE")

	url = fmt.Sprintf("%s?lon=%s&lat=%s&appid=%s", url, longitude, latitude, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var weatherData WeatherResp
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println(weatherData)
}
