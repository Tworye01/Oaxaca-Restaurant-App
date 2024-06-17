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

// Handler for retrieving the menu items
// dish godoc
//
//	@Summary		Lists all the dishes
//	@Description	Returns the json of all dishes
//	@Tags			dishes
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/dishes/ [get]
func ListDishes(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gets dishes from table
		var dishes []models.Dish
		err := s.DB.Table("dishes").Preload("Item").Find(&dishes).Error
		if err != nil {
			// Handles error when dishes cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving dishes " + err.Error())
		}

		// Sends menu as JSON if no errors have occured
		return c.JSON(dishes)
	}
}

// Handler for retrieving the menu items
// dish godoc
//
//	@Summary		Gets a dish
//	@Description	Returns the dish based on its ID
//	@Tags			dishes
//	@Param			id	path	integer	true	"ID of the dish"
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/dishes/{id} [get]
func GetDish(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pk, err := c.ParamsInt("id")
		if err != nil {
			// Handles error when drinks cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving dish")
		}

		// Gets drinks from table
		var dish models.Dish
		err = s.DB.Table("dishes").Preload("Item").Where("id = ?", pk).Find(&dish).Error
		if err != nil {
			// Handles error when drinks cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving dish")
		}

		// Sends menu as JSON if no errors have occured
		return c.JSON(dish)
	}
}

// handler for adding menu items
// dish godoc
//
//	@Summary		Posts a dish
//	@Description	Posts a dish into the database
//	@Tags			dishes
//	@Param			data	body	models.Dish	true	"JSON object of an order to be added or updated"
//	@Produce		plain
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/dishes/ [post]
func PostDish(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing form")
		}

		bdy, ok := form.Value["body"]
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing form body")
		}

		var d models.Dish
		err = json.Unmarshal([]byte(bdy[0]), &d)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error unmrashalling json data in body field")
		}

		err = s.Add(&d, "dishes")
		if err != nil {
			// Handles error when dish cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error adding dish")
		}

		img, err := c.FormFile("img")
		if err != nil {
			return c.Status(http.StatusNoContent).SendString("Created dish with no image")
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

		return c.Status(http.StatusNoContent).SendString("Created dish with image")
	}
}

// Handler for updating a dish
// dish godoc
//
//	@Summary		Updates a dish
//	@Description	Updates the dish on the ID
//	@Tags			dishes
//	@Param			data	body	models.Dish	true	"JSON object of an order to be added or updated"
//	@Produce		plain
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/dishes [patch]
func PatchDish(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing form")
		}

		bdy, ok := form.Value["body"]
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing form body")
		}

		var d models.Dish
		err = json.Unmarshal([]byte(bdy[0]), &d)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error unmrashalling json data in body field")
		}

		e := s.Contains(&models.Dish{}, "dishes", "id = ?", d.ID)
		if !e {
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating drink")
		}
		e = s.Contains(&models.MenuItem{}, "menu_items", "id = ?", d.Item.ID)
		if !e {
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating drink")
		}

		err = s.Update(&d, "dishes")
		if err != nil {
			// Handles error when dish cannot be updated
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating dish")
		}

		err = s.Update(&d.Item, "menu_items")
		if err != nil {
			// Handles error when dish cannot be updated
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating dish")
		}

		img, err := c.FormFile("img")
		if err != nil {
			return c.Status(http.StatusNoContent).SendString("Created dish with no image")
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

		return c.Status(http.StatusNoContent).SendString("Created dish with image")
	}
}

// dishes godoc
//
//	@Summary		Put a dish
//	@Description	Puts an dish, if it does not exist it is added otherwise it updates the order of the same ID
//	@Tags			dishes
//	@Param			data	body	models.Dish	true	"JSON object of an order to be added or updated"
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/dishes [put]
func PutDish(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var d models.Dish

		err := c.BodyParser(&d)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error parsing body for Dish model")
		}

		db := s.Contains(&d, "dishes")
		if db {
			err := s.Update(&d, "dishes")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error updating dish")
			}
		} else {
			err := s.Add(&d, "dishes")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error adding dish")
			}
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// Handler for deleting a dish
// dish godoc
//
//	@Summary		Deletes a dish
//	@Description	Deletes the dish on the ID
//	@Tags			dishes
//	@Param			id	path	integer	true	"ID of the dish"
//	@Produce		plain
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/menu/dishes/{id} [delete]
func DeleteDish(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			// Handles error when dish cannnot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving dish")
		}

		d := models.Dish{}
		err = s.Get(&d, "dishes", "id = ?", id)
		if err != nil {
			// Handles error when drinks cannot be deleted
			return c.Status(fiber.StatusInternalServerError).SendString("Error removing drink")
		}

		err = s.DB.Table("dishes").Where("id = ?", id).Delete(&models.Dish{}).Error
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
