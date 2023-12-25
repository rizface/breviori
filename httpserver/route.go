package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/rizface/breviori/urlshortener"
)

func RegisterRoutes(r *chi.Mux) {
	deps := deps{
		Shortener: urlshortener.New(),
	}

	r.Post("/short", handlerURLShortener(deps))
}
