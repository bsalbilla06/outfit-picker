package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Printf("%s\n", os.Getenv("OPENWEATHERMAP_KEY"))
}
