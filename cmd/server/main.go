package main

import (
	"net/http"

	"github.com/maltea/go-dependency-injection/pkg/http/rest"
	"github.com/maltea/go-dependency-injection/pkg/logging"
)

func main() {
	logger := logging.CreateLogger("Log: ")
	logger.Logf("Server listeing on :3000")

	handler := rest.CreateHandler(logger)

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(handler.HandleRequest),
	}

	server.ListenAndServe()
}
