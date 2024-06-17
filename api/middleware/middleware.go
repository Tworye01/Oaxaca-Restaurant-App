package middleware

import (
	"fmt"
	"strings"
	"team-project/auth"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

// Middleware checking if user's role can access a route
func RequiresRole(role models.Role, s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if role == 0 {
			return c.Next()
		}

		// Gets authorisation token
		var token string
		header := c.GetReqHeaders()
		authorization, ok := header["Authorization"]
		if !ok {
			return c.Status(fiber.StatusBadRequest).SendString("Error retrieving token")
		}
		tok := authorization[0]

		if strings.HasPrefix(tok, "Bearer ") {
			token = strings.TrimPrefix(tok, "Bearer ")
		} else if c.Cookies("token") != "" {
			token = c.Cookies("token")
		}

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Error retrieving token")
		}

		// Gets username via token
		username, password, err := auth.GetUsernameFromToken(token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving username")
		}

		// Gets role via username
		var user models.Credentials
		err = s.Get(&user, "credentials", "username = ?", username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving user role")
		}

		// Proceed if users role is high enough for required role
		if user.Role >= role && user.Password == password {
			return c.Next()
		}

		// Denies access if role is not sufficient
		return c.Status(fiber.StatusForbidden).SendString("Access denied")
	}
}

// Basic middleware, prints out middleware on use.
func BasicMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("middleware")

		return c.Next()
	}
}
