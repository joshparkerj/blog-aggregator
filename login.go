package main

import (
	"errors"
	"fmt"
)

func Login(s *State, cmd Command) (err error) {
	if len(cmd.Args) == 0 {
		err = errors.New("need a username to login")
		return
	}

	err = s.configuration.SetUser(cmd.Args[0])
	if err != nil {
		return
	}

	fmt.Println("user has been set")
	return
}
