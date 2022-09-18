package mock_dto

import (
	dto "game-currency/src/app/dtos/conversion"

	"github.com/stretchr/testify/mock"
)

type MockConversionDTO struct {
	mock.Mock
}

func NewMockConversionDTO() *MockConversionDTO {
	return &MockConversionDTO{}
}

var _ dto.ConversInterface = &MockConversionDTO{}

func (m *MockConversionDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
