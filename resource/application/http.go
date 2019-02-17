package application

import (
	"context"
	"net/http"

	"github.com/halium-project/go-server-utils/response"
)

type HTTPHandler struct {
	application ControllerInterface
}

type ControllerInterface interface {
	GetAll(ctx context.Context, cmd *GetAllCmd) map[string]Application
}

func NewHTTPHandler(application ControllerInterface) *HTTPHandler {
	return &HTTPHandler{
		application: application,
	}
}

func (t *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	applications := t.application.GetAll(r.Context(), &GetAllCmd{})

	response.Write(w, http.StatusOK, applications)
}
