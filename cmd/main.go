package main

import (
	"GO-FJ/internal/config"
	"GO-FJ/internal/server"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
		return
	}

	server.New(cfg).Run()
}
