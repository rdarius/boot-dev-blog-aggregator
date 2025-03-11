package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/database"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/parser"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/rss"
	"time"
)

func ScrapeFeedsHandler(s *config.State, cmd config.Command) error {

	ctx := context.Background()

	nextFeed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	_, err = s.DB.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return err
	}

	feedData, err := rss.FetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}

	fmt.Println("Feed Title: " + feedData.Channel.Title)

	for _, i := range feedData.Channel.Item {

		pubTime, err := parser.ParseTimestamp(i.PubDate)
		if err != nil {
			return err
		}
		_, err = s.DB.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			FeedID:      nextFeed.ID,
			Url:         i.Link,
			Title:       i.Title,
			Description: i.Description,
			PublishedAt: pubTime,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
