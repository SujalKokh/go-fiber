package controllers

import (
	"fmt"
	"go-fiber/api/responses"
	"go-fiber/infrastructure"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json: "name"`
	Email string `json: "email"`
}

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
func (controller UsersController) HandleGetAllUsers(
	c *fiber.Ctx,
) error {
	return responses.JSON(c, http.StatusOK, "hello")
}

func (controller UsersController) CreateNewUser(
	c *fiber.Ctx,
) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		fmt.Println("error = ", err)
		return responses.JSON(c, http.StatusInternalServerError, "failed to parse json body")
	}
	// getting user if no error
	fmt.Println("user = ", user)
	fmt.Println("user = ", user.Name)
	fmt.Println("user = ", user.Email)
	return responses.JSON(c, http.StatusOK, "test")
}
