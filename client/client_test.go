package client_test

import (
	"context"
	"testing"
	"time"

	factorio_client "github.com/Pawilonek/factorio-mod-portal-client/client"
	"github.com/Pawilonek/factorio-mod-portal-client/http"
)

func TestMappedValues(t *testing.T) {
	ctx := context.TODO()

	mockedClient := &http.MockedClient{}
	client := factorio_client.New(nil, mockedClient)

	response, err := client.List(ctx)
	if err != nil {
		t.Fatalf("List should not return an error, %s", err)
	}

	// Check if date is properly converted to golang type
	lastReleaseDate := response.Results[0].LatestRelease.ReleasedAt
	explectedDate := time.Date(2019, 4, 2, 1, 22, 18, int(431*time.Millisecond), time.UTC)
	if !lastReleaseDate.Equal(explectedDate) {
		t.Fatalf("Expected value %v, but got: %v", explectedDate, lastReleaseDate)
	}

	expectedName := "120U-235"
	if response.Results[23].Name != expectedName {
		t.Fatalf("Expected value %s, but got: %s", expectedName, response.Results[23].Name)
	}

	expectedDownloadCount := 1284
	if response.Results[23].DownloadsCount != expectedDownloadCount {
		t.Fatalf("Expected value %d, but got: %d", expectedDownloadCount, response.Results[23].DownloadsCount)
	}

	expectedNumberOfPages := 384
	if response.Pagination.PageCount != expectedNumberOfPages {
		t.Fatalf("Expected value %d, but got: %d", expectedNumberOfPages, response.Pagination.PageCount)
	}

	expectedCount := 9579
	if response.Pagination.Count != expectedCount {
		t.Fatalf("Expected value %d, but got: %d", expectedCount, response.Pagination.Count)
	}
}
