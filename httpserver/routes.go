package httpserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rizface/breviori/database"
	"github.com/rizface/breviori/urlshortener"
)

func RegisterRoutes(r *chi.Mux) {
	shortnedURLRepo := database.NewShortnedURL()
	cacheShortnedURL := database.NewRedisClient()

	var (
		deps = deps{
			Shortener: urlshortener.New(shortnedURLRepo, cacheShortnedURL),
		}
	)

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello")
	})

	r.Post("/short", handlerURLShortener(deps))
	r.Get("/redirect/{key}", handlerRedirection(deps))
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
}
