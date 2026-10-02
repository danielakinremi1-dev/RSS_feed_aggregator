package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/danielakinremi1-dev/RSS_feed_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit := 2
	if len(cmd.Args) > 0 {
		newLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Error parsing requested browse limit: %w", err)
		}
		limit = newLimit
	}

	queryArgs := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), queryArgs)
	if err != nil {
		return fmt.Errorf("Error getting user posts: %w", err)
	}

	for _, post := range posts {

		fmt.Printf("Post: %v\n Description: %v\n Published at %v from url %v",
			post.Title, post.Description, post.PublishedAt, post.Url)
	}

	return nil
}
