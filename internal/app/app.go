package app

import (
	"net/http"

	httpfunctions "github.com/baking-code/average-number-service-go/internal/httpFunctions"
	"github.com/baking-code/average-number-service-go/internal/rest"
	"github.com/baking-code/average-number-service-go/internal/service"
	"github.com/go-chi/chi/v5"
)

func Server() {
	r := chi.NewRouter()
	service := service.NewAverageNumberService(httpfunctions.Fetch)
	handler := rest.NewHandler(service)
	handler.Register(r)
	service.Run()
	http.ListenAndServe(":3333", r)
}
