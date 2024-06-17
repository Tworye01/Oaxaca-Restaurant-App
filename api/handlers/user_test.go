package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/login/credentials", GetUser(s))

	req, err := http.NewRequest("GET", "api/v1/login/credentials", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, 404, res.StatusCode)
}

func TestPostUser(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Post("/api/v1/login/credentials", PostUser(s))

	req, err := http.NewRequest("POST", "/api/v1/login/credentials", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestUpdateUser(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Post("/api/v1/login/credentials", UpdateUser(s))

	req, err := http.NewRequest("PATCH", "/api/v1/login/credentials", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, 405, res.StatusCode)
}

func TestDeleteUser(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Post("/api/v1/login/credentials", DeleteUser(s))

	req, err := http.NewRequest("DELETE", "/api/v1/login/credentials", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, 405, res.StatusCode)

}
