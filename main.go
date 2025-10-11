package main

import (
	"fmt"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"

	"example/sub-alert-server/controllers"
	_ "example/sub-alert-server/docs"
)

// @title Subscription Alerts And Tracker API
// @version 1.0
// @description This is the API Documentation for the subscription alerts and tracking server
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)
	
	app.Get("/users", controllers.GetUsers)
	app.Get("/user", controllers.GetUser)


	fmt.Println("🚀 Server running on http://localhost:8080/swagger/index.html")
	app.Listen(":8080")
}
