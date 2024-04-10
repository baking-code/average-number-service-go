package main

import (
	"github.com/baking-code/average-number-service-go/internal/app"
	httpfunctions "github.com/baking-code/average-number-service-go/internal/httpFunctions"
)

func main() {
	app.Server(httpfunctions.Fetch)
}
