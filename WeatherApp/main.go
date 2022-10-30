package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	"weatherapp/weatherdata"
)

func main() {
	http.HandleFunc("/hello", getWeather)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func getWeather(w http.ResponseWriter, req *http.Request) {

	summaries := [...]string{"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}

	randomTemp := (rand.Int31n(50-(-20)) - 20)

	randomPicker := rand.Int31n(9-1) + 1
	randomSummary := summaries[randomPicker]

	responseData := weatherdata.WeatherData{Date: time.Now().Format("2 Jan 2006"), TempCelsius: randomTemp, Summary: randomSummary}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseData)
}
