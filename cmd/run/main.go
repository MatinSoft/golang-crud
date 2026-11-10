package main

import (
	"github.com/maadco/go-crud-test/cmd/bootstrap"
	_ "github.com/maadco/go-crud-test/docs"
)

// @title Customer CRUD API
// @version 1.0
// @description Customer CRUD API built with Gin, DDD, Clean Architecture and CQRS.
// @host localhost:8080
// @BasePath /
func main() {
	app := bootstrap.New()

	app.Router.Run(":8080")
}
