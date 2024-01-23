package handler

import (
	"errors"
	"fiber-demo/database"
	"fiber-demo/models"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var user models.User
	result := db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"successful": false,
			"message":    fmt.Sprintf("No user found with given id: %s", id),
		})
	}
	return c.JSON(fiber.Map{
		"successful": true,
		"data":       user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"successful": false,
			"message":    "Invalid input",
			"data":       err.Error(),
		})
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"successful": false,
			"message":    "Invalid input",
			"data":       err.Error(),
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"successful": false,
			"message":    "Something went wrong",
		})
	}
	user.Password = string(hashedPassword)

	if result := db.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"successful": false,
			"message":    "Something went wrong",
			"data":       err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"successful": true,
		"message":    "User created successfully",
	})
}
