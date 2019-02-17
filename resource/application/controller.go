package application

import (
	"context"
	"strings"
	"time"
)

type Controller struct {
	storage StorageInterface
}

type StorageInterface interface {
	GetAll(ctx context.Context) []Application
}

func NewController(storage StorageInterface) *Controller {
	return &Controller{
		storage: storage,
	}
}

func (t *Controller) GetAll(ctx context.Context, cmd *GetAllCmd) map[string]Application {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	apps := t.storage.GetAll(ctx)

	res := make(map[string]Application, len(apps))
	for _, app := range apps {
		appID := strings.Replace(strings.ToLower(app.Name), " ", "-", -1)
		res[appID] = app
	}

	return res
}
