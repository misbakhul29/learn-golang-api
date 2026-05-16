package main

import (
	"belajargo/internal/config"
	"belajargo/internal/database"
	"belajargo/internal/routes"
	"belajargo/internal/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	cfg := config.LoadEnv()

	db, err := database.ConnectDB()
	if err != nil {
		services.Log(services.Logger{
			Name:    "DATABASE",
			Message: fmt.Sprintf("Failed to connect database: %v", err),
		})
		panic(err)
	}

	mux := http.NewServeMux()
	formats := make(map[string]huma.Format)
	for k, v := range huma.DefaultFormats {
		formats[k] = v
	}

	formats["application/json"] = huma.Format{
		Marshal: func(w io.Writer, v any) error {
			enc := json.NewEncoder(w)
			enc.SetIndent("", "    ")
			return enc.Encode(v)
		},
		Unmarshal: json.Unmarshal,
	}

	api := humago.New(mux, huma.Config{
		OpenAPI: &huma.OpenAPI{
			OpenAPI: "3.1.0",
			Info: &huma.Info{
				Title:       "API Belajar Go",
				Version:     "1.0.0",
				Description: "Belajar membuat API dengan Go",
				Contact: &huma.Contact{
					Name:  "Misbakhul Munir",
					URL:   "https://misbakhul.my.id",
					Email: "misbakhul2904@gmail.com",
				},
			},
			Components: &huma.Components{
				SecuritySchemes: map[string]*huma.SecurityScheme{
					"api_key": {
						Type: "apiKey",
						In:   "header",
						Name: "X-API-Key",
					},
				},
			},
			Servers: []*huma.Server{
				{URL: cfg.Address()},
			},
			Tags: []*huma.Tag{
				{
					Name:        "Users",
					Description: "Endpoints untuk mengelola user",
				},
			},
			Security: []map[string][]string{
				{
					"api_key": {},
				},
				{
					"rate_limit": {},
				},
			},
		},
		DocsPath:    "/docs",
		OpenAPIPath: "/docs",
		SchemasPath: "/schemas",
		Formats:     formats,
	})

	services.Log(services.Logger{
		Name:    "SERVER",
		Message: fmt.Sprintf("Starting server on %s", cfg.Address()),
	})

	err = http.ListenAndServe(cfg.ListenAddr(), routes.Handlers(db, api, mux, cfg))
	if err != nil {
		panic(err)
	}
}
