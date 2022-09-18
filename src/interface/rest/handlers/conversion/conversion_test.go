package balance_handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mockDTO "game-currency/mocks/app/dtos/conversion"

	mockUC "game-currency/mocks/app/use_cases/conversion"
	mockResp "game-currency/mocks/interface/rest/response"
	dto "game-currency/src/app/dtos/conversion"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockConversionHandler struct {
	mock.Mock
}

type ConversionHandlerTest struct {
	suite.Suite
	mockUC            *mockUC.MockConversionUseCase
	mockResp          *mockResp.MockResponse
	h                 ConversionHandlerInterface
	w                 *httptest.ResponseRecorder
	dtoTest           *dto.ConversReqDTO
	dtoTestFail       *dto.ConversReqDTO
	mockDTO           *mockDTO.MockConversionDTO
	dtoCreateTset     *dto.CreateConversReqDTO
	dtoCreateTestFail *dto.CreateConversReqDTO
}

func (suite *ConversionHandlerTest) SetupTest() {
	suite.mockUC = new(mockUC.MockConversionUseCase)
	suite.mockResp = new(mockResp.MockResponse)
	suite.mockDTO = new(mockDTO.MockConversionDTO)
	suite.h = NewConversionHandler(suite.mockResp, suite.mockUC)
	suite.w = httptest.NewRecorder()
	suite.dtoTest = &dto.ConversReqDTO{
		CurrencyIdFrom: 1,
		CurrencyIdTo:   2,
		Amount:         580,
	}
	suite.dtoTestFail = &dto.ConversReqDTO{
		CurrencyIdFrom: 5,
		CurrencyIdTo:   2,
		Amount:         580,
	}

	suite.dtoCreateTset = &dto.CreateConversReqDTO{
		CurrencyIdFrom: 3,
		CurrencyIdTo:   2,
		Rate:           29,
	}
}

func (s *ConversionHandlerTest) TestConversionSuccess() {

	r := httptest.NewRequest("GET", "/conversion/1/2/580", nil)

	idFrom := "1"
	idTo := "2"
	ammount := "580"
	ctx := r.Context()
	r = r.WithContext(ctx)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("idFrom", idFrom)
	rctx.URLParams.Add("idTo", idTo)
	rctx.URLParams.Add("ammount", ammount)

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUC.Mock.On("ConversResult", s.dtoTest).Return(mock.Anything, nil)
	s.mockResp.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.GetConversionResult).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)
}

func (s *ConversionHandlerTest) TestConversionFailConvertPayloadIdFrom() {

	r := httptest.NewRequest("GET", "/conversion/A/2/580", nil)

	idFrom := "A"
	idTo := "2"
	ammount := "580"
	ctx := r.Context()
	r = r.WithContext(ctx)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("idFrom", idFrom)
	rctx.URLParams.Add("idTo", idTo)
	rctx.URLParams.Add("ammount", ammount)

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetConversionResult).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)
}

func (s *ConversionHandlerTest) TestConversionFailConvertPayloadIdTo() {
	r := httptest.NewRequest("GET", "/conversion/1/A/580", nil)
	idFrom := "1"
	idTo := "A"
	ammount := "580"
	ctx := r.Context()
	r = r.WithContext(ctx)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("idFrom", idFrom)
	rctx.URLParams.Add("idTo", idTo)
	rctx.URLParams.Add("ammount", ammount)

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetConversionResult).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)
}

func (s *ConversionHandlerTest) TestConversionFailConvertPayloadAmmount() {

	r := httptest.NewRequest("GET", "/conversion/1/2/ABC", nil)

	idFrom := "1"
	idTo := "2"
	ammount := "ABC"
	ctx := r.Context()
	r = r.WithContext(ctx)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("idFrom", idFrom)
	rctx.URLParams.Add("idTo", idTo)
	rctx.URLParams.Add("ammount", ammount)

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetConversionResult).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)
}

func (s *ConversionHandlerTest) TestConversionNotFound() {

	r := httptest.NewRequest("GET", "/conversion/5/2/580", nil)

	idFrom := "5"
	idTo := "2"
	ammount := "580"
	ctx := r.Context()
	r = r.WithContext(ctx)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("idFrom", idFrom)
	rctx.URLParams.Add("idTo", idTo)
	rctx.URLParams.Add("ammount", ammount)

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUC.Mock.On("ConversResult", s.dtoTestFail).Return(nil, errors.New("Data Not Found"))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetConversionResult).ServeHTTP(s.w, r)

	s.Equal(404, s.w.Result().StatusCode)
}
func (s *ConversionHandlerTest) TestConversionFail() {

	r := httptest.NewRequest("GET", "/conversion/1/2/580", nil)

	idFrom := "1"
	idTo := "2"
	ammount := "580"
	ctx := r.Context()
	r = r.WithContext(ctx)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("idFrom", idFrom)
	rctx.URLParams.Add("idTo", idTo)
	rctx.URLParams.Add("ammount", ammount)

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUC.Mock.On("ConversResult", s.dtoTest).Return(nil, errors.New(mock.Anything))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetConversionResult).ServeHTTP(s.w, r)

	s.Equal(500, s.w.Result().StatusCode)
}

func (s *ConversionHandlerTest) TestCreateConversionSuccess() {
	bodyBytes, _ := json.Marshal(s.dtoCreateTset)
	r := httptest.NewRequest("POST", "/conversion", bytes.NewBuffer(bodyBytes))

	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUC.Mock.On("CreateConversionRate", s.dtoCreateTset).Return(nil)
	s.mockResp.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.CreateConversionRate).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)
}

func (s *ConversionHandlerTest) TestCreateConversionFailDecode() {
	r := httptest.NewRequest("POST", "/conversion/", nil)
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.CreateConversionRate).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *ConversionHandlerTest) TestCreateConversionFailValidateDTO() {
	bodyBytes, _ := json.Marshal(s.dtoCreateTestFail)
	r := httptest.NewRequest("POST", "/conversion/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.CreateConversionRate).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *ConversionHandlerTest) TestCreateConversionFailDataExist() {
	bodyBytes, _ := json.Marshal(s.dtoCreateTset)
	r := httptest.NewRequest("POST", "/conversion", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUC.Mock.On("CreateConversionRate", s.dtoCreateTset).Return(errors.New("Conversion Already Exist"))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.CreateConversionRate).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *ConversionHandlerTest) TestCreateConversionFail() {
	bodyBytes, _ := json.Marshal(s.dtoCreateTset)
	r := httptest.NewRequest("POST", "/conversion", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUC.Mock.On("CreateConversionRate", s.dtoCreateTset).Return(errors.New(mock.Anything))
	s.mockResp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.CreateConversionRate).ServeHTTP(s.w, r)

	s.Equal(500, s.w.Result().StatusCode)

}

func TestHandler(t *testing.T) {
	suite.Run(t, new(ConversionHandlerTest))
}
