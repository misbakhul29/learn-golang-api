package common

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"golang.org/x/time/rate"
)

func ApiKeyMiddleware(api huma.API, expectedKey string) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		key := ctx.Header("X-API-Key")
		if key != expectedKey {
			WriteResponse(api, ctx, http.StatusUnauthorized, NewErrorResponse("Invalid or missing API Key", StatusUnauthorized, http.StatusUnauthorized))
			return
		}
		next(ctx)
	}
}

func RateLimitMiddleware(api huma.API, r rate.Limit, b int) func(huma.Context, func(huma.Context)) {
	limiter := rate.NewLimiter(r, b)
	return func(ctx huma.Context, next func(huma.Context)) {
		if !limiter.Allow() {
			WriteResponse(api, ctx, http.StatusTooManyRequests, NewErrorResponse("Rate limit exceeded", StatusTooManyRequests, http.StatusTooManyRequests))
			return
		}
		next(ctx)
	}
}
