package currency_handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mockDTO "game-currency/mocks/app/dtos/currency"
	mockUC "game-currency/mocks/app/use_cases/currency"
	mockResp "game-currency/mocks/interface/rest/response"
	dto "game-currency/src/app/dtos/currency"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockBalanceHandler struct {
	mock.Mock
}

type CurrencyHandlerTest struct {
	suite.Suite
	mockUC      *mockUC.MockCurrencyUseCase
	mockResp    *mockResp.MockResponse
	h           CurrencyHandlerInterface
	w           *httptest.ResponseRecorder
	dtoTest     *dto.CurrencyReqDTO
	dtoTestFail *dto.CurrencyReqDTO
	mockDTO     *mockDTO.MockCurrencyDTO
}

func (suite *CurrencyHandlerTest) SetupTest() {
	suite.mockUC = new(mockUC.MockCurrencyUseCase)
	suite.mockResp = new(mockResp.MockResponse)
	suite.mockDTO = new(mockDTO.MockCurrencyDTO)
	suite.h = NewCurrencyHandler(suite.mockResp, suite.mockUC)
	suite.w = httptest.NewRecorder()
	suite.dtoTest = &dto.CurrencyReqDTO{
		Name: "disena",
	}
}

func (s *CurrencyHandlerTest) TestGetCurrenciesListSuccess() {

	r := httptest.NewRequest("GET", "/currency", nil)

	s.mockUC.Mock.On("GetCurrenciesList").Return(mock.Anything, nil)
	s.mockResp.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.GetCurrencyList).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)
}

func (s *CurrencyHandlerTest) TestGetCurrenciesFail() {

	r := httptest.NewRequest("GET", "/currency", nil)

	s.mockUC.Mock.On("GetCurrenciesList").Return(nil, errors.New(mock.Anything))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetCurrencyList).ServeHTTP(s.w, r)

	s.Equal(500, s.w.Result().StatusCode)
}

func (s *CurrencyHandlerTest) TestCreateCurrenciesSuccess() {
	bodyBytes, _ := json.Marshal(s.dtoTest)
	r := httptest.NewRequest("POST", "/currency", bytes.NewBuffer(bodyBytes))

	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUC.Mock.On("CreateCurrency", s.dtoTest).Return(nil)
	s.mockResp.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.CreateCurrency).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)
}

func (s *CurrencyHandlerTest) TestCreateCurrenciesFailDecode() {
	r := httptest.NewRequest("POST", "/currency/", nil)
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.CreateCurrency).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *CurrencyHandlerTest) TestCreateFailValidateDTO() {
	bodyBytes, _ := json.Marshal(s.dtoTestFail)
	r := httptest.NewRequest("POST", "/currency/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.CreateCurrency).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *CurrencyHandlerTest) TestCreateFail() {
	bodyBytes, _ := json.Marshal(s.dtoTest)
	r := httptest.NewRequest("POST", "/currency/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUC.Mock.On("CreateCurrency", s.dtoTest).Return(errors.New(mock.Anything))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.CreateCurrency).ServeHTTP(s.w, r)

	s.Equal(500, s.w.Result().StatusCode)

}

func TestHandler(t *testing.T) {
	suite.Run(t, new(CurrencyHandlerTest))
}
