package middlewares

import (
	"belajargo/internal/dto"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"golang.org/x/time/rate"
)

func WriteResponse(api huma.API, ctx huma.Context, status int, body any) {
	ctx.SetStatus(status)
	ctx.AppendHeader("Content-Type", "application/json")
	enc := json.NewEncoder(ctx.BodyWriter())
	enc.SetIndent("", "    ")
	enc.Encode(body)
}

func ApiKeyMiddleware(api huma.API, expectedKey string) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		key := ctx.Header("X-API-Key")
		if key != expectedKey {
			WriteResponse(api, ctx, http.StatusUnauthorized, dto.NewErrorResponse("Invalid or missing API Key", dto.StatusUnauthorized, http.StatusUnauthorized))
			return
		}
		next(ctx)
	}
}

func RateLimitMiddleware(api huma.API, r rate.Limit, b int) func(huma.Context, func(huma.Context)) {
	type clientInfo struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var mu sync.Mutex
	clients := make(map[string]*clientInfo)

	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, c := range clients {
				if time.Since(c.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(ctx huma.Context, next func(huma.Context)) {
		ip := ctx.Header("X-Forwarded-For")
		if ip == "" {
			ip = ctx.Header("X-Real-IP")
		}
		if ip == "" {
			ip = ctx.RemoteAddr()
		}

		mu.Lock()
		c, exists := clients[ip]
		if !exists {
			c = &clientInfo{
				limiter: rate.NewLimiter(r, b),
			}
			clients[ip] = c
		}
		c.lastSeen = time.Now()
		limiter := c.limiter
		mu.Unlock()

		if !limiter.Allow() {
			WriteResponse(api, ctx, http.StatusTooManyRequests, dto.NewErrorResponse("Rate limit exceeded", dto.StatusTooManyRequests, http.StatusTooManyRequests))
			return
		}
		next(ctx)
	}
}
