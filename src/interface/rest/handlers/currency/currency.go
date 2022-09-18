package currency_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : game-currency
 */

import (
	"encoding/json"
	"net/http"

	dto "game-currency/src/app/dtos/currency"
	usecases "game-currency/src/app/use_cases/currency"
	common_error "game-currency/src/infra/errors"
	"game-currency/src/interface/rest/response"
)

type CurrencyHandlerInterface interface {
	GetCurrencyList(w http.ResponseWriter, r *http.Request)
	CreateCurrency(w http.ResponseWriter, r *http.Request)
}

type currencyHandler struct {
	response response.IResponseClient
	usecase  usecases.CurrencyUsecaseInterface
}

func NewCurrencyHandler(r response.IResponseClient, h usecases.CurrencyUsecaseInterface) CurrencyHandlerInterface {
	return &currencyHandler{
		response: r,
		usecase:  h,
	}
}

func (h *currencyHandler) GetCurrencyList(w http.ResponseWriter, r *http.Request) {

	data, err := h.usecase.GetCurrenciesList()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Currencies",
		data,
		nil,
	)
}

func (h *currencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	postDTO := dto.CurrencyReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = h.usecase.CreateCurrency(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Create Currency",
		nil,
		nil,
	)
}
