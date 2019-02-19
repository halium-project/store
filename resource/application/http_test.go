package application

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_HTTP_GetAll_success(t *testing.T) {
	controllerMock := new(ControllerMock)
	httpHandler := NewHTTPHandler(controllerMock)

	controllerMock.On("GetAll", &GetAllCmd{}).Return(map[string]Application{
		ValidApp.ID: ValidApp,
	}).Once()

	r := httptest.NewRequest("GET", "http://example.com/applications", nil)
	w := httptest.NewRecorder()

	httpHandler.GetAll(w, r)

	res := w.Result()
	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.JSONEq(t, `{
		"some-app": {
			"id": "some-app",
			"name": "Some App",
			"description": "A simple app",
			"redirectURIs": ["some-url"],
			"grantTypes": ["implicit", "refresh_token"],
			"responseTypes": ["token", "code"],
			"scopes": ["permissions"],
			"public": true
		}
	}`, string(body))

	controllerMock.AssertExpectations(t)
}
