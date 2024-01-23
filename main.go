package main

import (
	"fiber-demo/config"
	"fiber-demo/database"
	"fiber-demo/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber Demo",
	})

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	config.LoadConfig()
	database.ConnectDB()

	api := app.Group("/api")
	apiV1 := api.Group("/v1")
	router.SetupV1Routes(apiV1)

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
