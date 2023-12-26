package httpserver

import (
	"github.com/go-chi/chi/v5"
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

	r.Post("/short", handlerURLShortener(deps))
	r.Get("/redirect/{key}", handlerRedirection(deps))
}
