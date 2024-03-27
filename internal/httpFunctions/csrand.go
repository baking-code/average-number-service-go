package httpfunctions

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type csrandResponse struct {
	Status string `json:"status"`
	Random int    `json:"random"`
}

func Fetch() int {
	url := "https://csrng.net/csrng/csrng.php?min=0&max=100"
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("No response from request", err)
		panic(err)
	}
	defer resp.Body.Close()
	var result []csrandResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Cannot read from Body", err)
		panic(err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		slog.Error("Can not unmarshal JSON", err)
		panic(err)
	}
	if result[0].Status != "success" {
		slog.Error("Bad response", err)
		panic(fmt.Errorf("bad response"))
	}
	return result[0].Random
}
