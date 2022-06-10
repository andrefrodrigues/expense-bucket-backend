package server

import (
	"expense-bucket-api/model"
	"expense-bucket-api/routes"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() error {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "from=${remote_ip} method=${method}, uri=${uri}, status=${status}\n",
	}))

	dsn := os.Getenv("CONNECTION")
	db, err := model.SetupDatabase(dsn)

	if err != nil {
		return err
	}

	config := routes.RoutesConfig{
		Echo: e,
		DB:   db,
	}
	err = routes.Setup(config)
	if err != nil {
		return err
	}
	return e.Start(":3000")
}
