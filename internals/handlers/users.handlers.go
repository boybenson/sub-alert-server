package handlers

import (
	"example/sub-alert-server/internals/services"

	"github.com/gofiber/fiber/v2"
)

// GetUsers godoc
// @Summary Get list of users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} types.User
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
	users := services.GetUsers()
	return c.JSON(users)
}

