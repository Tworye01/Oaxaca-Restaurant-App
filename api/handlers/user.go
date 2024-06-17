package handlers

import (
	"net/http"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

/* Waiters can:
- change the menu
- change status of order
- confirm order
- view order times
- cancel order
- see that client needs help

Kitchen can:
- notify waiters that order is ready
- view order confirmation
- view order times

Customer can:
- view menu
- order
- call waiter

Admin can:
- delete users
- update users
*/

// Handler for retrieving a user
// dish godoc
//
//	@Summary		Gets a user
//	@Description	Returns the user based on their username
//	@Tags			user
//	@Param			username	path	integer	true	"username"
//	@Produce		json
//	@Success		204
//	@Failure		500
//	@Router			/api/v1/login/credentials/{username} [get]
func GetUser(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		usr := c.Params("username")

		var user models.Credentials
		// Gets user from table
		err := s.Get(&user, "credentials", "username = ?", usr)
		if err != nil {
			// Handles error when user cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving user")
		}

		// Sends user as JSON if no errors have occured
		return c.JSON(user)
	}
}

// Handler for updating a user
// user godoc
//
//	@Summary		Updates a user
//	@Description	Updates user based on their credentials
//	@Tags			user
//	@Param			data	body	models.Credentials	true	"user details"
//	@Produce		plain
//	@Success		204
//	@Failure		500
//	@Router			/api/v1/login/credentials [patch]
func UpdateUser(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var d models.Credentials
		err := c.BodyParser(&d)
		if err != nil {
			// Handles error when user cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving user")
		}

		err = s.Update(&d, "credentials")
		if err != nil {
			// Handles error when user cannot be updated
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating user")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// Handler for banishing a user
// user godoc
//
//	@Summary		Deletes a user
//	@Description	Deletes the user via their username
//	@Tags			user
//	@Param			user	path	integer	true	"username"
//	@Produce		plain
//	@Success		204
//	@Failure		500
//	@Router			/api/v1/login/credentials/{user} [delete]
func DeleteUser(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("user")

		err := s.Delete(&models.Credentials{}, "credentials", "username = ?", name)
		if err != nil {
			// Handles error when user cannot be deleted
			return c.Status(fiber.StatusInternalServerError).SendString("Error removing user")
		}

		return c.SendStatus(http.StatusNoContent)
	}

}

// Handler to add a user
// user godoc
//
//	@Summary		Posts a user
//	@Description	Posts a user into the database
//	@Tags			user
//	@Param			data	body	models.Credentials	true	"user details"
//	@Produce		plain
//	@Success		204
//	@Failure		500
//	@Router			/api/v1/login/credentials/ [post]
func PostUser(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var d models.Credentials
		err := c.BodyParser(&d)
		if err != nil {
			// Handles error when user cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving user")
		}
		// Adds user info and initially sets role to customer
		err = s.Add(&d, "credentials")
		d.Role = 0
		if err != nil {
			// Handles error when user cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error adding user")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
