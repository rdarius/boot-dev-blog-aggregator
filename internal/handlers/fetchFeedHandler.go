package handlers

import (
	"context"
	"fmt"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/rss"
	"os"
)

func FetchFeedHandler(s *config.State, cmd config.Command) error {
	var feedUrl string
	if len(cmd.Args) < 1 {
		feedUrl = "https://www.wagslane.dev/index.xml"
	} else {
		feedUrl = cmd.Args[0]
	}

	ctx := context.Background()

	feed, err := rss.FetchFeed(ctx, feedUrl)
	if err != nil {
		fmt.Println("Failed to fetch feed")
		os.Exit(1)
	}

	fmt.Printf("%v", feed)

	return nil
}
