package main

import (
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
