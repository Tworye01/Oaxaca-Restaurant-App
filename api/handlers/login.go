package handlers

import (
	"net/http"
	"team-project/auth"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

// Handler for authenticating users
// login godoc
//
//	@Summary		Authenticates user
//	@Description	Logs user in if valid login details are given and returns a login token.
//	@Tags			login
//	@Param			data	body	models.Credentials	true	"user's login details"
//	@Produce		plain
//	@Success		200
//	@Failure		400
//	@Failure		561
//	@Router			/api/v1/login [post]
func LoginHandler(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginDetailsGiven models.Credentials
		err := c.BodyParser(&loginDetailsGiven)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid request")
		}

		var user models.Credentials
		err = s.Get(&user, "credentials", "username = ?", loginDetailsGiven.Username)

		// Unauthorized status response if user is not found
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString("Invalid login details")
		}
		if user.Password != loginDetailsGiven.Password {
			return c.Status(http.StatusUnauthorized).SendString("Invalid login details")
		}
		tok, err := auth.CreateToken(loginDetailsGiven.Username, loginDetailsGiven.Password)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString("Error creating token")

		}
		// Send token if successful
		return c.Status(http.StatusOK).SendString(tok)
	}
}

// Handler that can only used when the bearer has the correct authorization
func ProtectedHandler(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		h := c.GetReqHeaders()
		toks, ok := h["Authorization"]
		if !ok && len(toks) <= 0 {
			return c.Status(http.StatusUnauthorized).SendString("Invalid Authorization Header")
		}
		tok := toks[0]
		if len(tok) < 8 {
			return c.Status(http.StatusUnauthorized).SendString("Invalid Login Token")
		}

		tok = tok[len("Bearer "):]
		err := auth.VerifyToken(tok)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString("Invalid Login Token")
		}
		return c.SendString(tok)
	}

}
