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

type PostHandler struct {
	service *api_services.PostServiceImpl
}

func NewPostHandler(service *api_services.PostServiceImpl) *PostHandler {
	return &PostHandler{service: service}
}

func RegisterPostRoutes(r *dto.Router, handler *PostHandler) {
	op := services.NewOp(r, "get-posts", http.MethodGet, "/api/posts", "List all posts", []string{"Posts"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			IncludeUsers bool `query:"include_users" doc:"Include users"`
		}) (*dto.BaseResponse[[]models.Posts], error) {
			posts, err := handler.service.GetAll(input.IncludeUsers)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal mengambil data post dari database", err)
			}
			return dto.NewResponse(posts, "Berhasil mengambil data post", dto.StatusSuccess, http.StatusOK), nil
		},
	)

	op = services.NewOp(r, "get-post", http.MethodGet, "/api/posts/{id}", "Get post by id", []string{"Posts"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID uuid.UUID `path:"id" doc:"Post ID"`
		}) (*dto.BaseResponse[*models.Posts], error) {
			post, err := handler.service.GetByID(input.ID)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal mengambil data post dari database", err)
			}
			return dto.NewResponse(post, "Berhasil mengambil data post", dto.StatusSuccess, http.StatusOK), nil
		},
	)

	op = services.NewOp(r, "create-post", http.MethodPost, "/api/posts", "Create a new post", []string{"Posts"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			Body dto.InputPosts
		}) (*dto.BaseResponse[*models.Posts], error) {
			post, err := handler.service.Create(&input.Body)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal membuat post", err)
			}
			return dto.NewResponse(post, "Berhasil membuat post", dto.StatusCreated, http.StatusCreated), nil
		},
	)

	op = services.NewOp(r, "update-post", http.MethodPut, "/api/posts/{id}", "Update post by id", []string{"Posts"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID   uuid.UUID      `path:"id" doc:"Post ID"`
			Body dto.InputPosts `json:"body" doc:"Post data"`
		}) (*dto.BaseResponse[*models.Posts], error) {
			post, err := handler.service.Update(input.ID, &input.Body)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal memperbarui post", err)
			}
			return dto.NewResponse(post, "Berhasil memperbarui post", dto.StatusSuccess, http.StatusOK), nil
		},
	)

	op = services.NewOp(r, "delete-post", http.MethodDelete, "/api/posts/{id}", "Delete post by id", []string{"Posts"})
	huma.Register(r.API, op,
		func(ctx context.Context, input *struct {
			ID uuid.UUID `path:"id" doc:"Post ID"`
		}) (*dto.BaseResponse[*models.Posts], error) {
			err := handler.service.Delete(input.ID)
			if err != nil {
				return nil, huma.NewError(http.StatusInternalServerError, "Gagal menghapus post", err)
			}
			return dto.NewResponse((*models.Posts)(nil), "Berhasil menghapus post", dto.StatusSuccess, http.StatusOK), nil
		},
	)
}
