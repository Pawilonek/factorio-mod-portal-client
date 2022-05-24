package http

import (
	"context"
	_ "embed"
	"fmt"
)

//go:embed mocked_responses/api_mods.json
var apiMods string

//go:embed mocked_responses/api_mods_page_2.json
var apiModsPage2 string

//go:embed mocked_responses/api_mods_krastorio2.json
var apiModsKrastorio string

var mockedPages = map[string]string{
	"https://mods.factorio.com/api/mods":            apiMods,
	"https://mods.factorio.com/api/mods?page=2":     apiModsPage2,
	"https://mods.factorio.com/api/mods/Krastorio2": apiModsKrastorio,
}

type MockedClient struct {
}

func (client *MockedClient) Get(ctx context.Context, url string) (string, error) {
	response, ok := mockedPages[url]
	if !ok {
		return "", fmt.Errorf("Not found")
	}

	return response, nil
}
