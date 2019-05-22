package mid

import (
	"craft-advanced-ultimate-go/sv/internal/platform/web"
	"log"
	"net/http"
	"context"
	"time"
	"go.opencensus.io/trace"
)

func Logger(log *log.Logger) web.Middleware {
	f := func(before web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			err := before(ctx, w, r, params)

			ctx, span := trace.StartSpan(ctx, "internal.mid.Logger")
			defer span.End()

			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return web.Shutdown("web value missing from context")
			}

			log.Printf("%s : (%d) : %s %s -> %s (%s)\n",
				v.TraceID,
				v.StatusCode,
				r.Method, r.URL.Path,
				r.RemoteAddr, time.Since(v.Now),
			)

			// For consistency return the error we received.
			return err
		}

		return h
	}

	return f
}
