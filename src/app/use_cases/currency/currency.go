package currency_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : game-currency
 */

import (
	dto "game-currency/src/app/dtos/currency"
	"game-currency/src/domain/repositories"
	"log"
)

type CurrencyUsecaseInterface interface {
	GetCurrenciesList() ([]*dto.CurrencyRespDTO, error)
	CreateCurrency(data *dto.CurrencyReqDTO) error
}

type currencyUseCase struct {
	CurrencyRepo repositories.ConversionRateRepository
}

func NewCurrencyUseCase(currencyRepo repositories.ConversionRateRepository) *currencyUseCase {
	return &currencyUseCase{
		CurrencyRepo: currencyRepo,
	}
}

func (uc *currencyUseCase) GetCurrenciesList() ([]*dto.CurrencyRespDTO, error) {

	dataCurrencies, err := uc.CurrencyRepo.GetCurrenciesList()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToCurrencies(dataCurrencies), nil
}

func (uc *currencyUseCase) CreateCurrency(data *dto.CurrencyReqDTO) error {

	err := uc.CurrencyRepo.CreateCurrency(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
