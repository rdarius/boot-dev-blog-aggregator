package main

import (
	"fmt"
	"github.com/rdarius/boot-dev-blog-aggregator/internal/config"
	"log"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	cfg.SetUser()

	cfg2, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", cfg2)

}
