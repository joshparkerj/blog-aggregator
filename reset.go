package main

import (
	"context"
)

func Reset(s *State, cmd Command) (err error) {
	err = s.DB.DeleteUsers(context.Background())
	return
}
