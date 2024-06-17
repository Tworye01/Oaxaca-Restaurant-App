package handlers

import (
	"net/http"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

// Function for calculating total cost of items in the basket
func CalculateBasketCost(i []models.MenuItem, ids []int) float32 {
	var c float32 = 0.0

	for _, id := range ids {
		for _, item := range i {
			if item.ID == uint(id) {
				c += item.Cost
				break
			}
		}
	}

	return c
}

// Returns the cost of a given ids item
func calculateItemCost(i []models.MenuItem, id int) float32 {
	for _, item := range i {
		if item.ID == uint(id) {
			return item.Cost
		}
	}

	return 0
}

// Handler for adding items to basket
// basket godoc
//
//	@Summary		Adds items to basket
//	@Description	Gets data from the request body and item information from the database, calculates the total cost, and returns the basket items
//	@Tags			basket
//	@Param			data		body	models.Basket	true	"JSON object representing the basket"
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Router			/api/v1/baskets/items/ [post]
func PostBasketItems(s *database.Store) fiber.Handler {
	type Response struct {
		Items      []models.MenuItem `json:"items"`
		ItemsCount int               `json:"itemsCount"`
		Cost       float32           `json:"cost"`
	}

	return func(c *fiber.Ctx) error {
		var b models.Basket
		err := c.BodyParser(&b)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error parsing basket body")
		}

		if len(b.Items) <= 0 {
			r := Response{
				Items:      []models.MenuItem{},
				ItemsCount: 0,
				Cost:       0.0,
			}
			return c.Status(http.StatusOK).JSON(r)
		}

		ids := make([]int, 0)

		for _, item := range b.Items {
			ids = append(ids, item.ItemID)
		}

		r := Response{}
		err = s.DB.Table("menu_items").Find(&r.Items, ids).Error
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error fetching item ids from database")
		}

		r.ItemsCount = len(r.Items)
		r.Cost = CalculateBasketCost(r.Items, ids)

		return c.Status(http.StatusOK).JSON(r)
	}
}

// Handler for payment
// basket godoc
//
//	@Summary		Posts table's bill based on basket
//	@Description	Gets item information from basket and table number, calculates total cost, adds orders & bill to database
//	@Tags			basket
//	@Param			data		body	models.Basket	true	"JSON object representing the basket"
//	@Param			table	path	int	true	"table number"
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Router			/api/v1/baskets/pay/{table} [post]
func PostBasketPayment(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var b models.Basket
		err := c.BodyParser(&b)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error parsing basket body")
		}

		tbl, err := c.ParamsInt("table")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error parsing table number")
		}

		ids := make([]int, 0)
		items := make([]models.MenuItem, 0)

		for _, item := range b.Items {
			ids = append(ids, item.ItemID)
		}

		err = s.DB.Table("menu_items").Find(&items, ids).Error
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error fetching item ids from database")
		}

		bill := models.Bill{
			TableRefer: tbl,
			Cost:       CalculateBasketCost(items, ids),
			Paid:       true,
		}

		err = s.Add(&bill, "bills")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error saving bill to database")
		}

		for _, item := range b.Items {
			o := models.Order{
				TableRefer: tbl,
				Status:     models.Preparation,
				Cost:       calculateItemCost(items, item.ID),
				Note:       item.Note,

				ItemRefer: uint(item.ItemID),
				BillRefer: bill.ID,
			}

			err = s.Add(&o, "orders")
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString("error saving bill to database")
			}
		}

		return c.SendStatus(http.StatusOK)
	}
}

// Handler for basket
// basket godoc
//
//	@Summary		Posts table's bill based on basket
//	@Description	Gets item information from basket and table number, calculates total cost, adds orders & bill to database
//	@Tags			basket
//	@Param			data		body	models.Basket	true	"JSON object representing the basket"
//	@Param			table	path	int	true	"table number"
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Router			/api/v1/baskets/{table} [post]
func PostBasket(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var b models.Basket
		err := c.BodyParser(&b)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error parsing basket body")
		}

		tbl, err := c.ParamsInt("table")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error parsing table number")
		}

		ids := make([]int, 0)
		items := make([]models.MenuItem, 0)

		for _, item := range b.Items {
			ids = append(ids, item.ItemID)
		}

		err = s.DB.Table("menu_items").Find(&items, ids).Error
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error fetching item ids from database")
		}

		bill := models.Bill{
			TableRefer: tbl,
			Cost:       CalculateBasketCost(items, ids),
			Paid:       false,
		}

		err = s.Add(&bill, "bills")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error saving bill to database")
		}

		for _, item := range b.Items {
			o := models.Order{
				TableRefer: tbl,
				Status:     models.Preparation,
				Cost:       calculateItemCost(items, item.ID),
				Note:       item.Note,

				ItemRefer: uint(item.ItemID),
				BillRefer: bill.ID,
			}

			err = s.Add(&o, "orders")
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString("error saving bill to database")
			}
		}

		return c.SendStatus(http.StatusOK)
	}
}
