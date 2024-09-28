package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/joshparkerj/blog-aggregator/internal/database"
)

func ScrapeFeeds(s *State) (err error) {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return
	}

	params := database.MarkFeedFetchedParams{
		Url: feed.Url,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err = s.DB.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return
	}

	fetchedFeed, err := FetchFeed(context.Background(), feed.Url)

	fmt.Printf(" **** **** **** **** %v %v **** **** **** ****\n", feed.Name, time.Now().Format("03:04:05"))
	for _, item := range fetchedFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	return
}
