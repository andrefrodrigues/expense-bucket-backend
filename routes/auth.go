package routes

import (
	"expense-bucket-api/middleware"
	"expense-bucket-api/model"
	"expense-bucket-api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupAuthRoutes(config ApiConfig) error {
	group := config.group
	authService := config.AuthService

	signupHandler := func(c echo.Context) error {
		payload := service.SignupDto{}
		if err := (&echo.DefaultBinder{}).Bind(&payload, c); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		output, err := authService.Signup(c.Request().Context(), payload)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, output)
	}

	loginHandler := func(c echo.Context) error {
		payload := service.LoginDto{}
		if err := (&echo.DefaultBinder{}).Bind(&payload, c); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		output, err := authService.Login(c.Request().Context(), payload)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, output)
	}

	meHandler := func(c echo.Context) error {
		data := c.Get(middleware.USER_KEY)
		user, ok := data.(*model.User)
		if !ok {
			return c.String(http.StatusUnauthorized, "No user")
		}

		return c.JSON(http.StatusOK, user.ToDTO())
	}
	group.POST("/signup", signupHandler)
	group.POST("/login", loginHandler)
	group.GET("/me", meHandler, config.AuthMiddleware)
	return nil
}
