package routes

import (
	"expense-bucket-api/auth"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) error {
	api := app.Group("/api")
	err := auth.SetupRoutes(api)
	if err != nil {
		return err
	}

	return nil
}
