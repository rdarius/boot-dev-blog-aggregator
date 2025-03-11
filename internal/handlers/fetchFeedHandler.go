package handlers

import (
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"log"
	"time"
)

func FetchFeedHandler(s *config.State, cmd config.Command) error {
	if len(cmd.Args) < 1 {
		log.Fatal("usage: boot-dev-blog-aggregator agg time_between_reqs")
	}

	t, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(t)
	for ; ; <-ticker.C {
		err := ScrapeFeedsHandler(s, cmd)
		if err != nil {
			return err
		}
	}

	return nil
}
