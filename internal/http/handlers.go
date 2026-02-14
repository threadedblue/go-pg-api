package httpapi

import (
	"encoding/json"
	"net/http"

	"semita.wk/go-pg-api/internal/repo"
)

type Handlers struct {
	widgets *repo.WidgetRepo
}

func NewHandlers(w *repo.WidgetRepo) *Handlers {
	return &Handlers{widgets: w}
}

func (h *Handlers) ListWidgets(w http.ResponseWriter, r *http.Request) {
	items, err := h.widgets.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
