package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// image godoc
//
//	@Summary		Get the static images for the file directory
//	@Description	Goes to the image/menu/... directory and sends the file.
//	@Tags			image
//	@Param			dir		path	string	true	"dishes or drinks"	format(string)
//	@Param			file	path	string	true	"name of the file"	format(string)
//	@Produce		plain
//	@Success		200
//	@Failure		500
//	@Router			/images/{dir}/{file} [get]
func GetImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		d := c.Params("dir")
		f := c.Params("file")
		p := "images/menu/" + d + "/" + f

		return c.SendFile(p)
	}
}

// image godoc
//
//	@Summary		Posts the image into the file directory
//	@Description	Takes the file and saves it into the path
//	@Tags			image
//	@Param			dir		path	string	true	"dishes or drinks"	format(string)
//	@Param			file	path	string	true	"name of the file"	format(string)
//	@Produce		plain
//	@Success		200
//	@Failure		500
//	@Router			/images/{dir}/{file} [post]
func PostImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		d := c.Params("dir")
		f := c.Params("file")
		file, err := c.FormFile("file")
		if err != nil {
			c.SendStatus(http.StatusBadRequest)
		}
		p := "images/menu/" + d + "/" + f

		c.SaveFile(file, p)

		return c.SendFile(p)
	}
}
