package routes

import (
	"expense-bucket-api/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) error {
	authService := service.NewAuthService()

	api.Post("signup", func(c *fiber.Ctx) error {
		payload := service.SignupDto{}
		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}
		output, err := authService.Signup(payload)
		if err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(output)
	})
	return nil
}
