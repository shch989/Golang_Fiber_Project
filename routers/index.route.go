package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shch989/Golang_Fiber_Project/controllers"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", controllers.UserHandlerGetAll)
	r.Post("/user", controllers.UserHandlerCreate)
}
