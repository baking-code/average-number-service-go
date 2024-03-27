package service

import "context"

type Service interface {
	GetAverage(ctx context.Context) float64
	GetStoredCount(ctx context.Context, n int) int
	Run() chan bool
}
