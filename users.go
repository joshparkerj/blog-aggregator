package main

import (
	"context"
	"fmt"
)

func Users(s *State, cmd Command) (err error) {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return
	}

	for _, user := range users {
		if user.Name == s.Configuration.CurrentUserName {
			fmt.Printf(" * %v (current)\n", user.Name)
		} else {
			fmt.Printf(" * %v\n", user.Name)
		}
	}
	return
}
