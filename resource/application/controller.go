package application

import (
	"context"
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

func (t *Controller) GetAll(ctx context.Context, cmd *GetAllCmd) []Application {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	return t.storage.GetAll(ctx)
}
