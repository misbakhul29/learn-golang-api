package users

import (
	"belajargo/internal/models"
	"belajargo/internal/routes/common"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

func UsersRoute(r common.Router) {
	op := r.NewOp("get-users", http.MethodGet, "/api/users", "List all users", []string{"Users"})
	huma.Register(r.API, huma.Operation(op),
		func(ctx context.Context, input *struct{}) (*common.BaseResponse[[]models.Users], error) {
			users, err := models.GetAllUsers(r.DB)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal mengambil data user dari database", err)
			}
			return common.NewResponse(users, "Berhasil mengambil data user", common.StatusSuccess, http.StatusOK), nil
		},
	)

	op = r.NewOp("get-user", http.MethodGet, "/api/users/{id}", "Get user by id", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID uuid.UUID `path:"id" doc:"User ID"`
		}) (*common.BaseResponse[*models.Users], error) {
			users, err := models.GetUserById(r.DB, input.ID)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal mengambil data user dari database", err)
			}
			return common.NewResponse(users, "Berhasil mengambil data user", common.StatusSuccess, http.StatusOK), nil
		},
	)

	op = r.NewOp("create-user", http.MethodPost, "/api/users", "Create a new user", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			Body models.InputUser
		}) (*common.BaseResponse[*models.Users], error) {
			users, err := models.CreateUser(r.DB, &input.Body)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal membuat user", err)
			}
			return common.NewResponse(users, "Berhasil membuat user", common.StatusCreated, http.StatusCreated), nil
		},
	)

	op = r.NewOp("update-user", http.MethodPut, "/api/users/{id}", "Update user by id", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID   uuid.UUID        `path:"id" doc:"User ID"`
			Body models.InputUser `json:"body" doc:"User data"`
		}) (*common.BaseResponse[*models.Users], error) {
			users, err := models.UpdateUserById(r.DB, &input.Body, input.ID)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal memperbarui user", err)
			}
			return common.NewResponse(users, "Berhasil memperbarui user", common.StatusSuccess, http.StatusOK), nil
		},
	)

	op = r.NewOp("delete-user", http.MethodDelete, "/api/users/{id}", "Delete user by id", []string{"Users"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID uuid.UUID `path:"id" doc:"User ID"`
		}) (*common.BaseResponse[*models.Users], error) {
			users, err := models.DeleteUserById(r.DB, input.ID)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal menghapus user", err)
			}
			return common.NewResponse(users, "Berhasil menghapus user", common.StatusSuccess, http.StatusOK), nil
		},
	)
}
