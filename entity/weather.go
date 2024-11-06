package entity

type Weather struct {
	TempCelsius    float64 `json:"temp_celsius"`
	TempFahrenheit float64 `json:"temp_fahrenheit"`
	TempKelvin     float64 `json:"temp_kelvin"`
}

func NewWeather(tempCelsius float64) Weather {
	tempFahrenheit := CelsiusToFahrenheit(tempCelsius)
	tempKelvin := CelsiusToKelvin(tempCelsius)

	return Weather{
		TempCelsius:    tempCelsius,
		TempFahrenheit: tempFahrenheit,
		TempKelvin:     tempKelvin,
	}
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}
