package controllers

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/shch989/Golang_Fiber_Project/database"
	"github.com/shch989/Golang_Fiber_Project/models/entity"
	"github.com/shch989/Golang_Fiber_Project/models/request"
	"github.com/shch989/Golang_Fiber_Project/utils"
)

func LoginHandler(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	log.Println(loginRequest)

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"massage": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User
	err := database.DB.Where("email = ?", loginRequest.Email).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "wrong credential11111111",
			"error":   err.Error(),
		})
	}

	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(404).JSON(fiber.Map{
			"message": "wrong credential222222222",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return c.Status(404).JSON(fiber.Map{
			"message": "wrong credential333333333333",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
