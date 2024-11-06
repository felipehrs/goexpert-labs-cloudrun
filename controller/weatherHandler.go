package controller

import (
	"encoding/json"
	"net/http"

	"github.com/felipehrs/go-expert-cloud-run/dto"
	"github.com/felipehrs/go-expert-cloud-run/entity"
	apiErrors "github.com/felipehrs/go-expert-cloud-run/errors"
	service "github.com/felipehrs/go-expert-cloud-run/service"
)

func GetWeatherHandler(w http.ResponseWriter, r *http.Request, fetchAddress func(string) (*http.Response, error), fetchWeather func(string) (*http.Response, error)) {
	zipCode := r.URL.Query().Get("cep")
	if zipCode == "" {
		http.Error(w, apiErrors.InvalidZipCode.Error(), http.StatusUnprocessableEntity)
		return
	}

	if !entity.IsValidZipCode(zipCode) {
		http.Error(w, apiErrors.InvalidZipCode.Error(), http.StatusUnprocessableEntity)
		return
	}

	address, err := service.SearchAddress(zipCode, fetchAddress)
	if err != nil {
		http.Error(w, "Error fetching address: "+err.Error(), http.StatusInternalServerError)
		return
	}

	weather, err := service.SearchWeather(address.Localidade, fetchWeather)
	if err != nil {
		http.Error(w, "Error fetching weather: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.TemperatureResponse{
		TempC: weather.TempCelsius,
		TempF: weather.TempFahrenheit,
		TempK: weather.TempKelvin,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
