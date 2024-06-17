package handlers

import (
	"strings"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

// fliter menu
//
//	@Summary		Get the static images for the file directory
//	@Description	Goes to the image/menu/... directory and sends the file.
//	@Tags			filter
//	@Param			query	path	int		true	"spice"
//	@Param			query	path	int		true	"course"
//	@Param			query	path	string	true	"allergens"
//	@Param			query	path	int		true	"calories"
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/dishes/query [get]
func GetFilter(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {

		m := c.Queries()
		// Construct the query
		spice := c.QueryInt("spice", 5)
		course := c.QueryInt("course", 5)
		calories := c.QueryInt("calories", 1000)
		vegetarian := c.QueryInt("vegetarian", 3)
		// Construct the interface
		var dishes []models.Dish

		query := s.DB.Table("dishes").Joins("Item")

		// Checks if spice has changed
		if spice != 5 {
			query = query.Where("spice = ?", spice)
		}

		// Checks against Max calories
		if calories != 1000 {
			query = query.Where("calories <= ?", calories)
		}

		// Check for allergens
		allergenString := m["allergens"]
		allergens := strings.Split(allergenString, ",")

		for _, allergen := range allergens {
			allergen = strings.TrimSpace(allergen)
			if allergen != "" {
				query = query.Where("allergens NOT LIKE ?", "%"+allergen+"%")
			}
		}

		// Checks if course has changed
		if course != 5 {
			query = query.Where("course = ?", course)
		}

		// Checks if vegetarian
		if vegetarian == 0 {
			query = query.Where("vegetarian = ?", false)
		}

		if vegetarian == 1 {
			query = query.Where("vegetarian = ?", true)
		}

		err := query.Find(&dishes).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving dishes")
		}

		return c.JSON(dishes)
	}
}
