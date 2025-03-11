package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"os"
)

func LoginHandler(s *config.State, cmd config.Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("missing required argument USERNAME")
	}

	ctx := context.Background()

	u, err := s.DB.GetUser(ctx, cmd.Args[0])
	if err != nil {
		fmt.Println("User does not exist")
		os.Exit(1)
	}

	err = s.Config.SetUser(u.Name)
	if err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Printf("User set to %s\n", s.Config.CurrentUserName)

	return nil
}
