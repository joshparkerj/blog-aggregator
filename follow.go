package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joshparkerj/blog-aggregator/internal/database"
)

func Follow(s *State, cmd Command) (err error) {
	feedUrl := cmd.Args[0]
	userName := s.Configuration.CurrentUserName
	fmt.Println(userName)
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       feedUrl,
		Name:      userName,
	}

	followFeed, err := s.DB.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return
	}

	fmt.Printf("Followed feed!\n * Feed name: %v\n * Current user: %v\n", followFeed.FeedName, followFeed.UserName)
	return
}
