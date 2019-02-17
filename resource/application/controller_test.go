package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validAppDescription = Application{
	Name:        "Some App",
	Description: "A simple app",
	OauthInfos: OauthInfos{
		RedirectURIs:  []string{"some-url"},
		GrantTypes:    []string{"implicit", "refresh_token"},
		ResponseTypes: []string{"token", "code"},
		Scopes:        []string{"permissions"},
		Public:        true,
	},
}

func Test_Controller_GetAll_success(t *testing.T) {
	storageMock := new(StorageMock)
	controller := NewController(storageMock)

	storageMock.On("GetAll").Return([]Application{validAppDescription}).Once()

	res := controller.GetAll(context.Background(), &GetAllCmd{})

	assert.EqualValues(t, []Application{validAppDescription}, res)

	storageMock.AssertExpectations(t)
}
