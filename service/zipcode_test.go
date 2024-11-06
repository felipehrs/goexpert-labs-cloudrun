package service_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/felipehrs/go-expert-cloud-run/service"
)

func TestSearchAddressUnmarshalError(t *testing.T) {
	mockFetchAddress := func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("invalid json")),
		}, nil
	}

	_, err := SearchAddress("01001000", mockFetchAddress)
	assert.Error(t, err)
}

func TestSearchAddressFechError(t *testing.T) {
	mockFetchAddress := func(url string) (*http.Response, error) {
		addressJSON := `{"cep": "01001-000", "localidade": "São Paulo", "uf": "SP", "estado": "São Paulo"}`
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(addressJSON)),
		}, nil
	}

	address, err := SearchAddress("01001000", mockFetchAddress)
	assert.Nil(t, err)
	assert.Equal(t, address.CEP, "01001-000")
	assert.Equal(t, address.UF, "SP")
	assert.Equal(t, address.Estado, "São Paulo")
}

func TestSearchAddressFech(t *testing.T) {
	mockFetchAddress := func(url string) (*http.Response, error) {
		return nil, errors.New("error fetching data")
	}

	_, err := SearchAddress("01001000", mockFetchAddress)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error fetching data")
}
