package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/baking-code/average-number-service-go/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

// Register connects the handlers to the router.
func (t *Handler) Register(r chi.Router) {
	r.Get("/", t.getCurrentRandom)
	r.Get("/{count}", t.getCount)
}

func (t *Handler) getCurrentRandom(w http.ResponseWriter, r *http.Request) {
	num := t.service.GetAverage(r.Context())
	w.Write([]byte(fmt.Sprint(num)))
}

func (t *Handler) getCount(w http.ResponseWriter, r *http.Request) {
	countRaw := chi.URLParam(r, "count")
	count, err := strconv.Atoi(countRaw)
	if err != nil {
		message := fmt.Sprintf("cannot parse %s as a number", countRaw)
		slog.Error(message, err)
		http.Error(w, message, 400)
	} else {
		num := t.service.GetStoredCount(r.Context(), count)
		w.Write([]byte(fmt.Sprint(num)))
	}
}
