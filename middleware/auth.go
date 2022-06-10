package middleware

import (
	"expense-bucket-api/auth"
	"expense-bucket-api/service"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type MiddlewareConfig struct {
	UserService *service.UserService
}

const USER_KEY = "user"

func BuildAuthMiddleware(config MiddlewareConfig) echo.MiddlewareFunc {
	userService := config.UserService
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
			if !auth.IsTokenValid(token) {
				return c.String(http.StatusUnauthorized, "Invalid token")
			}
			tokenData, err := auth.GetTokenData(token)
			if err != nil {
				return c.String(http.StatusUnauthorized, err.Error())
			}
			user, err := userService.GetUserByEmail(c.Request().Context(), tokenData.Email)
			if err != nil {
				return c.String(http.StatusUnauthorized, err.Error())
			}
			c.Set(USER_KEY, user)
			next(c)
			return nil
		}
	}
}
