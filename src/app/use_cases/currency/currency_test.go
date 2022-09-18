package currency_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : game-currency
 */

import (
	"errors"
	"testing"

	mockDTO "game-currency/mocks/app/dtos/currency"
	mockRepo "game-currency/mocks/domain/repositories"

	dto "game-currency/src/app/dtos/currency"
	models "game-currency/src/infra/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseCurrencyTest struct {
	suite.Suite
	repo *mockRepo.MockConversionRepo

	usecase     CurrencyUsecaseInterface
	models      *[]models.Currencies
	dtoTest     *dto.CurrencyReqDTO
	mockDTO     *mockDTO.MockCurrencyDTO
}

func (suite *UsecaseCurrencyTest) SetupTest() {
	suite.repo = new(mockRepo.MockConversionRepo)
	suite.mockDTO = new(mockDTO.MockCurrencyDTO)
	suite.usecase = NewCurrencyUseCase(suite.repo)

	suite.models = &[]models.Currencies{
		models.Currencies{
			ID:   1,
			Name: "sickle",
		},
	}

	suite.dtoTest = &dto.CurrencyReqDTO{
		Name: "sickle",
	}

}

func (uc *UsecaseCurrencyTest) TestGetCurrenciesList() {
	uc.repo.Mock.On("GetCurrenciesList").Return(uc.models, nil)
	_, err := uc.usecase.GetCurrenciesList()
	uc.Equal(nil, err)
}

func (uc *UsecaseCurrencyTest) TestGetCurrenciesListFail() {
	uc.repo.Mock.On("GetCurrenciesList").Return(nil, errors.New(mock.Anything))

	_, err := uc.usecase.GetCurrenciesList()
	uc.Error(errors.New(mock.Anything), err)
}

func (uc *UsecaseCurrencyTest) TestCreateCurrency() {
	uc.repo.Mock.On("CreateCurrency",uc.dtoTest).Return(nil)
	err := uc.usecase.CreateCurrency(uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseCurrencyTest) TestCreateCurrencyFail() {
	uc.repo.Mock.On("CreateCurrency",uc.dtoTest).Return(errors.New(mock.Anything))
	err := uc.usecase.CreateCurrency(uc.dtoTest)
	uc.Error(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseCurrencyTest))
}
