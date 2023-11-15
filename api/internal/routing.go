package internal

import (
	"github.com/go-chi/chi/v5"
)

func InitRouter(r chi.Router) {
	h := NewHandler()
	r.Get("/health", h.Health())
	r.Post("/", h.Handle())
}
