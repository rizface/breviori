package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rizface/breviori/urlshortener"
)

func RegisterRoutes(r *chi.Mux) {
	var (
		ctx  = context.Background()
		deps = deps{
			Shortener: urlshortener.New(),
		}
	)

	r.Post("/short", handlerURLShortener(ctx, deps))
	r.Get("/redirect/{key}", func(w http.ResponseWriter, r *http.Request) {
		// http.RedirectHandler("google.com", http.StatusOK).ServeHTTP(w, r)
		fmt.Println(chi.URLParam(r, "key"))
	})
}
