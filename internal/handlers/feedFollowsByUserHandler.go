package handlers

import (
	"context"
	"fmt"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
)

func GetFeedFollowsByUserHandler(s *config.State, cmd config.Command) error {
	ctx := context.Background()
	user, err := s.DB.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	ff, err := s.DB.GetFeedFollowsByUser(ctx, user.ID)

	if err != nil {
		return err
	}

	for _, feed := range ff {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}
