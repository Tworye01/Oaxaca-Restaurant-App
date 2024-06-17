package middleware

import (
	"net/http"
	"os"
	"team-project/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	v := m.Run()
	os.Exit(v)
}

func TestRequiresRole(t *testing.T) {
	s := database.New()
	app := fiber.New()
	app.Get("/api/v1/role", RequiresRole(1, s))

	req, err := http.NewRequest("GET", "/api/v1/role", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)

	app.Get("/api/v1/role2", RequiresRole(0, s))
	req, err = http.NewRequest("GET", "/api/v1/role2", nil)
	assert.NoError(t, err)
	res, err = app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestBasicMiddleware(t *testing.T) {
	app := fiber.New()
	app.Get("/api/v1/middleware", BasicMiddleware())

	req, err := http.NewRequest("GET", "/api/v1/middleware", nil)
	assert.NoError(t, err)
	res, err := app.Test(req, 1000)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
