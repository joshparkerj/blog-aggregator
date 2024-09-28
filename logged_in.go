package main

import (
	"context"

	"github.com/joshparkerj/blog-aggregator/internal/database"
)

func MiddlewareLoggedIn(
	handler func(s *State, cmd Command, user database.User) error,
) (loggedInHandler func(*State, Command) error) {

	loggedInHandler = func(s *State, cmd Command) (err error) {
		user, err := s.DB.GetUser(context.Background(), s.Configuration.CurrentUserName)
		if err != nil {
			return
		}
		err = handler(s, cmd, user)
		return
	}

	return
}
