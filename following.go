package main

import (
	"context"
	"fmt"

	"github.com/joshparkerj/blog-aggregator/internal/database"
)

func Following(s *State, cmd Command, user database.User) (err error) {
	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return
	}

	for _, feed := range feeds {
		fmt.Printf(" ********\n * Feed name: %v\n * Feed url: %v\n * Feed created by: %v\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}
	return
}
