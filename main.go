package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/database"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/handlers"
	"log"
	"os"
)

var commands config.Commands

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)

	state := config.State{
		DB:     dbQueries,
		Config: &cfg,
	}

	commands = config.Commands{
		Commands: map[string]func(*config.State, config.Command) error{},
	}

	commands.Register("login", handlers.LoginHandler)
	commands.Register("register", handlers.RegisterHandler)
	commands.Register("reset", handlers.ResetUsersHandler)
	commands.Register("users", handlers.GetUsersHandler)
	commands.Register("agg", handlers.FetchFeedHandler)
	commands.Register("addfeed", middlewareLoggedIn(handlers.AddFeedHandler))
	commands.Register("feeds", handlers.ListFeedsHandler)
	commands.Register("follow", middlewareLoggedIn(handlers.FollowFeedHandler))
	commands.Register("unfollow", middlewareLoggedIn(handlers.UnfollowFeedHandler))
	commands.Register("following", middlewareLoggedIn(handlers.GetFeedFollowsByUserHandler))

	if len(os.Args) < 2 {
		log.Fatal("usage: boot-dev-blog-aggregator <command> [args...]")
		return
	}

	name := os.Args[1]
	args := os.Args[2:]

	err = commands.Run(&state, config.Command{Name: name, Args: args})
	if err != nil {
		log.Fatal(err)
	}
}

func middlewareLoggedIn(handler func(s *config.State, cmd config.Command, user database.User) error) func(*config.State, config.Command) error {
	return func(state *config.State, command config.Command) error {
		user, err := state.DB.GetUser(context.Background(), state.Config.CurrentUserName)
		if err != nil {
			log.Fatal("user not logged in")
		}
		return handler(state, command, user)
	}
}
