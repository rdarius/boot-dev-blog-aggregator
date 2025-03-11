package handlers

import (
	"context"
	"fmt"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/database"
	"log"
)

func UnfollowFeedHandler(s *config.State, cmd config.Command, user database.User) error {
	if len(cmd.Args) < 1 {
		log.Fatal("usage: boot-dev-blog-aggregator unfollow URL")
	}
	ctx := context.Background()

	url := cmd.Args[0]

	err := s.DB.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	})

	if err != nil {
		return err
	}

	fmt.Println("Feed unfollow successfully")

	return nil
}
