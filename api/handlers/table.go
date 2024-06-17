package handlers

import (
	"net/http"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

// table godoc
//
//	@Summary		Gets orders for a specific table
//	@Description	Fetches the orders for specific table number and returns them.
//	@Tags			table
//	@Param			table	path	int	true	"table number"	format(int)
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Router			/api/v1/table/orders/{table} [get]
func GetTableOrders(s *database.Store) fiber.Handler {
	type Response struct {
		Orders      []models.Order `json:"orders"`
		OrdersCount int            `json:"ordersCount"`
	}

	return func(c *fiber.Ctx) error {
		tbl, err := c.ParamsInt("table")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("bad table parameter given, please give an integer")
		}

		r := Response{}
		err = s.DB.Table("orders").Preload("Bill").Preload("Item").Where("table_refer = ?", tbl).Find(&r.Orders).Error
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error fetching data from the database, maybe your table doesn't exist?")
		}

		r.OrdersCount = len(r.Orders)

		return c.Status(http.StatusOK).JSON(r)
	}
}

// order godoc
//
//	@Summary		Gets bill for a specific table
//	@Description	Fetches the bill for specific table number and returns it.
//	@Tags			table
//	@Param			table	path	int	true	"table number"	format(int)
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Router			/api/v1/table/bills/{table} [get]
func GetTableBills(s *database.Store) fiber.Handler {
	type Response struct {
		Bills      []models.Bill `json:"bills"`
		BillsCount int           `json:"billsCount"`
	}

	return func(c *fiber.Ctx) error {
		tbl, err := c.ParamsInt("table")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("bad table parameter given, please give an integer")
		}

		r := Response{}
		err = s.DB.Table("bills").Where("table_refer = ?", tbl).Find(&r.Bills).Error
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error fetching data from the database, maybe your table doesn't exist?")
		}

		r.BillsCount = len(r.Bills)

		return c.Status(http.StatusOK).JSON(r)
	}
}
