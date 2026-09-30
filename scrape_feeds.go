package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/danielakinremi1-dev/RSS_feed_aggregator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("Error fetching next feed to scrape")
		return err
	}

	currentTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	queryArgs := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: currentTime,
		ID:            nextFeed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), queryArgs)
	if err != nil {
		fmt.Println("Error marking feed as fetched")
		return err
	}

	RSSFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Println("Error fetching online feed data")
		return err
	}

	itemPublishedTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	itemDescription := sql.NullString{
		Time:  time.Now(),
		Valid: true,
	}

	for _, item := range RSSFeed.Channel.Item {
		post := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString,
			PublishedAt: item.PubDate,
			FeedID:      uuid.New(),
		}

	}

	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(RSSFeed.Channel.Item))

	return nil
}
