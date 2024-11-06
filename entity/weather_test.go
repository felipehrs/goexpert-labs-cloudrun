package entity_test

import (
	"testing"

	"github.com/felipehrs/go-expert-cloud-run/entity"
)

func TestNewWeather(t *testing.T) {
	tests := []struct {
		name                   string
		tempCelsius            float64
		expectedTempFahrenheit float64
		expectedTempKelvin     float64
	}{
		{"Normal temperature", 25, 77, 298},
		{"Freezing point", 0, 32, 273},
		{"Boiling point", 100, 212, 373},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := entity.NewWeather(tt.tempCelsius)
			if w.TempFahrenheit != tt.expectedTempFahrenheit {
				t.Errorf("NewWeather() TempFahrenheit = %v, expected %v", w.TempFahrenheit, tt.expectedTempFahrenheit)
			}
			if w.TempKelvin != tt.expectedTempKelvin {
				t.Errorf("NewWeather() TempKelvin = %v, expected %v", w.TempKelvin, tt.expectedTempKelvin)
			}
		})
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name                   string
		tempCelsius            float64
		expectedTempFahrenheit float64
	}{
		{"Normal temperature", 25, 77},
		{"Freezing point", 0, 32},
		{"Boiling point", 100, 212},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := entity.CelsiusToFahrenheit(tt.tempCelsius)
			if result != tt.expectedTempFahrenheit {
				t.Errorf("CelsiusToFahrenheit() = %v, expected %v", result, tt.expectedTempFahrenheit)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name               string
		tempCelsius        float64
		expectedTempKelvin float64
	}{
		{"Normal temperature", 25, 298},
		{"Freezing point", 0, 273},
		{"Boiling point", 100, 373},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := entity.CelsiusToKelvin(tt.tempCelsius)
			if result != tt.expectedTempKelvin {
				t.Errorf("CelsiusToKelvin() = %v, expected %v", result, tt.expectedTempKelvin)
			}
		})
	}
}
