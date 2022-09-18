package balance_dto

import validation "github.com/go-ozzo/ozzo-validation"

type CurrencyDTOInterface interface {
	Validate() error
}

type CurrencyReqDTO struct {
	Name string `json:"name"`
}

func (dto *CurrencyReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Name, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type CurrencyRespDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
