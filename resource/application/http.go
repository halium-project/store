package application

import (
	"context"
	"encoding/json"
	"net/http"
)

type HTTPHandler struct {
	application ControllerInterface
}

type ControllerInterface interface {
	GetAll(ctx context.Context, cmd *GetAllCmd) []Application
}

func NewHTTPHandler(application ControllerInterface) *HTTPHandler {
	return &HTTPHandler{
		application: application,
	}
}

func (t *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	applications := t.application.GetAll(r.Context(), &GetAllCmd{})

	err := json.NewEncoder(w).Encode(applications)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
