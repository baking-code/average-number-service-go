package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/baking-code/average-number-service-go/internal/rest"
	"github.com/baking-code/average-number-service-go/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func Server(f func() int) {
	logger := httplog.NewLogger("average-number-service", httplog.Options{
		LogLevel:         slog.LevelDebug,
		JSON:             true,
		Concise:          true,
		RequestHeaders:   true,
		ResponseHeaders:  true,
		MessageFieldName: "message",
		LevelFieldName:   "level",
		TimeFieldFormat:  time.RFC1123,
		QuietDownPeriod:  10 * time.Second,
	})
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger, []string{"/ping"}))
	r.Use(middleware.Heartbeat("/ping"))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			httplog.LogEntrySetField(ctx, "user", slog.StringValue("user1"))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	service := service.NewAverageNumberService(f)
	handler := rest.NewHandler(service)
	handler.Register(r)
	service.Run()
	port := ":3333"
	slog.Info(fmt.Sprintf("Server now running on port %s", port))
	go func() {
		err := http.ListenAndServe(port, r)
		if err != nil {
			slog.Error("error running server", err)
			panic(err)
		}
	}()
}
