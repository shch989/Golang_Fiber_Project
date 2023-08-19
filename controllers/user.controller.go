package controllers

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/shch989/Golang_Fiber_Project/database"
	"github.com/shch989/Golang_Fiber_Project/models/entity"
	"github.com/shch989/Golang_Fiber_Project/models/request"
)

func UserHandlerGetAll(c *fiber.Ctx) error {
	var users []entity.User

	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}

	return c.JSON(users)
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to store data.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entity.User
	err := database.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}
