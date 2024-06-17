package handlers

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetImage(t *testing.T) {
	app := fiber.New()

	app.Get("/images/:dir/:file", GetImage())

	req, err := http.NewRequest("GET", "/images/:dir/:file", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
