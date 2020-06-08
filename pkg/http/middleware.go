package http

import (
	"net/http"

	"github.com/rs/cors"
)

/**
* TODO: API Token validation
 */

func JSONApi(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func CORS(h http.Handler) http.Handler {
	corsConfig := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowedMethods: []string{"POST", "PUT", "GET", "PATCH", "OPTIONS", "HEAD", "DELETE"},
		Debug:          true,
	})
	return corsConfig.Handler(h)
}

func UseMiddleware(h http.Handler) http.Handler {
	h = JSONApi(h)
	h = CORS(h)
	return h
}
