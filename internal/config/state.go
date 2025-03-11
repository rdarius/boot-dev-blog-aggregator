package config

import "github.com/rdarius/boot-dev-blog-aggregator/internal/database"

type State struct {
	DB     *database.Queries
	Config *Config
}
