package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maadco/go-crud-test/application"
	_ "github.com/maadco/go-crud-test/docs"
	"github.com/maadco/go-crud-test/infrastructure"
	"github.com/maadco/go-crud-test/presentation"
)

// @title Customer CRUD API
// @version 1.0
// @description Customer CRUD API built with Gin, DDD, Clean Architecture and CQRS.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	repo := infrastructure.NewInMemoryCustomerRepository()
	eventStore := infrastructure.NewInMemoryEventStore()
	service := application.NewCustomerService(repo, eventStore)

	presentation.RegisterRoutes(r, service)

	r.Run(":8080")
}
