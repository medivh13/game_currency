package conversion_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : game-currency
 */

import (
	"errors"
	dto "game-currency/src/app/dtos/conversion"
	"game-currency/src/domain/repositories"
	"log"
)

type ConversionUsecaseInterface interface {
	ConversResult(data *dto.ConversReqDTO) (*dto.ConversRespDTO, error)
	CreateConversionRate(data *dto.CreateConversReqDTO) error
}

type conversUseCase struct {
	ConversionRepo repositories.ConversionRateRepository
}

func NewConversionUseCase(conversionRepo repositories.ConversionRateRepository) *conversUseCase {
	return &conversUseCase{
		ConversionRepo: conversionRepo,
	}
}

func (uc *conversUseCase) ConversResult(data *dto.ConversReqDTO) (*dto.ConversRespDTO, error) {
	var resp dto.ConversRespDTO

	dataConversionRate, err := uc.ConversionRepo.GetConversionRate(data.CurrencyIdFrom, data.CurrencyIdTo)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if data.CurrencyIdFrom == dataConversionRate.FromID && data.CurrencyIdTo == dataConversionRate.ToID {
		resp.Result = data.Amount * dataConversionRate.Rate
	} else {
		resp.Result = data.Amount / dataConversionRate.Rate
	}

	return &resp, nil
}

func (uc *conversUseCase) CreateConversionRate(data *dto.CreateConversReqDTO) error {

	dataConversionRate, err := uc.ConversionRepo.GetConversionRate(data.CurrencyIdFrom, data.CurrencyIdTo)
	if err != nil {
		if err.Error() != "Data Not Found" {
			log.Println(err)
			return err
		}
	}

	if dataConversionRate != nil {
		err := errors.New("Conversion Already Exist")
		log.Println(err)
		return err
	}

	err = uc.ConversionRepo.CreateConversionRate(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
