package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Post("api/v1/login", LoginHandler(s))

	req, err := http.NewRequest("POST", "api/v1/login", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, 404, res.StatusCode)
}
