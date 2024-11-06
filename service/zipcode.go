package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/felipehrs/go-expert-cloud-run/entity"
)

func SearchAddress(cep string, fetch func(string) (*http.Response, error)) (entity.Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := fetch(url)
	if err != nil {
		return entity.Address{}, err
	}
	defer resp.Body.Close()

	var address entity.Address
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &address); err != nil {
		return entity.Address{}, err
	}

	return address, nil
}
