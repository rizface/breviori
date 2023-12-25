package httpserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rizface/breviori/urlshortener"
)

func RegisterRoutes(r *chi.Mux) {
	deps := deps{
		Shortener: urlshortener.New(),
	}

	r.Post("/short", handlerURLShortener(deps))
	r.Get("/redirecst", func(w http.ResponseWriter, r *http.Request) {
		// http.RedirectHandler("google.com", http.StatusOK).ServeHTTP(w, r)
		w.Header().Add("Location", "http://www.google.com")
		w.WriteHeader(http.StatusTemporaryRedirect)

	})
}
