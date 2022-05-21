package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Pawilonek/factorio-mod-portal-golang-client/response"
)

const urlApi = "https://mods.factorio.com"
const urlAssets = "https://assets-mod.factorio.com"

type Config struct {
	Timeout   int
	UrlApi    string
	UrlAssets string
}

type Client struct {
	config     Config
	httpClient httpClient
}

func New(config Config) *Client {
	if config.Timeout == 0 {
		config.Timeout = int(time.Second)
	}

	if config.UrlApi == "" {
		config.UrlApi = urlApi
	}

	if config.UrlAssets == "" {
		config.UrlAssets = urlAssets
	}

	return &Client{
		config:     config,
		httpClient: httpClient{},
	}
}

func (c Client) List(ctx context.Context) (*response.List, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.config.Timeout))
	defer cancel()

	url := fmt.Sprintf("%s/api/mods", c.config.UrlApi)
	resp, err := c.httpClient.get(ctx, url)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	decoder := json.NewDecoder(strings.NewReader(resp))
	asdf := response.List{}
	decoder.Decode(&asdf)

	return &asdf, nil
}

type httpClient struct {
}

func (client httpClient) get(ctx context.Context, url string) (string, error) {
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
