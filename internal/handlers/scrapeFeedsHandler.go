package handlers

import (
	"context"
	"fmt"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/rss"
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
		fmt.Println("Item Title: " + i.Title)
	}

	return nil
}
