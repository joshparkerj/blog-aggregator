package main

import (
	"errors"
	"fmt"
	"time"
)

func Agg(s *State, cmd Command) (err error) {
	if len(cmd.Args) != 1 {
		err = errors.New("this command takes one parameter: time between requests (a duration string, like 1s, 1m, 1h)")
		return
	}

	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		err = fmt.Errorf("the parameter was not a valid duration format: got %v (%v)", cmd.Args[0], err)
		return
	}

	ticker := time.NewTicker(duration)
	fmt.Printf("Collecting feeds every %v\n", duration)
	for {
		err = ScrapeFeeds(s)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				fmt.Println("ERROR: try adding some feeds to aggregate first")
			}

			return
		}

		<-ticker.C
	}
}
