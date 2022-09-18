package mock_cases

import (
	dto "game-currency/src/app/dtos/conversion"
	usecase "game-currency/src/app/use_cases/conversion"

	"github.com/stretchr/testify/mock"
)

type MockConversionUseCase struct {
	mock.Mock
}

func NewMockConversionUseCase() *MockConversionUseCase {
	return &MockConversionUseCase{}
}

var _ usecase.ConversionUsecaseInterface = &MockConversionUseCase{}

func (m *MockConversionUseCase) ConversResult(data *dto.ConversReqDTO) (*dto.ConversRespDTO, error) {
	args := m.Called(data)
	var (
		err      error
		respData *dto.ConversRespDTO
	)
	if n, ok := args.Get(0).(*dto.ConversRespDTO); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}

func (m *MockConversionUseCase) CreateConversionRate(data *dto.CreateConversReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}
