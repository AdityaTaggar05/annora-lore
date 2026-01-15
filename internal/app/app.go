package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/config"
	"github.com/AdityaTaggar05/annora-lore/internal/handlers/https"
	"github.com/AdityaTaggar05/annora-lore/internal/infrastructure/neo4j"
	"github.com/AdityaTaggar05/annora-lore/internal/repository"
	"github.com/AdityaTaggar05/annora-lore/internal/services"
)

type App struct {
	Config *config.Config
	Server *http.Server
}

func New(cfg *config.Config) (*App, error) {
	ctx := context.Background()
	neoDB, err := neo4j.NewNeo4jDB(ctx, &cfg.Neo4j)

	if err != nil {
		log.Fatalf("error: couldn't connect to Neo4j")
		os.Exit(1)
	}
	log.Println("connected succesfully to Neo4j")

	loreRepo := repository.NewLoreRepository(neoDB)
	loreService := services.NewLoreService(loreRepo)
	loreHandler := https.NewHandler(loreService)

	router := https.NewRouter(*loreHandler)

	return &App{
		Config: cfg,
		Server: &http.Server{
			Addr:         ":" + cfg.Server.Port,
			Handler:      router,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
		},
	}, nil
}

func (a *App) Start() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Lore service listening on %s\n", a.Server.Addr)

		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v\n", err)
		}
	}()

	<-stop
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
		return err
	}

	log.Println("Auth service stopped gracefully")
	return nil
}
