package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, error := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if error != nil {
		return nil, error
	}
	request.Header.Set("User-Agent", "gator")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	response, error := client.Do(request)
	if error != nil {
		return nil, error
	}
	io.ReadAll(response)
	return nil, nil
}
