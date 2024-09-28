package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/joshparkerj/blog-aggregator/internal/database"
)

func Browse(s *State, cmd Command, user database.User) (err error) {
	limit := 2
	if len(cmd.Args) >= 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Printf("could not parse %v as int. using default limit of 2", cmd.Args[0])
			limit = 2
		}

		if len(cmd.Args) > 1 {
			fmt.Printf("too many args. discarding these: %v\n", cmd.Args[1:])
		}
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), params)
	if err != nil {
		return
	}

	for _, post := range posts {
		fmt.Printf(" ********\n * Post title: %v\n * Post url: %v\n * Post description: %v\n", post.Title, post.Url, post.Description)
	}
	return
}
