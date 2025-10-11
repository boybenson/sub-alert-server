package main

import (
	"fmt"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"

	_ "example/sub-alert-server/docs"
	"example/sub-alert-server/internals/handlers"
)

// @title Subscription Alerts And Tracker API
// @version 1.0
// @description This is the API Documentation for the subscription alerts and tracking server
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)
	
	app.Get("/users", handlers.GetUsers)
	app.Get("/user", handlers.GetUser)


	fmt.Println("🚀 Server running on http://localhost:8080/swagger/index.html")
	app.Listen(":8080")
}
