package main

import (
	"context"
	"fmt"

	"github.com/joshparkerj/blog-aggregator/internal/database"
)

func Unfollow(s *State, cmd Command, user database.User) (err error) {
	feedUrl := cmd.Args[0]
	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    feedUrl,
	}

	err = s.DB.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return
	}

	fmt.Printf("Unfollowed feed!\n")
	return
}
