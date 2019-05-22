package handlers

import (
	"log"
	"net/http"
	"os"

	"craft-advanced-ultimate-go/sv/internal/platform/web"
	"craft-advanced-ultimate-go/sv/internal/mid"
)

// API returns a handler for a set of routes.
func API(shutdown chan os.Signal, log *log.Logger) http.Handler {
	app := web.NewApp(shutdown, log, mid.Logger(log), mid.Errors(log), mid.Metrics())

	check := Check{Log: log}
	app.Handle("GET", "/v1/health", check.Health)

	return app
}
