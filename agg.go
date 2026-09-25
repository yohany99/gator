package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	rssfeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", rssfeed)
	return nil
}

func scrapeFeeds(s *state) error {
	return nil
}
