package middleware

import (
	"net/http"

	// "github.com/go-chi/chi"
	// "github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func CORSMiddleware(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	return c.Handler(h)
}

func Recoverer() {
	// chi.NewRouter()

	// middleware.Recoverer()
}
