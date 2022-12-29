package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"weatherapp/weatherdata"

	"gopkg.in/yaml.v3"
)

func main() {

	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	addr := config.Server.Host + ":" + config.Server.Port
	fmt.Printf("Starting server at: %v", addr)

	http.HandleFunc("/getweather", getRandomWeather)

	log.Fatal(http.ListenAndServe(addr, nil))
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

func getRandomWeather(w http.ResponseWriter, request *http.Request) {

	summaries := [...]string{"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}
	summariesCount := int32(len(summaries))

	randomTemp := (rand.Int31n(50-(-20)) - 20)

	randomPicker := rand.Int31n(summariesCount-1) + 1
	randomSummary := summaries[randomPicker]

	responseData := weatherdata.WeatherData{Date: time.Now().Format("2 Jan 2006"), TempCelsius: randomTemp, Summary: randomSummary}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseData)
}
