package handlers

import (
	"context"
	"fmt"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/database"
	"strconv"
)

func BrowseHandler(s *config.State, cmd config.Command, user database.User) error {
	var limit int32
	if len(cmd.Args) < 1 {
		limit = 2
	} else {
		l, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = int32(l)
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\nURL: %s\nArticle: %sPublished: %s\n\n", post.Title, post.Url, post.Description, post.PublishedAt)
	}

	return nil
}
