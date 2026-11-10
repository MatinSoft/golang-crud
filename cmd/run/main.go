package main

import (
	"log"

	"github.com/maadco/go-crud-test/cmd/bootstrap"
	_ "github.com/maadco/go-crud-test/docs"
	"github.com/maadco/go-crud-test/internal/config"
)

// @title Customer CRUD API
// @version 1.0
// @description Customer CRUD API built with Gin, DDD, Clean Architecture and CQRS.
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	port := cfg.App.Port

	if port == "" {
		port = "8080"
	}

	app := bootstrap.New()

	log.Printf("Starting server on port %s...", port)
	app.Router.Run(":" + port)
}
