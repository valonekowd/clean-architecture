package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Router interface {
	Get(pattern string, h http.HandlerFunc)
	Post(pattern string, h http.HandlerFunc)
	Put(pattern string, h http.HandlerFunc)
	Delete(pattern string, h http.HandlerFunc)
	Route(pattern string, fn func(r Router)) Router
	Use(middlewares ...func(http.Handler) http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

var URLParam = chi.URLParam

func QueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
