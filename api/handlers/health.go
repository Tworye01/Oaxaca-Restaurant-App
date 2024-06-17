package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// health
//
//	@Summary		Get the health
//	@Description	Returns the status of the middleware
//	@Tags			health
//	@Produce		plain
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/health [get]
func GetHealth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Healthy")
	}
}
