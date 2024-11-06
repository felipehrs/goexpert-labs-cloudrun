package service_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/felipehrs/go-expert-cloud-run/service"
)

func TestSearchWeather(t *testing.T) {
	mockResponse := `{"current":{"temp_c":25,"aqi":"AQI1"}}`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	weather, err := SearchWeather("Sao Paulo", func(url string) (*http.Response, error) {
		return mockServer.Client().Get(mockServer.URL)
	})

	require.NoError(t, err)
	assert.Equal(t, 25.0, weather.TempCelsius)
}

func TestSearchWeatherInvalidRequest(t *testing.T) {
	_, err := SearchWeather("Sao Paulo", func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(bytes.NewBufferString("Error")),
		}, nil
	})

	require.Error(t, err)
	require.Equal(t, err.Error(), "failed to fetch weather, status code: 400")
}

func TestSearchWeatherUnmarshalError(t *testing.T) {
	_, err := SearchWeather("Sao Paulo", func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("Invalid JSON")),
		}, nil
	})

	require.Error(t, err)
	require.Contains(t, err.Error(), "error unmarshalling weather response:")
}

func TestSearchWeatherFetchError(t *testing.T) {
	_, err := SearchWeather("Sao Paulo", func(url string) (*http.Response, error) {
		return nil, errors.New("error fetching data")
	})

	require.Error(t, err)
	require.Contains(t, err.Error(), "error fetching data")
}
