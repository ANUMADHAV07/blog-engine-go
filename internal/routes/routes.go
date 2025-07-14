package routes

import (
	"github.com/ANUMADHAV07/blog-engine-go.git/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoute(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	// r.Get("/handler",)

	return r
}
