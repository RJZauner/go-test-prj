package health

import (
	"github.com/labstack/echo/v5"
)

func HealthGroupRoutes(e *echo.Echo) {
	e.GET("/health", HealthHandler)
}
