package handlers

import (
	"fmt"
	"net/http"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

// bill godoc
//
//	@Summary		Gets bill from its id
//	@Description	Returns a status, depending on if the bill was retrieved successfully or not
//	@Tags			bill
//	@Param			id	path	int	true	"id of the bill"	format(int)
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Router			/api/v1/bills/{id} [get]
func GetBill(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error parsing id parameter of type integer")
		}

		var o models.Order
		err = s.Get(&o, "bill", "id = ?", id)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Error getting bill of id: %d", id))
		}

		return c.Status(http.StatusOK).JSON(o)
	}
}

// bill godoc
//
//	@Summary		Patches a bill
//	@Description	Returns a status, depending on if the bill was updated successfully or not
//	@Tags			bill
//	@Param			data	body	models.Order	true	"JSON object of an order to be updated"
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Router			/api/v1/bills [patch]
func PatchBill(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var o models.Order
		err := c.BodyParser(&o)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error parsing order")
		}

		cpy := o
		e := s.Contains(&cpy, "bill")
		if !e {
			return c.Status(http.StatusNoContent).SendString("Bill does not exist")
		}

		err = s.Update(&o, "bill")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error updating bill")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// bill godoc
//
//	@Summary		Gets a bill and its orders
//	@Description	Returns a status, depending on if the Bill and Orders were successfully fetched
//	@Tags			bill
//	@Param			id	path	int	true	"id of the bill"	format(int)
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/bill/orders/{id} [get]
func GetBillOrders(s *database.Store) fiber.Handler {
	type Response struct {
		Bill        models.Bill
		Orders      []models.Order `json:"orders"`
		OrdersCount int            `json:"ordersCount"`
	}

	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error parsing bill id, did you use an int?")
		}

		r := Response{}
		err = s.Get(&r.Bill, "bills", "id = ?", id)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error fetching bill from the database")
		}

		err = s.DB.Table("orders").Preload("Item").Where("bill_refer = ?", id).Find(&r.Orders).Error
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("error fetching orders from the database")
		}

		r.OrdersCount = len(r.Orders)

		return c.Status(http.StatusOK).JSON(r)
	}
}

// bill godoc
//
//	@Summary		Patches a bill depending on ID
//	@Description	Returns a status, depending on if the Order was deleted successfully or not
//	@Tags			bill
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/bills/:id [patch]
func PatchBills(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error parsing bill ID, did you use an int?")
		}

		var updatedBill models.Bill
		if err := c.BodyParser(&updatedBill); err != nil {
			fmt.Println(err)
			return c.Status(http.StatusBadRequest).SendString("Error parsing request")
		}

		// Check if the bill exists
		if err := s.Contains(&models.Bill{}, "bills", "id = ?", id); !err {
			fmt.Println(err)
			return c.Status(http.StatusNotFound).SendString("Bill not found")
		}

		// Update the bill in the database
		if err := s.Update(&updatedBill, "bills"); err != nil {
			fmt.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Error updating bill")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
