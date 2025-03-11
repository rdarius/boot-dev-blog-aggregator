package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/database"
	"os"
	"time"
)

func RegisterHandler(s *config.State, cmd config.Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("missing required argument USERNAME")
	}

	ctx := context.Background()

	_, err := s.DB.GetUser(ctx, cmd.Args[0])
	if err == nil {
		fmt.Println("User already exists")
		os.Exit(1)
	}

	u, err := s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		fmt.Printf("Failed to create user %v\n", err)
		os.Exit(1)
	}

	err = s.Config.SetUser(u.Name)
	if err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Printf("User Created %v\n", u)

	return nil
}
