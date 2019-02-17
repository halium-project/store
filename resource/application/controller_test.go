package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_GetAll_success(t *testing.T) {
	storageMock := new(StorageMock)
	controller := NewController(storageMock)

	storageMock.On("GetAll").Return([]Application{ValidApp}).Once()

	res := controller.GetAll(context.Background(), &GetAllCmd{})

	assert.EqualValues(t, map[string]Application{
		ValidAppID: ValidApp,
	}, res)

	storageMock.AssertExpectations(t)
}
