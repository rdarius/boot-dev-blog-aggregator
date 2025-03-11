package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/database"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/rss"
	"log"
	"os"
	"time"
)

func FollowFeedHandler(s *config.State, cmd config.Command, user database.User) error {
	if len(cmd.Args) < 1 {
		log.Fatal("usage: boot-dev-blog-aggregator follow URL")
	}
	ctx := context.Background()

	url := cmd.Args[0]

	ff, err := s.DB.GetFeedByUrl(ctx, url)
	if err != nil {
		// no feed, get it and save it

		fd, err := rss.FetchFeed(ctx, url)
		if err != nil {
			fmt.Printf("Failed to fetch feed: %v\n", err)
			os.Exit(1)
		}

		res, err := s.DB.CreateFeed(ctx, database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			UserID:    user.ID,
			Url:       url,
			Name:      fd.Channel.Title,
		})
		if err != nil {
			fmt.Printf("Failed to create feed: %v\n", err)
			os.Exit(1)
		}
		ff = res
	}

	feed, err := s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		UserID:    user.ID,
		FeedID:    ff.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%v", feed)

	return nil
}
