package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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
	fmt.Printf("got %v items\n", len(fetchedFeed.Channel.Item))
	timeLayout := "Mon, 02 Jan 2006 15:04:04 -0700"
	for _, item := range fetchedFeed.Channel.Item {
		publishTime, err := time.Parse(timeLayout, item.PubDate)
		if err != nil {
			fmt.Printf("the time format was not parseable: %v\n", item.PubDate)
			fmt.Println(err)
		}

		timeValid := err == nil
		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title,
				Valid:  true,
			},
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url: item.Link,
			PublishedAt: sql.NullTime{
				Time:  publishTime,
				Valid: timeValid,
			},
			FeedID: feed.ID,
		}

		_, err = s.DB.CreatePost(context.Background(), postParams)
		if err != nil {
			// handle some errors and ignore some here
			if err.Error() != "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
				log.Fatal(err)
			}
		}
	}

	return
}
