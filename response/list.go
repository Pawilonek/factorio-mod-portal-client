package response

import (
	"time"
)

type Links struct {
	// URL to the first page of the results, or null if you're already on the first page.
	First string `json:"first"`
	// URL to the previous page of the results, or null if you're already on the first page.
	Prev string `json:"prev"`
	// URL to the next page of the results, or null if you're already on the last page.
	Next string `json:"next"`
	// URL to the last page of the results, or null if you're already on the last page.
	Last string `json:"last"`
}

type Pagination struct {
	// Total number of mods that match your specified filters.
	Count int `json:"count"`
	// The current page number.
	Page int `json:"page"`
	// The total number of pages returned.
	PageCount int `json:"page_count"`
	// The number of results per page.
	PageSize int `json:"page_size"`
	// Utility links to mod portal api requests, preserving all filters and search queries.
	Links Links `json:"links"`
}

type InfoJsonShort struct {
	FactorioVersion string `json:"factorio_version"`
}

type LatestRelease struct {
	// Path to download for a mod. starts with "/download" and does not include a full url
	DownloadUrl string `json:"download_url"`
	// The file name of the release.
	FileName string `json:"file_name"`
	// A copy of the mod's info.json file, only contains factorio_version in short version
	InfoJson InfoJsonShort `json:"info_json"`
	// Time when the mod was released.
	ReleasedAt time.Time `json:"released_at"`
	// The version string of this mod release. Used to determine dependencies.
	Version string `json:"version"`
	// The sha1 key for the file
	Sha1 string `json:"sha1"`
}

type ModList struct {
	// The mod's machine-readable ID string.
	Name string `json:"name"`
	// The mod's human-readable name.
	Title string `json:"title"`
	// The Factorio username of the mod's author.
	Owner string `json:"owner"`
	// A shorter mod description.
	Summary string `json:"summary"`
	// Number of downloads.
	DownloadCount int `json:"download_count"`
	// A single tag describing the mod.
	Category string `json:"category"`
	// todo: not described on wiki
	Score float32 `json:"score"`
	// The latest version of the mod available for download.
	LatestRelease LatestRelease `json:"latest_release"`
}

type List struct {
	// Pagination
	Pagination Pagination `json:"pagination"`
	// A list of mods, matching any filters you specified.
	Results []ModList `json:"results"`
}
