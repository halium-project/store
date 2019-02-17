package application

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ControllerMock struct {
	mock.Mock
}

func (t *ControllerMock) GetAll(ctx context.Context, cmd *GetAllCmd) map[string]Application {
	return t.Called(cmd).Get(0).(map[string]Application)
}
