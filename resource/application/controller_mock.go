package application

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ControllerMock struct {
	mock.Mock
}

func (t *ControllerMock) GetAll(ctx context.Context, cmd *GetAllCmd) []Application {
	return t.Called(cmd).Get(0).([]Application)
}
