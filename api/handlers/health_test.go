package handlers

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	app := fiber.New()

	app.Get("/api/v1/health", GetHealth())

	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

}
