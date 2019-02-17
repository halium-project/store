package application

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Storage struct {
	applications []Application
}

func NewStorage() *Storage {
	return &Storage{}
}

func (t *Storage) GetAll(ctx context.Context) []Application {
	return t.applications
}

func (t *Storage) LoadFromFile(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("failed to read the file content: %s", err))
	}

	var applications []Application
	err = json.Unmarshal(file, &applications)
	if err != nil {
		panic(fmt.Sprintf("failed to decode the application description file: %s", err))
	}

	t.applications = applications
}
