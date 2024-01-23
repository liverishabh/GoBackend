package router

import (
	"fiber-demo/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupV1Routes(api fiber.Router) {
	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
}
