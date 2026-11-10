package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maadco/go-crud-test/application"
	"github.com/maadco/go-crud-test/infrastructure"
	"github.com/maadco/go-crud-test/presentation"
)

func main() {
	r := gin.Default()

	repo := infrastructure.NewInMemoryCustomerRepository()
	eventStore := infrastructure.NewInMemoryEventStore()
	service := application.NewCustomerService(repo, eventStore)

	presentation.RegisterRoutes(r, service)

	r.Run(":8080")
}
