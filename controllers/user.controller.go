package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shch989/Golang_Fiber_Project/database"
	"github.com/shch989/Golang_Fiber_Project/models/entity"
)

func UserHandlerGetAll(c *fiber.Ctx) error {
	var users []entity.User

	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}

	return c.JSON(users)
}
