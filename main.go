package main

import (
	"context"
	"fmt"

	factorio_client "github.com/Pawilonek/factorio-mod-portal-client/client"
	"github.com/Pawilonek/factorio-mod-portal-client/request"
)

func main() {
	ctx := context.Background()

	config := factorio_client.Config{}
	client := factorio_client.New(&config, nil)

	list, err := client.List(ctx, &request.ListParams{
		Page:           1,
		PageSize:       2,
		HideDeprecated: true,
		Sort:           request.SortUpdatedAt,
		SortOrder:      request.SortOrderDesc,
		Version:        "1.1",
		Namelist: []string{
			"Atomic_Overhaul",
			"factoryplanner",
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("::::: List")

	for key, mod := range list.Results {
		fmt.Println("::::: List:", key)
		fmt.Println("name:", mod.Name, "title:", mod.Title, "owner:", mod.Owner, "summary:", mod.Summary, "download count:", mod.DownloadsCount, "category:", mod.Category, "score:", mod.Score)
	}

	fmt.Println("::::: Get")

	mod, err := client.Get(ctx, "Krastorio2")
	if err != nil {
		panic(err)
	}

	fmt.Println("name:", mod.Name, "title:", mod.Title, "owner:", mod.Owner, "summary:", mod.Summary, "download count:", mod.DownloadsCount, "category:", mod.Category, "score:", mod.Score)

}
