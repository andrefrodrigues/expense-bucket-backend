package server

import (
	"expense-bucket-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() error {
	app := fiber.New()
	app.Use(logger.New())
	routes.Setup(app)
	err := app.Listen(":3000")

	return err
}
