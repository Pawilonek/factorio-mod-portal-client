package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	config := Config{}
	client := New(config)

	list, err := client.List(ctx)
	if err != nil {
		panic(err)
	}

	for _, mod := range list.Results {
		fmt.Println("name:", mod.Name, "title:", mod.Title, "owner:", mod.Owner, "summary:", mod.Summary, "download count:", mod.DownloadCount, "category:", mod.Category, "score:", mod.Score)
	}

	fmt.Println(list.Pagination.Count)
	fmt.Println(list.Pagination.Links.First)
	fmt.Println(list.Pagination.Links.Next)
	fmt.Println(list.Pagination.Links.Last)
	fmt.Println(list.Pagination.Links.Prev)
	fmt.Println(list.Pagination.Page)
	fmt.Println(list.Pagination.PageCount)
	fmt.Println(list.Pagination.PageSize)
}
