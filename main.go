package main

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/halium-project/go-server-utils/endpoint"
	"github.com/halium-project/go-server-utils/server"
)

const addr = ":42001"

func main() {
	ctx := context.Background()

	router := mux.NewRouter()

	// Expose utility endpoints.
	router.HandleFunc("/ping", endpoint.Pinger).Methods("GET")

	server.ServeHandler(router)
}
