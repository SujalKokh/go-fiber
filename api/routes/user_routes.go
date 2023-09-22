package routes

import (
	"go-fiber/api/controllers"
	"go-fiber/infrastructure"
)

// UserRoute -> structure for user routes
type UsersRoute struct {
	controller controllers.UsersController
	router     infrastructure.Router
}

// NewUserRoute -> constructor function for user routes
func NewUsersRoute(
	controller controllers.UsersController,
	router infrastructure.Router,
) UsersRoute {
	return UsersRoute{
		controller: controller,
		router:     router,
	}
}

// Setup -> setups the routes
func (u UsersRoute) Setup() {
	users := u.router.Group("/users")
	users.Get("/", u.controller.HandleGetAllUsers)
	users.Post("/", u.controller.CreateNewUser)
}
