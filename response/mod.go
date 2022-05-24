package response

import "time"

type Release struct {
	// Path to download for a mod
	// todo: starts with "/download" and does not include a full url.
	DownloadUrl string `json:"download_url"`
	// The file name of the release. Always seems to follow the pattern "{name}_{version}.zip"
	FileName string `json:"file_name"`
	// A copy of the mod's info.json file, only contains factorio_version in short version
	InfoJosn map[string]string `json:"info_json"`
	// Time wen the mod was released.
	ReleasedAt time.Time `json:"released_at"`
	// The sha1 key for the file
	Sha1 string `json:"sh1"`
	// The version string of this mod release. Used to determine dependencies.
	Version string `json:"version"`
}

// Mod contains response for the mod details
type Mod struct {
	// The mod's machine-readable ID string.
	Name string `json:"name"`
	// The mod's human-readable name.
	Title string `json:"title"`
	// The Factorio username of the mod's author.
	Owner string `json:"owner"`
	// A shorter mod description.
	Summary string `json:"summary"`
	// Number of downloads.
	DownloadsCount int `json:"downloads_count"`
	// A single tag describing the mod.
	Category string `json:"category"`
	// todo: not described on wiki
	Score float32 `json:"score"`
	// Thumbnmail of the mod
	Thumbnail string `json:"thumbnail"`
	// List of all releases
	Releases []Release `json:"releases"`
}
