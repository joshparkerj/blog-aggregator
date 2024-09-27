package main

import (
	"context"
	"fmt"
)

func Agg(s *State, cmd Command) (err error) {
	// this link is hardcoded for now
	link := "https://www.wagslane.dev/index.xml"
	feed, err := FetchFeed(context.Background(), link)
	fmt.Println(feed)
	return
}
