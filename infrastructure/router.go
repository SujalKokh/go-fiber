package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Router -> Fiber Router
type Router struct {
	*fiber.App
}

//NewRouter : all the routes are defined here
func NewRouter(env Env) Router {

	httpRouter := fiber.New()

	httpRouter.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "http://localhost",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	httpRouter.Use(logger.New())

	httpRouter.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Fiber API up and running ♾️",
		})
	})

	return Router{
		httpRouter,
	}
}
