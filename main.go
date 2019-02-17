package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/halium-project/go-server-utils/endpoint"
	"github.com/halium-project/go-server-utils/server"
	"github.com/halium-project/store/resource/application"
)

const addr = ":42001"

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Sprintf("Usage: %s [application description file]", os.Args[0]))
	}

	router := mux.NewRouter()

	// Expose the applications resource.
	applicationStorage := application.NewStorage()
	applicationStorage.LoadFromFile(os.Args[1])
	applicationController := application.NewController(applicationStorage)
	applicationHTTPHandler := application.NewHTTPHandler(applicationController)
	router.HandleFunc("/applications", applicationHTTPHandler.GetAll).Methods("GET")

	// Expose utility endpoints.
	router.HandleFunc("/ping", endpoint.Pinger).Methods("GET")

	server.ServeHandler(addr, router)
}
