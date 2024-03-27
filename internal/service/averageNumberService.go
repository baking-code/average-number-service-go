package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/baking-code/average-number-service-go/internal/utils/average"
)

type AverageNumberService struct {
	counter       average.AverageNumberCounter
	fetchFunction func() int
}

func (s *AverageNumberService) getData() {
	slog.Info("running fetch function")
	randomNumber := s.fetchFunction()

	s.counter.AddNewNumber(randomNumber)
}

func (s *AverageNumberService) GetStoredCount(ctx context.Context, n int) int {
	return s.counter.Store[n]
}

func (s *AverageNumberService) GetAverage(ctx context.Context) float64 {
	return s.counter.Average
}

func (s *AverageNumberService) Run() chan bool {
	done := make(chan bool, 1)
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		s.getData()
		for {
			select {
			case <-ticker.C:
				s.getData()

			case <-done:

			}
		}
	}()
	return done
}

func NewAverageNumberService(f func() int) Service {
	return &AverageNumberService{
		counter:       average.NewAverageNumberCounter(),
		fetchFunction: f,
	}
}
