package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) error {
	api := app.Group("/api")
	err := SetupAuthRoutes(api)
	if err != nil {
		return err
	}

	return nil
}
