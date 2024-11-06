package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/felipehrs/go-expert-cloud-run/entity"
)

func SearchWeather(city string, fetch func(string) (*http.Response, error)) (entity.Weather, error) {
	apiKey := "5903bf6c2eb54a7f863170134240311" // Replace with your weather API key
	//url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)
	resp, err := fetch(url)
	if err != nil {
		return entity.Weather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.Weather{}, fmt.Errorf("failed to fetch weather, status code: %d", resp.StatusCode)
	}

	var data map[string]interface{}
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		return entity.Weather{}, fmt.Errorf("error unmarshalling weather response: %v", err)
	}

	tempCelsius := data["current"].(map[string]interface{})["temp_c"].(float64)

	return entity.NewWeather(tempCelsius), nil
}
