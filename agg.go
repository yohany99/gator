package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}
	timeBetweenReqs := cmd.args[0]
	timeDuration, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	fmt.Printf("Collecting feeds every %s...", timeDuration)
	ticker := time.NewTicker(timeDuration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	// rssfeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return fmt.Errorf("couldn't fetch feed: %w", err)
	// }
	// fmt.Printf("Feed: %+v\n", rssfeed)
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed to fetch: %w", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed %s fetched: %v", feed.Name, err)
	}
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't collect feed %s: %v", feed.Name, err)
	}
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	fmt.Printf("Feed %s collected, %v posts found\n", feed.Name, len(feedData.Channel.Item))
	return nil
}
