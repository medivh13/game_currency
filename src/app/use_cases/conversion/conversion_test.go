package conversion_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : game-currency
 */

import (
	"errors"
	"testing"

	mockDTO "game-currency/mocks/app/dtos/conversion"
	mockRepo "game-currency/mocks/domain/repositories"

	dto "game-currency/src/app/dtos/conversion"
	models "game-currency/src/infra/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseConversionTest struct {
	suite.Suite
	repo *mockRepo.MockConversionRepo

	usecase       ConversionUsecaseInterface
	models        *models.ConversionRates
	dtoTest       *dto.ConversReqDTO
	dtoTest2      *dto.ConversReqDTO
	dtoTestFail   *dto.ConversReqDTO
	dtoTestCreate *dto.CreateConversReqDTO
	mockDTO       *mockDTO.MockConversionDTO
}

func (suite *UsecaseConversionTest) SetupTest() {
	suite.repo = new(mockRepo.MockConversionRepo)
	suite.mockDTO = new(mockDTO.MockConversionDTO)
	suite.usecase = NewConversionUseCase(suite.repo)

	suite.models = &models.ConversionRates{
		ID:     1,
		FromID: 1,
		ToID:   2,
		Rate:   29,
	}

	suite.dtoTest = &dto.ConversReqDTO{
		CurrencyIdFrom: 2,
		CurrencyIdTo:   1,
		Amount:         580,
	}

	suite.dtoTest2 = &dto.ConversReqDTO{
		CurrencyIdFrom: 1,
		CurrencyIdTo:   2,
		Amount:         580,
	}

	suite.dtoTestFail = &dto.ConversReqDTO{
		CurrencyIdFrom: 1,
		CurrencyIdTo:   1,
		Amount:         580,
	}

	suite.dtoTestCreate = &dto.CreateConversReqDTO{
		CurrencyIdFrom: 1,
		CurrencyIdTo:   2,
		Rate:           20,
	}
}

func (uc *UsecaseConversionTest) TestGetConversionSuccess() {
	uc.repo.Mock.On("GetConversionRate", uc.dtoTest.CurrencyIdFrom, uc.dtoTest.CurrencyIdTo).Return(uc.models, nil)
	_, err := uc.usecase.ConversResult(uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseConversionTest) TestGetConversionSuccess2() {
	uc.repo.Mock.On("GetConversionRate", uc.dtoTest2.CurrencyIdFrom, uc.dtoTest2.CurrencyIdTo).Return(uc.models, nil)
	_, err := uc.usecase.ConversResult(uc.dtoTest2)
	uc.Equal(nil, err)
}

func (uc *UsecaseConversionTest) TestGetConversionFail() {
	uc.repo.Mock.On("GetConversionRate", uc.dtoTestFail.CurrencyIdFrom, uc.dtoTestFail.CurrencyIdTo).Return(nil, errors.New(mock.Anything))
	_, err := uc.usecase.ConversResult(uc.dtoTestFail)
	uc.Error(errors.New(mock.Anything), err)
}

func (uc *UsecaseConversionTest) TestCreateConversionSuccess() {
	uc.repo.Mock.On("GetConversionRate", uc.dtoTestCreate.CurrencyIdFrom, uc.dtoTestCreate.CurrencyIdTo).Return(nil, nil)
	uc.repo.Mock.On("CreateConversionRate", uc.dtoTestCreate).Return(nil)

	err := uc.usecase.CreateConversionRate(uc.dtoTestCreate)
	uc.Equal(nil, err)
}

func (uc *UsecaseConversionTest) TestCreateConversionFailErroRetriveExistinConversion() {
	uc.repo.Mock.On("GetConversionRate", uc.dtoTestCreate.CurrencyIdFrom, uc.dtoTestCreate.CurrencyIdTo).Return(nil, errors.New(mock.Anything))

	err := uc.usecase.CreateConversionRate(uc.dtoTestCreate)
	uc.Error(errors.New(mock.Anything), err)
}

func (uc *UsecaseConversionTest) TestCreateConversionFailConversionExist() {
	uc.repo.Mock.On("GetConversionRate", uc.dtoTestCreate.CurrencyIdFrom, uc.dtoTestCreate.CurrencyIdTo).Return(uc.models, nil)

	err := uc.usecase.CreateConversionRate(uc.dtoTestCreate)
	uc.Error(errors.New(mock.Anything), err)
}

func (uc *UsecaseConversionTest) TestCreateConversionFail() {
	uc.repo.Mock.On("GetConversionRate", uc.dtoTestCreate.CurrencyIdFrom, uc.dtoTestCreate.CurrencyIdTo).Return(nil, nil)
	uc.repo.Mock.On("CreateConversionRate", uc.dtoTestCreate).Return(errors.New(mock.Anything))
	err := uc.usecase.CreateConversionRate(uc.dtoTestCreate)
	uc.Error(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseConversionTest))
}
