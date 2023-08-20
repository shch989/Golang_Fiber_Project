package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shch989/Golang_Fiber_Project/config"
	"github.com/shch989/Golang_Fiber_Project/controllers"
	middleware "github.com/shch989/Golang_Fiber_Project/middlewares"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"./public/asset")
	// Auth
	r.Post("/login", controllers.LoginHandler)
	// User
	r.Get("/user", middleware.Auth, controllers.UserHandlerGetAll)
	r.Post("/user", controllers.UserHandlerCreate)
	r.Get("/user/:id", controllers.UserHandlerGetById)
	r.Put("/user/:id", controllers.UserHandlerUpdate)
	r.Delete("/user/:id", controllers.UserHandlerDelete)
}
