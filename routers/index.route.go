package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shch989/Golang_Fiber_Project/config"
	"github.com/shch989/Golang_Fiber_Project/controllers"
)

func middleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" || token != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()
}

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"./public/asset")
	r.Get("/user", middleware, controllers.UserHandlerGetAll)
	r.Post("/user", controllers.UserHandlerCreate)
	r.Get("/user/:id", controllers.UserHandlerGetById)
	r.Put("/user/:id", controllers.UserHandlerUpdate)
	r.Delete("/user/:id", controllers.UserHandlerDelete)
}
