package handlers

import (
	"example/sub-alert-server/internals/models"
	"example/sub-alert-server/internals/services"

	"github.com/gofiber/fiber/v2"
)

// GetUsers godoc
// @Summary Get list of users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
	users := services.GetUsers()
	return c.JSON(users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user to the database
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdUser, err := services.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}
