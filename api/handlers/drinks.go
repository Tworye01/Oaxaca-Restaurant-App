package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"team-project/database"
	"team-project/models"
	"team-project/utils"

	"github.com/gofiber/fiber/v2"
)

// handler for retrieving the menu items
// drink godoc
//
//	@Summary		Lists all the drinks
//	@Description	Returns the json of drinks
//	@Tags			drinks
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/drinks [get]
func ListDrinks(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gets dishes from table
		var drinks []models.Drink
		err := s.DB.Table("drinks").Preload("Item").Find(&drinks).Error
		if err != nil {
			// Handles error when dishes cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving drinks")
		}

		// Sends menu as JSON if no errors have occured
		return c.JSON(drinks)
	}
}

// Handler for retrieving the menu items
// drink godoc
//
//	@Summary		Gets a drink
//	@Description	Gets a drink, using the primary key
//	@Tags			drinks
//	@Param			id	path	integer	true	"ID of the dish"
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/drinks.{id} [put]
func GetDrink(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pk, err := c.ParamsInt("id")
		if err != nil {
			// Handles error when drinks cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving drink")
		}

		// Gets drinks from table
		var drink models.Drink
		err = s.DB.Table("drinks").Preload("Item").Where("id = ?", pk).Find(&drink).Error
		if err != nil {
			// Handles error when drinks cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving drink")
		}

		// Sends menu as JSON if no errors have occured
		return c.JSON(drink)
	}
}

// handler for adding menu items
// drink godoc
//
//	@Summary		Posts a drink
//	@Description	Added a drink to the database regardless if it alread exists or not
//	@Tags			drinks
//	@Param			data	body	models.Drink	true	"JSON object of an order to be added or updated"
//	@Produce		plain
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/drinks [post]
func PostDrink(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing form")
		}

		bdy, ok := form.Value["body"]
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing form body")
		}

		var d models.Drink
		err = json.Unmarshal([]byte(bdy[0]), &d)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error unmrashalling json data in body field")
		}

		err = s.Add(&d, "drinks")
		if err != nil {
			// Handles error when drinks cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error adding drink")
		}

		img, err := c.FormFile("img")
		if err != nil {
			return c.Status(http.StatusNoContent).SendString("Created drink with no image")
		}

		ext, err := utils.GetFileExtension(img.Filename)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Bad file was given")
		}

		img.Filename = fmt.Sprintf("%d.%s", d.Item.ID, ext)
		err = c.SaveFile(img, fmt.Sprintf("images/menu/items/%s", img.Filename))
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error saving the given file")
		}

		d.Item.Image = fmt.Sprintf("http://localhost:8080/%s", fmt.Sprintf("images/items/%s", img.Filename))
		err = s.Update(&d.Item, "menu_items")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error saving the given file")
		}

		return c.Status(http.StatusNoContent).SendString("Created drink with image")
	}

}

// Handler for updating a drink
// drink godoc
//
//	@Summary		Updates a drink
//	@Description	Updates the drink at the certain ID
//	@Tags			drinks
//	@Param			data	body	models.Drink	true	"JSON object of an order to be added or updated"
//	@Produce		plain
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/drinks/{id} [patch]
func PatchDrink(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing form")
		}

		bdy, ok := form.Value["body"]
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing form body")
		}

		var d models.Drink
		err = json.Unmarshal([]byte(bdy[0]), &d)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error unmrashalling json data in body field")
		}

		e := s.Contains(&models.Drink{}, "drinks", "id = ?", d.ID)
		if !e {
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating drink")
		}
		e = s.Contains(&models.MenuItem{}, "menu_items", "id = ?", d.Item.ID)
		if !e {
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating drink")
		}

		err = s.Update(&d, "drinks")
		if err != nil {
			// Handles error when drink cannot be updated
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating drink")
		}

		err = s.Update(&d.Item, "menu_items")
		if err != nil {
			// Handles error when drink cannot be updated
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating drink")
		}

		img, err := c.FormFile("img")
		if err != nil {
			return c.Status(http.StatusNoContent).SendString("Created drink with no image")
		}

		ext, err := utils.GetFileExtension(img.Filename)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Bad file was given")
		}

		img.Filename = fmt.Sprintf("%d.%s", d.Item.ID, ext)
		err = c.SaveFile(img, fmt.Sprintf("images/menu/items/%s", img.Filename))
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error saving the given file")
		}

		d.Item.Image = fmt.Sprintf("http://localhost:8080/%s", fmt.Sprintf("images/items/%s", img.Filename))
		err = s.Update(&d.Item, "menu_items")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error saving the given file")
		}

		return c.Status(http.StatusNoContent).SendString("Created drink with image")
	}
}

// drink godoc
//
//	@Summary		Put a drink
//	@Description	Puts an drink, if it does not exist it is added otherwise it updates the order of the same ID
//	@Tags			drinks
//	@Param			data	body	models.Drink	true	"JSON object of an order to be added or updated"
//	@Produce		plain
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/drinks [put]
func PutDrink(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var d models.Drink

		err := c.BodyParser(&d)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error parsing body for Drink model")
		}

		db := s.Contains(&d, "drinks")
		if db {
			err := s.Update(&d, "drinks")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error updating drink")
			}
		} else {
			err := s.Add(&d, "drinks")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error adding drink")
			}
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// Handler for deleting a drink
// drink godoc
//
//	@Summary		Deletes a drink
//	@Description	Deletes a drink using it's ID
//	@Tags			drinks
//	@Param			id	path	integer	true	"ID of the dish"
//	@Produce		plain
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/drinks/{id} [delete]
func DeleteDrink(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			// Handles error when drinks cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error removing drink")
		}

		d := models.Drink{}
		err = s.Get(&d, "drinks", "id = ?", id)
		if err != nil {
			// Handles error when drinks cannot be deleted
			return c.Status(fiber.StatusInternalServerError).SendString("Error removing drink")
		}

		err = s.DB.Table("drinks").Where("id = ?", id).Delete(&models.Drink{}).Error
		if err != nil {
			// Handles error when drinks cannot be deleted
			return c.Status(fiber.StatusInternalServerError).SendString("Error removing drink")
		}

		err = s.DB.Table("menu_items").Where("id = ?", d.ItemRefer).Delete(&models.MenuItem{}).Error
		if err != nil {
			// Handles error when drinks cannot be deleted
			return c.Status(fiber.StatusInternalServerError).SendString("Error removing drink")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
