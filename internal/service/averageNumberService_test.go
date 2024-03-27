package service

import (
	"context"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	iter := 0
	service := NewAverageNumberService(func() int {
		iter++
		return rand.IntN(100)
	})
	stop := service.Run()
	time.Sleep(2100 * time.Millisecond)
	assert.Equal(t, 3, iter)
	assert.Equal(t, true, service.GetAverage(context.Background()) < 100.0)
	stop <- true
	// verify we can terminate
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 3, iter)

}
