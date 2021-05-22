package routes

import "go.uber.org/fx"

// Module exports dependency of routes
var Module = fx.Options()

type Routes []Routes

type Route interface {
	Setup()
}

//NewRoute setups the routes
func NewRoute() Routes {
	return Routes{}
}

// Setup all the routes
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
