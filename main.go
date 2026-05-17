package main

import (
	"belajargo/internal/config"
	"belajargo/internal/database"
	"belajargo/internal/handlers"
	"belajargo/internal/services"
	"belajargo/internal/services/redis"
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

	rdb, err := redis.ConncetRedis(cfg)
	if err != nil {
		services.Log(services.Logger{
			Name:    "REDIS",
			Message: fmt.Sprintf("Failed to connect redis: %v", err),
		})
		panic(err)
	}
	defer rdb.Conn().Close()

	redisStore := redis.NewStore(rdb)

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
				Description: `API Belajar Go - RESTful API Performa Tinggi

Sebuah RESTful API modern dan performa tinggi yang dirancang dengan arsitektur Clean Architecture menggunakan Go dan framework Huma v2.

Fitur Utama:
- **Modular & Clean Architecture**: Pemisahan tanggung jawab yang ketat antara Controller/Handler, Service, dan Repository.
- **Endpoint Security**: Dilindungi dengan middleware X-API-Key untuk otentikasi yang aman.
- **Client IP-Based Rate Limiting**: Pembatasan request yang adil per-IP menggunakan algoritma Token Bucket yang thread-safe dan dilengkapi auto-cleanup background memory.
- **High-Performance Caching**: Dukungan Redis Cache-Aside (Lazy Loading) dan Active Cache Invalidation untuk response time sub-milidetik pada data User dan Post yang dinamis.
- **Auto-Generated OpenAPI & Swagger Docs**: Dokumentasi API interaktif yang otomatis dibuat oleh Huma v2 di endpoint /docs.`,
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

	services.Log(services.Logger{
		Name:    "SERVER",
		Message: fmt.Sprintf("Documentation server on %s/docs", cfg.Address()),
	})

	err = http.ListenAndServe(cfg.ListenAddr(), handlers.Handlers(db, api, mux, cfg, redisStore))
	if err != nil {
		panic(err)
	}
}
