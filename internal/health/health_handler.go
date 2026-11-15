package health

import (
	"net/http"
	"github.com/labstack/echo/v5"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(c *echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{Status: "healthy"})
}
