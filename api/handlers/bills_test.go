package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBill(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/order/bill/:id", GetBill(s))

	req, err := http.NewRequest("GET", "/api/v1/order/bill/3", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestPatchBill(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Patch("/api/v1/bill", PatchBill(s))

	req, err := http.NewRequest("PATCH", "/api/v1/bill", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetBillOrders(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/bill/orders/:id", GetBillOrders(s))

	req, err := http.NewRequest("GET", "/api/v1/bill/orders/0", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
