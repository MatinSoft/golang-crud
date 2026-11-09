package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maadco/go-crud-test/application"
	"github.com/maadco/go-crud-test/infrastructure"
	"github.com/maadco/go-crud-test/presentation"
)

// @title Go CRUD Clean Architecture CQRS API
// @version 1.0
// @description Customer CRUD service with Go, Gin, DDD, CQRS and Clean Architecture.
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
