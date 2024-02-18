package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"weatherapp/weatherdata"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func main() {

	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	addr := config.Server.Host + ":" + config.Server.Port
	fmt.Printf("Starting server at: %v", addr)

	r := gin.Default()

	r.SetTrustedProxies([]string{"config.Server.Host"})
	r.GET("/weather", func(c *gin.Context) {
		data := getRandomWeather()

		c.JSON(http.StatusOK, data)
	})

	log.Fatal(r.Run(addr)) // host:port
}

func getConfig() (*Config, error) {
	config := &Config{}

	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func getRandomWeather() weatherdata.WeatherData {

	randomTempCelsius := (rand.Int31n(50-(-20)) - 20)
	tempFahrenheit := int32(float32(randomTempCelsius)*1.8) + 32

	summaries := [...]string{"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}
	summariesCount := int32(len(summaries))
	randomPicker := rand.Int31n(summariesCount-1) + 1
	randomSummary := summaries[randomPicker]

	responseData := weatherdata.WeatherData{Date: time.Now().Format("1 Jan 2001"),
		TempCelsius:    randomTempCelsius,
		TempFahrenheit: tempFahrenheit,
		Summary:        randomSummary}

	return responseData
}
