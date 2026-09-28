package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/danielakinremi1-dev/RSS_feed_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Insufficient command arguments give for adding feed")
	}

	feedName := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		fmt.Println("Error getting current user details")
		os.Exit(1)
	}

	queryArgs := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       url,
		UserID:    user.ID,
	}

	feedData, err := s.db.CreateFeed(context.Background(), queryArgs)
	if err != nil {
		fmt.Println("Error creating RSSFeed entry")
		os.Exit(1)
	}

	fmt.Println("RSSFeed created successfully!")
	fmt.Printf("%+v\n", feedData)
	return nil
}
