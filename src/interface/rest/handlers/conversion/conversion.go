package balance_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : game-currency
 */

import (
	"encoding/json"
	"net/http"
	"strconv"

	dto "game-currency/src/app/dtos/conversion"
	usecases "game-currency/src/app/use_cases/conversion"
	common_error "game-currency/src/infra/errors"
	"game-currency/src/interface/rest/response"

	"github.com/go-chi/chi/v5"
)

type ConversionHandlerInterface interface {
	GetConversionResult(w http.ResponseWriter, r *http.Request)
	CreateConversionRate(w http.ResponseWriter, r *http.Request)
}

type conversHandler struct {
	response response.IResponseClient
	usecase  usecases.ConversionUsecaseInterface
}

func NewConversionHandler(r response.IResponseClient, h usecases.ConversionUsecaseInterface) ConversionHandlerInterface {
	return &conversHandler{
		response: r,
		usecase:  h,
	}
}

func (h *conversHandler) GetConversionResult(w http.ResponseWriter, r *http.Request) {
	getDTO := dto.ConversReqDTO{}
	idFrom, err := strconv.Atoi(chi.URLParam(r, "idFrom"))
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	getDTO.CurrencyIdFrom = int64(idFrom)

	idTo, err := strconv.Atoi(chi.URLParam(r, "idTo"))

	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	getDTO.CurrencyIdTo = int64(idTo)

	ammount, err := strconv.ParseFloat(chi.URLParam(r, "ammount"), 64)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	getDTO.Amount = ammount

	data, err := h.usecase.ConversResult(&getDTO)
	if err != nil {
		if err.Error() == "Data Not Found" {
			h.response.HttpError(w, common_error.NewError(common_error.DATA_NOT_FOUND, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Conversion Result",
		data,
		nil,
	)
}

func (h *conversHandler) CreateConversionRate(w http.ResponseWriter, r *http.Request) {
	postDTO := dto.CreateConversReqDTO{}
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
	err = h.usecase.CreateConversionRate(&postDTO)
	if err != nil {
		if err.Error() == "Conversion Already Exist" {
			h.response.HttpError(w, common_error.NewError(common_error.DATA_EXIST, err))
			return
		}
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
