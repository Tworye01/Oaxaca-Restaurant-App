package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestListOrder(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Get("/api/v1/orders", ListOrder(s))

	req, err := http.NewRequest("GET", "/api/v1/orders", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetOrder(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Get("/api/v1/orders/:id", GetOrder(s))

	req, err := http.NewRequest("GET", "/api/v1/orders/1", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetOrderStatus(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Get("/api/v1/orders/:id/status", GetOrderStatus(s))

	req, err := http.NewRequest("GET", "/api/v1/orders/1/status", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestPostOrder(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Post("/api/v1/orders", PostOrder(s))

	req, err := http.NewRequest("POST", "/api/v1/orders", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestPatchOrder(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Patch("/api/v1/orders", PatchOrder(s))

	req, err := http.NewRequest("PATCH", "/api/v1/orders", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestPatchOrderStatus(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Patch("/api/v1/orders/:id/:status", PatchOrderStatus(s))

	req, err := http.NewRequest("PATCH", "/api/v1/orders/1/2", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)

	req, err = http.NewRequest("PATCH", "/api/v1/orders/10/10", nil)
	assert.NoError(t, err)
	res, err = app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestPutOrder(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Put("/api/v1/orders", PutOrder(s))

	req, err := http.NewRequest("PUT", "/api/v1/orders", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestDeleteOrder(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Delete("/api/v1/orders/:id", DeleteOrder(s))

	req, err := http.NewRequest("DELETE", "/api/v1/orders/1", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}
