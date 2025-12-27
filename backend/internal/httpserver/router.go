// Package httpserver
package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func applyGenericMiddleware(mx *chi.Mux) {
	mx.Use(middleware.Logger)
}

func CreateRoutes() http.Handler {
	r := chi.NewRouter()
	applyGenericMiddleware(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	return r
}
