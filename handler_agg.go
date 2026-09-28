package main

import (
	"context"
	"fmt"
	"os"
)

func handlerAgg(s *state, cmd command) error {

	url := "https://www.wagslane.dev/index.xml"

	RSSFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		fmt.Println("Error fetching RSS feeds")
		os.Exit(1)
	}

	fmt.Printf("%+v\n", *RSSFeed)
	return nil
}
