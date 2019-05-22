package handlers

import (
	"log"
	"net/http"
	"context"
	"craft-advanced-ultimate-go/sv/internal/platform/web"
	"math/rand"
)

// Check provides support for orchestration health checks.
type Check struct {
	Log *log.Logger
}

// Health validates the service is healthy and ready to accept requests.
func (c *Check) Health(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	if n := rand.Intn(100); n%2 == 0 {
		// TODO Cause miatt nem megy
		//return web.Shutdown("nooooooo")
		// TODO Cause miatt nem megy
		return web.RespondErrorMessage("testinge", http.StatusNotFound)
	}

	status := struct {
		Status string
	}{
		Status: "Ok",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
