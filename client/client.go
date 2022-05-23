package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Pawilonek/factorio-mod-portal-client/http"
	"github.com/Pawilonek/factorio-mod-portal-client/request"
	"github.com/Pawilonek/factorio-mod-portal-client/response"
)

const urlApi = "https://mods.factorio.com"
const urlAssets = "https://assets-mod.factorio.com"

type Config struct {
	Timeout   time.Duration
	UrlApi    string
	UrlAssets string
}

type Client struct {
	config     Config
	httpClient http.Client
}

// New creats a new instance of Client based on passed dependencies or applying defauls
func New(passedConfig *Config, httpClient http.Client) *Client {
	// Copy passed config to prevent overriding usesr data
	var config Config
	if passedConfig == nil {
		config = Config{}
	} else {
		config = *passedConfig
	}

	// Set default values
	if config.Timeout == 0 {
		config.Timeout = time.Second
	}

	if config.UrlApi == "" {
		config.UrlApi = urlApi
	}

	if config.UrlAssets == "" {
		config.UrlAssets = urlAssets
	}

	if httpClient == nil {
		httpClient = &http.HttpClient{}
	}

	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

// Returns paginates list of mods
func (c Client) List(ctx context.Context, params *request.ListParams) (*response.List, error) {
	if params == nil {
		params = &request.ListParams{}
	}

	urlParams := params.GetAsUrlValues()
	urlParamsEncoded := urlParams.Encode()
	if len(urlParamsEncoded) > 0 {
		urlParamsEncoded = "?" + urlParamsEncoded
	}

	url := fmt.Sprintf("%s/api/mods%s", c.config.UrlApi, urlParamsEncoded)
	requestCtx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	resp, err := c.httpClient.Get(requestCtx, url)
	cancel()
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(strings.NewReader(resp))
	asdf := response.List{}
	err = decoder.Decode(&asdf)
	if err != nil {
		return nil, fmt.Errorf("coudn't decode response from list, err: %s", err)
	}

	return &asdf, nil
}
