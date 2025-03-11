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

func AddFeedHandler(s *config.State, cmd config.Command) error {

	if len(cmd.Args) < 2 {
		log.Fatal("usage: boot-dev-blog-aggregator addfeed NAME URL")
	}

	ctx := context.Background()
	user, err := s.DB.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	_, err = rss.FetchFeed(ctx, url)
	if err != nil {
		fmt.Printf("Failed to fetch feed: %v\n", err)
		os.Exit(1)
	}

	f, err := s.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.UUID{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		UserID:    user.ID,
		Url:       url,
		Name:      name,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%v", f)

	return nil
}
