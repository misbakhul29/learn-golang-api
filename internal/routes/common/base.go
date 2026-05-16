package common

import (
	"belajargo/internal/config"
	"encoding/json"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type StatusResponse string

const (
	StatusSuccess             StatusResponse = "success"
	StatusError               StatusResponse = "error"
	StatusRedirect            StatusResponse = "redirect"
	StatusUnauthorized        StatusResponse = "unauthorized"
	StatusForbidden           StatusResponse = "forbidden"
	StatusNotFound            StatusResponse = "not_found"
	StatusTooManyRequests     StatusResponse = "too_many_requests"
	StatusInternalServerError StatusResponse = "internal_server_error"
	StatusBadRequest          StatusResponse = "bad_request"
	StatusCreated             StatusResponse = "created"
	StatusDeleted             StatusResponse = "deleted"
	StatusUpdated             StatusResponse = "updated"
	StatusNoContent           StatusResponse = "no_content"
)

type BaseResponse[T any] struct {
	Body struct {
		Data       T              `json:"data" doc:"Data utama yang dikembalikan oleh API"`
		StatusText StatusResponse `json:"status" doc:"Status response (success, error, etc)"`
		Code       int            `json:"code" doc:"Application specific error/success code"`
		Message    string         `json:"message" doc:"Pesan penjelasan mengenai status response"`
	}
}

func NewResponse[T any](data T, message string, status StatusResponse, code int) *BaseResponse[T] {
	resp := &BaseResponse[T]{}
	resp.Body.Data = data
	resp.Body.StatusText = status
	resp.Body.Code = code
	resp.Body.Message = message
	return resp
}

func NewErrorResponse(message string, status StatusResponse, code int) *BaseResponse[any] {
	resp := &BaseResponse[any]{}
	resp.Body.Data = nil
	resp.Body.StatusText = status
	resp.Body.Code = code
	resp.Body.Message = message
	return resp
}

func WriteResponse(api huma.API, ctx huma.Context, status int, body any) {
	ctx.SetStatus(status)
	ctx.AppendHeader("Content-Type", "application/json")
	enc := json.NewEncoder(ctx.BodyWriter())
	enc.SetIndent("", "    ")
	enc.Encode(body)
}

type Router struct {
	DB  *gorm.DB
	API huma.API
	CFG *config.Config
}

func (r *Router) NewOp(id, method, path, summary string, tags []string) huma.Operation {
	// Generate schema otomatis dari struct BaseResponse
	registry := r.API.OpenAPI().Components.Schemas
	schema := registry.Schema(reflect.TypeOf(BaseResponse[any]{}), true, "BaseResponse")

	return huma.Operation{
		OperationID: id,
		Method:      method,
		Path:        path,
		Summary:     summary,
		Tags:        tags,
		Middlewares: huma.Middlewares{
			RateLimitMiddleware(r.API, rate.Limit(r.CFG.ApiRateLimit), r.CFG.ApiRateLimitBurst),
			ApiKeyMiddleware(r.API, r.CFG.ApiKey),
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
