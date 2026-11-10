package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/maadco/go-crud-test/internal/application"
	"github.com/maadco/go-crud-test/internal/infrastructure"
	"github.com/maadco/go-crud-test/internal/presentation"
)

// App holds the wired dependencies of the application.
type App struct {
	Router          *gin.Engine
	CustomerService *application.CustomerService
}

// New wires the infrastructure, application and presentation layers together.
func New() *App {
	r := gin.Default()

	repo := infrastructure.NewInMemoryCustomerRepository()
	eventStore := infrastructure.NewInMemoryEventStore()
	service := application.NewCustomerService(repo, eventStore)

	presentation.RegisterRoutes(r, service)

	return &App{
		Router:          r,
		CustomerService: service,
	}
}
