package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// Slices of temperature ranges where user will wear certain items of clothing
type TemperaturePreferences struct {
	Boots       [2]int
	Sandals     [2]int
	Sneakers    [2]int
	Socks       [2]int
	Sweatpants  [2]int
	Jeans       [2]int
	Slacks      [2]int
	Shorts      [2]int
	Skirts      [2]int
	TankTops    [2]int
	TShirts     [2]int
	LongSleeves [2]int
	CrewNecks   [2]int
	Hoodies     [2]int
	Jackets     [2]int
}

func main() {
	weatherData, err := RetrieveWeatherData(os.Getenv("OPENWEATHERMAP_KEY"), os.Getenv("LONGITUDE"), os.Getenv("LATITUDE"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(weatherData)
}
