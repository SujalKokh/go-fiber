package controllers

import (
	"go-fiber/infrastructure"

	"github.com/gofiber/fiber/v2"
)

// UsersController -> Structure for Users Controller
type UsersController struct {
	logger infrastructure.Logger
}

// NewUsersController -> constructor function for users controller
func NewUsersController(
	logger infrastructure.Logger,
) UsersController {
	return UsersController{
		logger: logger,
	}
}

// HandleGetAllUsers -> get all the users
func (controller UsersController) HandleGetAllUsers(c *fiber.Ctx) error {
	c.SendString("Test response")
	return nil
}
