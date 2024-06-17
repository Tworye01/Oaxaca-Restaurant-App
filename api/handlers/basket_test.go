package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPostBasketItems(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Post("/api/v1/baskets/items", PostBasketItems(s))

	req, err := http.NewRequest("POST", "/api/v1/baskets/items", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, 400, res.StatusCode)
}
