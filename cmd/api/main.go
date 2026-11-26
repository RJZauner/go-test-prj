package main

import (
	"github.com/labstack/echo/v5"
  "github.com/labstack/echo/v5/middleware"
	"github.com/RJZauner/go-test-prj/internal/health"
)

func main() {
	e := echo.New()

	// Comment added to test pipeline - take 2
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	health.HealthGroupRoutes(e)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
