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

func UserHandlerUpdate(c *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user entity.User

	userId := c.Params("id")

	err := database.DB.Where("id = ?", userId).First(&user).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		c.Status(500).JSON(fiber.Map{
			"message": "failed",
			"error":   errUpdate.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}
