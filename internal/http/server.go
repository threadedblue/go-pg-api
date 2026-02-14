package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(h *Handlers) http.Handler {
	r := chi.NewRouter()
	r.Get("/health", Health)
	r.Get("/widgets", h.ListWidgets)
	return r
}
