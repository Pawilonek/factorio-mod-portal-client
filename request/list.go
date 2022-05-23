package request

import (
	"net/url"
	"strconv"
	"strings"
)

const SortOrderAsc = "asc"
const SortOrderDesc = "desc"

const SortName = "name"
const SortCreatedAt = "created_at"
const SortUpdatedAt = "updated_at"

// Mod List Response
type ListParams struct {
	// Only return non-deprecated mods.
	HideDeprecated bool
	// Page number you would like to show. Makes it so you can see a certain part of the list without getting detail on all
	Page int
	// The amount of results to show in your search
	PageSize int
	// Sort results by this property. Defaults to name when not defined. Ignored for page_size=max queries.
	// {enum, one of name, created_at or updated_at}
	Sort string
	// Sort results ascending or descending. Defaults to descending when not defined. Ignored for page_size=max queries.
	// {enum, one of asc or desc}
	SortOrder string
	// Return only mods that match the given names.
	Namelist []string
	// Only return non-deprecated mods compatible with this Factorio version
	// {enum, one of 0.13, 0.14, 0.15, 0.16, 0.17, 0.18, 1.0 or 1.1}
	Version string
}

func (l *ListParams) GetAsUrlValues() url.Values {
	values := url.Values{}
	if l.HideDeprecated {
		values.Add("hide_deprecated", "true")
	}

	if l.Page > 0 {
		values.Add("page", strconv.Itoa(l.Page))
	}

	if l.PageSize > 0 {
		values.Add("page_size", strconv.Itoa(l.PageSize))
	}

	if l.Sort != "" {
		values.Add("sort", l.Sort)
	}

	if l.SortOrder != "" {
		values.Add("sort_order", l.SortOrder)
	}

	if len(l.Namelist) > 0 {
		namelist := strings.Join(l.Namelist, ",")
		values.Add("namelist", namelist)
	}

	if l.Version != "" {
		values.Add("version", l.Version)
	}

	return values
}
