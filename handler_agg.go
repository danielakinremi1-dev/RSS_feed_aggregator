package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("Insufficient command arguments given for aggregating feeds")
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		fmt.Println("Error parsing duration between aggregation requests")
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}
