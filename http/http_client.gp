package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct{}

// Get sends a http request to passed URL and returns received response body
func (client *HttpClient) Get(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("coudn't initialize http request, err: %s", err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("coudn't make an http request, err: %s", err)
	}

	if response.StatusCode >= 400 {
		return "", fmt.Errorf("invalid response from API, status: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("coudn't read API response, err: %s", err)
	}

	return string(body), err
}
