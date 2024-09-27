package main

import (
	"context"
	"errors"
	"fmt"
)

func Login(s *State, cmd Command) (err error) {
	if len(cmd.Args) == 0 {
		err = errors.New("need a username to login")
		return
	}

	user, err := s.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		err = fmt.Errorf("could not get user (%v)", err)
		return
	}

	err = s.Configuration.SetUser(user.Name)
	if err != nil {
		return
	}

	fmt.Println("user has been set")
	return
}
