package dto

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
