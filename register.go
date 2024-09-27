package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joshparkerj/blog-aggregator/internal/database"
)

func Register(s *State, cmd Command) (err error) {
	if len(cmd.Args) == 0 {
		err = fmt.Errorf("not enough args for command %v", cmd.Name)
		return
	}

	id := uuid.New()
	createUserParams := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := state.DB.CreateUser(context.Background(), createUserParams)
	if err != nil {
		err = fmt.Errorf("could not create a new user! (%v)", err)
		return
	}

	err = s.Configuration.SetUser(user.Name)
	if err != nil {
		return
	}

	fmt.Printf("User was created!\nUsername: %v\nUser id: %v\n User created at: %v\n User updated at: %v\n",
		user.Name,
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return
}
