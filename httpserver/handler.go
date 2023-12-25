package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/rizface/breviori/urlshortener"
)

type Shortener interface {
	Short(context.Context, string) (string, error)
}

func handlerURLShortener(ctx context.Context, deps deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Payload struct {
			URL string `json:"url"`
		}

		if err := json.NewDecoder(r.Body).Decode(&Payload); err != nil {
			writeHTTPResponse(w, httpResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid request body",
				Data:       nil,
			})
			return
		}

		if err := validation.Validate(Payload.URL, validation.Required, is.URL); err != nil {
			writeHTTPResponse(w, httpResponse{
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("Invalid URL: %v", err),
				Data:       nil,
			})
			return
		}

		shortURL, err := deps.Shortener.Short(ctx, Payload.URL)
		if errors.Is(err, urlshortener.ErrorKeyGen) {
			writeHTTPResponse(w, httpResponse{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    "Failed to generate short URL",
				Data:       nil,
			})
		}

		if err != nil {
			slog.Error(fmt.Sprintf("failed to short URL: %v", err))
			writeHTTPResponse(w, httpResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to short URL",
				Data:       nil,
			})
			return
		}

		writeHTTPResponse(w, httpResponse{
			StatusCode: http.StatusOK,
			Message:    "Success",
			Data: map[string]interface{}{
				"shortUrl": shortURL,
			},
		})
	}
}
