package main

import (
	"log"

	"github.com/rm-ryou/sample_todo_app/internal/api"
	"github.com/rm-ryou/sample_todo_app/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	api.Run(cfg)
}
