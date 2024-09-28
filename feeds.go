package main

import (
	"context"
	"fmt"
)

func Feeds(s *State, cmd Command) (err error) {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return
	}

	for _, feed := range feeds {
		fmt.Printf(" ********\n * Feed name: %v\n * Feed url: %v\n * Feed created by: %v\n", feed.Name, feed.Url, feed.Username)
	}
	return
}
