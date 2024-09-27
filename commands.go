package main

import (
	"fmt"
)

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

func (c *Commands) Run(state *State, cmd Command) (err error) {
	command, ok := c.Commands[cmd.Name]
	if !ok {
		err = fmt.Errorf("the command, %v, has not been registered", cmd.Name)
		return
	}

	err = command(state, cmd)

	return
}
