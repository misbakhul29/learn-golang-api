package services

import (
	"belajargo/internal/dto"
	"belajargo/internal/middlewares"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
	"golang.org/x/time/rate"
)

func NewOp(r *dto.Router, id, method, path, summary string, tags []string) huma.Operation {
	registry := r.API.OpenAPI().Components.Schemas
	schema := registry.Schema(reflect.TypeOf(dto.BaseResponse[any]{}), true, "BaseResponse")

	return huma.Operation{
		OperationID: id,
		Method:      method,
		Path:        path,
		Summary:     summary,
		Tags:        tags,
		Middlewares: huma.Middlewares{
			middlewares.RateLimitMiddleware(r.API, rate.Limit(r.CFG.ApiRateLimit), r.CFG.ApiRateLimitBurst),
			middlewares.ApiKeyMiddleware(r.API, r.CFG.ApiKey),
		},
		Security: []map[string][]string{
			{"api_key": {}},
		},
		DefaultStatus: 200,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Berhasil mendapatkan response",
				Content: map[string]*huma.MediaType{
					"application/json": {
						Schema: schema,
					},
				},
			},
			"400": {
				Description: "Bad Request - Permintaan tidak valid",
				Content: map[string]*huma.MediaType{
					"application/json": {
						Schema: schema,
					},
				},
			},
			"500": {
				Description: "Internal Server Error - Terjadi kesalahan pada server",
				Content: map[string]*huma.MediaType{
					"application/json": {
						Schema: schema,
					},
				},
			},
		},
	}
}
