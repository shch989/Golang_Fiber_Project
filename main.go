package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	route "github.com/shch989/Golang_Fiber_Project/routers"
)

func main() {
	app := fiber.New()

	// Install Route
	route.RouteInit(app)

	log.Fatal(app.Listen(":8080"))
}
