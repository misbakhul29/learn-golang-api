package routes

import (
	"belajargo/internal/config"
	"belajargo/internal/routes/common"
	"belajargo/internal/routes/users"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func Handlers(db *gorm.DB, api huma.API, mux *http.ServeMux, cfg *config.Config) http.Handler {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	users.UsersRoute(common.Router{
		DB:  db,
		API: api,
		CFG: cfg,
	})

	return mux
}
