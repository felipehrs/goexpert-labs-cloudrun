package entity

import "regexp"

type Address struct {
	CEP        string `json:"cep"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Estado     string `json:"estado"`
}

func IsValidZipCode(cep string) bool {
	cepRegex := regexp.MustCompile(`^\d{5}-?\d{3}$`)

	if cepRegex.MatchString(cep) {
		return true
	}

	return false
}
