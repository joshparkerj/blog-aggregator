package main

import (
	"context"
	"fmt"
)

func Following(s *State, cmd Command) (err error) {
	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), s.Configuration.CurrentUserName)
	if err != nil {
		return
	}

	for _, feed := range feeds {
		fmt.Printf(" ********\n * Feed name: %v\n * Feed url: %v\n * Feed created by: %v\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}
	return
}
