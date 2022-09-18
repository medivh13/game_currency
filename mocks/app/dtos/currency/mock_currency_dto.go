package mock_dto

import (
	dto "game-currency/src/app/dtos/currency"

	"github.com/stretchr/testify/mock"
)

type MockCurrencyDTO struct {
	mock.Mock
}

func NewMockCurrencyDTO() *MockCurrencyDTO {
	return &MockCurrencyDTO{}
}

var _ dto.CurrencyDTOInterface = &MockCurrencyDTO{}

func (m *MockCurrencyDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
