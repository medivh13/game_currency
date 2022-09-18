package balance_dto

import (
	models "game-currency/src/infra/models"
)

func ToGetCurrency(d *models.Currencies) *CurrencyRespDTO {
	return &CurrencyRespDTO{
		ID:   d.ID,
		Name: d.Name,
	}
}

func ToCurrencies(d *[]models.Currencies) []*CurrencyRespDTO {
	var data []*CurrencyRespDTO
	for _, val := range *d {
		data = append(data, ToGetCurrency(&val))
	}
	return data
}
