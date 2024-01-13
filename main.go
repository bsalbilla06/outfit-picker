package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	weatherData, err := RetrieveWeatherData(os.Getenv("OPENWEATHERMAP_KEY"), os.Getenv("LONGITUDE"), os.Getenv("LATITUDE"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(weatherData)
}
