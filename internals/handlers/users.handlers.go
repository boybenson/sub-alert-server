package handlers

import "github.com/gofiber/fiber/v2"


type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Benson", Age: 25},
}

// GetUsers godoc
// @Summary Get list of users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} User
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}


// GetUser godoc
// @Summary Get a user
// @Description Get a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Router /user [get]
func GetUser(c *fiber.Ctx) error {
	return c.JSON(users[0])
}
