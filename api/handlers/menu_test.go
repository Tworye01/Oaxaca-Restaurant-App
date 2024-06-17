package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetMenu(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/menu", GetMenu(s))

	req, err := http.NewRequest("GET", "/api/v1/menu", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}
