package handlers

import (
	"context"
	"fmt"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"os"
)

func GetUsersHandler(s *config.State, cmd config.Command) error {
	ctx := context.Background()

	u, err := s.DB.GetUsers(ctx)
	if err != nil {
		fmt.Println("Failed to get Users")
		os.Exit(1)
	}

	for _, u := range u {
		fmt.Printf("* %s", u.Name)

		if u.Name == s.Config.CurrentUserName {
			fmt.Print(" (current)")
		}

		fmt.Print("\n")
	}

	return nil
}
