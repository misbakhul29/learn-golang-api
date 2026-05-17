package api

import (
	"belajargo/internal/dto"
	"belajargo/internal/models"
	"belajargo/internal/services"
	"belajargo/internal/services/api_services"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *api_services.UserServiceImpl
}

func NewUserHandler(service *api_services.UserServiceImpl) *UserHandler {
	return &UserHandler{service: service}
}

func RegisterUserRoutes(r *dto.Router, handler *UserHandler) {
	op := services.NewOp(r, "get-users", http.MethodGet, "/api/users", "List all users", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct{}) (*dto.BaseResponse[[]models.Users], error) {
			users, err := handler.service.GetAll()
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal mengambil data user dari database", err)
			}
			return dto.NewResponse(users, "Berhasil mengambil data user", dto.StatusSuccess, http.StatusOK), nil
		},
	)

	op = services.NewOp(r, "get-user", http.MethodGet, "/api/users/{id}", "Get user by id", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID uuid.UUID `path:"id" doc:"User ID"`
		}) (*dto.BaseResponse[*models.Users], error) {
			user, err := handler.service.GetByID(input.ID)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal mengambil data user dari database", err)
			}
			return dto.NewResponse(user, "Berhasil mengambil data user", dto.StatusSuccess, http.StatusOK), nil
		},
	)

	op = services.NewOp(r, "create-user", http.MethodPost, "/api/users", "Create a new user", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			Body dto.InputUser
		}) (*dto.BaseResponse[*models.Users], error) {
			user, err := handler.service.Create(&input.Body)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal membuat user", err)
			}
			return dto.NewResponse(user, "Berhasil membuat user", dto.StatusCreated, http.StatusCreated), nil
		},
	)

	op = services.NewOp(r, "update-user", http.MethodPut, "/api/users/{id}", "Update user by id", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID   uuid.UUID     `path:"id" doc:"User ID"`
			Body dto.InputUser `json:"body" doc:"User data"`
		}) (*dto.BaseResponse[*models.Users], error) {
			user, err := handler.service.Update(input.ID, &input.Body)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal memperbarui user", err)
			}
			return dto.NewResponse(user, "Berhasil memperbarui user", dto.StatusSuccess, http.StatusOK), nil
		},
	)

	op = services.NewOp(r, "delete-user", http.MethodDelete, "/api/users/{id}", "Delete user by id", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID uuid.UUID `path:"id" doc:"User ID"`
		}) (*dto.BaseResponse[*models.Users], error) {
			err := handler.service.Delete(input.ID)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal menghapus user", err)
			}
			return dto.NewResponse((*models.Users)(nil), "Berhasil menghapus user", dto.StatusSuccess, http.StatusOK), nil
		},
	)
}
