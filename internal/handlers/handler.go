package handlers

import (
	"belajargo/internal/config"
	"belajargo/internal/dto"
	"belajargo/internal/handlers/api"
	"belajargo/internal/repositories"
	"belajargo/internal/services/api_services"
	"belajargo/internal/services/redis"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func Handlers(db *gorm.DB, HApi huma.API, mux *http.ServeMux, cfg *config.Config, rdb *redis.Store) http.Handler {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	userRepo := repositories.NewUserRepository(db)
	userService := api_services.NewUserService(userRepo, rdb)
	userHandler := api.NewUserHandler(userService)
	api.RegisterUserRoutes(&dto.Router{
		DB:  db,
		API: HApi,
		CFG: cfg,
	}, userHandler)

	postRepo := repositories.NewPostRepository(db)
	postService := api_services.NewPostService(postRepo, rdb)
	postHandler := api.NewPostHandler(postService)
	api.RegisterPostRoutes(&dto.Router{
		DB:  db,
		API: HApi,
		CFG: cfg,
	}, postHandler)

	return mux
}
