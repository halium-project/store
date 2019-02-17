package application

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type StorageMock struct {
	mock.Mock
}

func (t *StorageMock) GetAll(ctx context.Context) []Application {
	return t.Called().Get(0).([]Application)
}
