package mock_cases

import (
	dto "game-currency/src/app/dtos/currency"
	usecase "game-currency/src/app/use_cases/currency"

	"github.com/stretchr/testify/mock"
)

type MockCurrencyUseCase struct {
	mock.Mock
}

func NewMockCurrencyUseCase() *MockCurrencyUseCase {
	return &MockCurrencyUseCase{}
}

var _ usecase.CurrencyUsecaseInterface = &MockCurrencyUseCase{}

func (m *MockCurrencyUseCase) GetCurrenciesList() ([]*dto.CurrencyRespDTO, error) {
	args := m.Called()
	var err error
	var resp []*dto.CurrencyRespDTO

	if n, ok := args.Get(0).([]*dto.CurrencyRespDTO); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}

func (m *MockCurrencyUseCase) CreateCurrency(data *dto.CurrencyReqDTO) error {
	args := m.Called(data)
	var err error

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}
