package rest

import (
	"fmt"
	"net/http"

	"github.com/maltea/go-dependency-injection/pkg/logging"
)

type Handler struct {
	Logger logging.AppLogger
}

func (h *Handler) HandleRequest(rw http.ResponseWriter, req *http.Request) {
	h.Logger.Logf("Request URL: %s", req.URL.String())

	rw.Header().Set("content-type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hi")
}

// CreateHandler creates a new Handler
func CreateHandler(logger logging.AppLogger) *Handler {
	return &Handler{
		Logger: logger,
	}
}
