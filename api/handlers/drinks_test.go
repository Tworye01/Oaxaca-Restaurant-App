package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestListDrinks(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/menu/drinks", ListDrinks(s))

	req, err := http.NewRequest("GET", "/api/v1/menu/drinks", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestGetDrinks(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/menu/drinks/:id", GetDrink(s))

	req, err := http.NewRequest("GET", "/api/v1/menu/drinks/1", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestPostDrinks(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Post("/api/v1/menu/drinks", PostDrink(s))

	req, err := http.NewRequest("POST", "/api/v1/menu/drinks", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestPatchDrinks(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Patch("/api/v1/menu/drinks", PatchDrink(s))

	req, err := http.NewRequest("PATCH", "/api/v1/menu/drinks", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestDeleteDrinks(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Delete("/api/v1/menu/drinks", DeleteDrink(s))

	req, err := http.NewRequest("DELETE", "/api/v1/menu/drinks", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestPutDrinks(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Put("/api/v1/menu/drinks", PutDrink(s))

	req, err := http.NewRequest("PUT", "/api/v1/menu/drinks", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}
