package repositories

import (
	dtoConversion "game-currency/src/app/dtos/conversion"
	dto "game-currency/src/app/dtos/currency"
	models "game-currency/src/infra/models"
)

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : game-currency
 */

type ConversionRateRepository interface {
	GetConversionRate(fromID, toID int64) (*models.ConversionRates, error)
	GetCurrenciesList() (*[]models.Currencies, error)
	CreateCurrency(data *dto.CurrencyReqDTO) error
	CreateConversionRate(data *dtoConversion.CreateConversReqDTO) error
}
