package main

import (
	"fmt"
	"log"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"

	_ "example/sub-alert-server/docs"
	"example/sub-alert-server/internals/database"
	"example/sub-alert-server/internals/handlers"

	"github.com/joho/godotenv"
)

// @title Subscription Alerts And Tracker API
// @version 1.0
// @description This is the API Documentation for the subscription alerts and tracking server
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()


	database.ConnectDB()
	database.InitCollections()

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/users", handlers.GetUsers)


	fmt.Println("🚀 Server running on http://localhost:8080/swagger/index.html")
	app.Listen(":8080")
}
