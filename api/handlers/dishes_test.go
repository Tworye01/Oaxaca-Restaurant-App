package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestListDishes(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/menu/dishes", ListDishes(s))

	req, err := http.NewRequest("GET", "/api/v1/menu/dishes", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestGetDishes(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/menu/dishes/:id", GetDish(s))

	req, err := http.NewRequest("GET", "/api/v1/menu/dishes/1", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestPostDishes(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Post("/api/v1/menu/dishes", PostDish(s))

	req, err := http.NewRequest("POST", "/api/v1/menu/dishes", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestPatchDishes(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Patch("/api/v1/menu/dishes", PatchDish(s))

	req, err := http.NewRequest("PATCH", "/api/v1/menu/dishes", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestDeleteDishes(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Delete("/api/v1/menu/dishes", DeleteDish(s))

	req, err := http.NewRequest("DELETE", "/api/v1/menu/dishes", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}

func TestPutDishes(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Put("/api/v1/menu/dishes", PutDish(s))

	req, err := http.NewRequest("PUT", "/api/v1/menu/dishes", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

}
