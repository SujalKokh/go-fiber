package routes

import (
	"go.uber.org/fx"
)

// Module exports dependency of routes
var Module = fx.Options(
	fx.Provide(NewUsersRoute),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

//NewRoute setups the routes
func NewRoutes(
	users UsersRoute,
) Routes {
	return Routes{
		users,
	}
}

// Setup all the routes
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
