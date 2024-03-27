package rest

import (
	"fmt"
	"net/http"

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
	r.Get("/", t.get)
}

func (t *Handler) get(w http.ResponseWriter, r *http.Request) {

	num := t.service.GetAverage(r.Context())
	w.Write([]byte(fmt.Sprint(num)))
}
