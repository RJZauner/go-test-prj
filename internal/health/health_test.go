package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	e := echo.New()
	HealthGroupRoutes(e)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	// require http status of 200
	require.Equal(t, http.StatusOK, rec.Code)

	// assert response body is correct
	var got HealthResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &got))
	assert.Equal(t, HealthResponse{Status: "health"}, got)
}

