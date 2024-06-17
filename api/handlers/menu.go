package handlers

import (
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

// menu godoc
//
//	@Summary		Get the entire Menu
//	@Description	Returns all the Dishes and Drinks, regardless of availability
//	@Tags			menu
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/api/v1/menu [get]
func GetMenu(s *database.Store) fiber.Handler {
	type Response struct {
		Dishes []models.Dish
		Drinks []models.Drink
	}

	return func(c *fiber.Ctx) error {
		var menu Response

		// Gets dishes from table
		err := s.List(&menu.Dishes, "dishes")
		if err != nil {
			// Handles error when dishes cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving dishes")
		}

		// Gets drinks from table
		err = s.List(&menu.Drinks, "drinks")
		if err != nil {
			// Handles error when drinks cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving drinks")
		}

		// Sends menu as JSON if no errors have occured
		return c.JSON(menu)
	}
}
