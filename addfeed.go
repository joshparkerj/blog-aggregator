package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joshparkerj/blog-aggregator/internal/database"
)

func Addfeed(s *State, cmd Command, user database.User) (err error) {
	if len(cmd.Args) != 2 {
		for _, arg := range cmd.Args {
			fmt.Printf("got arg %v\n", arg)
		}

		err = errors.New("addfeed command requires two args (name and url)")
		return
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), params)
	if err != nil {
		return
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	s.DB.CreateFeedFollow(context.Background(), createFeedFollowParams)

	fmt.Printf("feed name: %v\nfeed id: %v\n feed created at: %v\nfeed updated at: %v\nfeed added by user with id: %v (user name is %v)\nfeed url: %v\n", feed.Name, feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.UserID, user.Name, feed.Url)

	return
}
