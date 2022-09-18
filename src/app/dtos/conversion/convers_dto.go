package items_dto

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ConversInterface interface {
	Validate() error
}

type ConversReqDTO struct {
	CurrencyIdFrom int64   `json:"currencyIdFrom"`
	CurrencyIdTo   int64   `json:"currencyIdTo"`
	Amount         float64 `json:"amount"`
}

type ConversRespDTO struct {
	Result float64 `json:"result"`
}

type CreateConversReqDTO struct {
	CurrencyIdFrom int64   `json:"currencyIdFrom"`
	CurrencyIdTo   int64   `json:"currencyIdTo"`
	Rate           float64 `json:"rate"`
}

func (dto *CreateConversReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.CurrencyIdFrom, validation.Required),
		validation.Field(&dto.CurrencyIdTo, validation.Required),
		validation.Field(&dto.Rate, validation.Required),
	); err != nil {
		return err
	}

	if dto.CurrencyIdFrom == dto.CurrencyIdTo {
		err := errors.New("ID From and ID TO cannot be the same")
		return err
	}
	return nil
}
