package handlers

import (
	"net/http"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetTableOrders(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/table/orders/:table", GetTableOrders(s))

	req, err := http.NewRequest("GET", "/api/v1/table/orders/0", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetTableBills(t *testing.T) {
	s := database.New()
	app := fiber.New()

	app.Get("/api/v1/table/bills/:table", GetTableBills(s))

	req, err := http.NewRequest("GET", "/api/v1/table/bills/0", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
