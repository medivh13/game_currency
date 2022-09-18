package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : game-currency
 */

import (
	dtoConversion "game-currency/src/app/dtos/conversion"
	dto "game-currency/src/app/dtos/currency"
	repos "game-currency/src/domain/repositories"
	models "game-currency/src/infra/models"

	"github.com/stretchr/testify/mock"
)

type MockConversionRepo struct {
	mock.Mock
}

func NewMockConversionRepo() *MockConversionRepo {
	return &MockConversionRepo{}
}

var _ repos.ConversionRateRepository = &MockConversionRepo{}

func (o *MockConversionRepo) GetConversionRate(fromID, toID int64) (*models.ConversionRates, error) {
	args := o.Called(fromID, toID)

	var (
		err      error
		respData *models.ConversionRates
	)
	if n, ok := args.Get(0).(*models.ConversionRates); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}

func (o *MockConversionRepo) GetCurrenciesList() (*[]models.Currencies, error) {
	args := o.Called()

	var (
		err      error
		respData *[]models.Currencies
	)
	if n, ok := args.Get(0).(*[]models.Currencies); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}

func (o *MockConversionRepo) CreateCurrency(data *dto.CurrencyReqDTO) error {
	args := o.Called(data)

	var (
		err error
	)
	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (o *MockConversionRepo) CreateConversionRate(data *dtoConversion.CreateConversReqDTO) error {
	args := o.Called(data)

	var (
		err error
	)
	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}
