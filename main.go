package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "example/sub-alert-server/docs"
)

// @title Fiber Example API
// @version 1.0
// @description Sample Swagger for Fiber
// @host localhost:8080
// @BasePath /
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Benson", Age: 25},
}

func main() {
	app := fiber.New()

  app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	fmt.Println("🚀 Server running on http://localhost:8080/swagger")
	app.Listen(":8080")
}
