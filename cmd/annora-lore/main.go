package main

import (
	"log"

	"github.com/AdityaTaggar05/annora-lore/internal/app"
	"github.com/AdityaTaggar05/annora-lore/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	server, err := app.New(cfg)
	
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
