package usecases

import (
	conversUC "game-currency/src/app/use_cases/conversion"
	currencyUC "game-currency/src/app/use_cases/currency"
)

type AllUseCases struct {
	ConversUseCase  conversUC.ConversionUsecaseInterface
	CurrencyUseCase currencyUC.CurrencyUsecaseInterface
}
