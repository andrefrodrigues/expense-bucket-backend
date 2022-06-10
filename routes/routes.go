package routes

import (
	"expense-bucket-api/middleware"
	"expense-bucket-api/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RoutesConfig struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

type ApiConfig struct {
	group          *echo.Group
	AuthService    *service.AuthService
	UserService    *service.UserService
	AuthMiddleware echo.MiddlewareFunc
}

func Setup(config RoutesConfig) error {
	e := config.Echo
	group := e.Group("/api")
	authService := service.NewAuthService(service.AuthServiceConfig{
		DB: config.DB,
	})
	userService := service.NewUserService(service.UserServiceConfig{
		DB: config.DB,
	})
	authMiddleware := middleware.BuildAuthMiddleware(middleware.MiddlewareConfig{
		UserService: userService,
	})
	apiConfig := ApiConfig{
		group:          group,
		AuthService:    authService,
		UserService:    userService,
		AuthMiddleware: authMiddleware,
	}
	err := SetupAuthRoutes(apiConfig)
	if err != nil {
		return err
	}

	return nil
}
