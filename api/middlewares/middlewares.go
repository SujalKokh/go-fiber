package middlewares

import "go.uber.org/fx"

var Module = fx.Options()

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares() Middlewares {
	return Middlewares{}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
