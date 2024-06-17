package handlers

import (
	"fmt"
	"net/http"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/fiber/v2"
)

// order godoc
//
//	@Summary		Get a list of all orders
//	@Description	Returns a JSON object of all orders
//	@Tags			order
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Router			/api/v1/orders [get]
func ListOrder(s *database.Store) fiber.Handler {
	type Response struct {
		Orders      []models.Order `json:"orders"`
		OrdersCount int            `json:"ordersCount"`
	}

	return func(c *fiber.Ctx) error {
		var o []models.Order
		err := s.DB.Table("orders").Preload("Item").Find(&o).Error
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error fetching orders from database")
		}

		r := Response{
			Orders:      o,
			OrdersCount: len(o),
		}

		return c.Status(http.StatusOK).JSON(r)
	}
}

// order godoc
//
//	@Summary		Get a specific order
//	@Description	Returns a JSON object of the order of the given ID
//	@Tags			order
//	@Param			id	path	int	true	"id of the order"	format(int)
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Router			/api/v1/orders/{id} [get]
func GetOrder(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error parsing id parameter of type integer")
		}

		var o models.Order
		err = s.DB.Table("orders").Preload("Item").Where("id = ?", id).Find(&o).Error
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Error getting order of id: %d", id))
		}

		return c.Status(http.StatusOK).JSON(o)
	}
}

// order godoc
//
//	@Summary		Get the status of a specific order
//	@Description	Returns the status of the given order ID
//	@Tags			order
//	@Param			id	path	int	true	"id of the order"	format(int)
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/api/v1/orders/{id}/status [get]
func GetOrderStatus(s *database.Store) fiber.Handler {
	type Response struct {
		OrderRefer uint          `json:"orderRefer"`
		Status     models.Status `json:"status"`
	}

	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error parsing id parameter of type integer")
		}

		var o models.Order
		err = s.Get(&o, "orders", "id = ?", id)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Error getting order of id: %d", id))
		}

		r := Response{
			OrderRefer: o.ID,
			Status:     o.Status,
		}

		return c.Status(http.StatusOK).JSON(r)
	}
}

// order godoc
//
//	@Summary		Post an order
//	@Description	Adds an order if it does not already exist in the orders
//	@Tags			order
//	@Param			data	body	models.Order	true	"JSON object of an order to be added"
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/orders [post]
func PostOrder(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var order models.Order
		err := c.BodyParser(&order)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error parsing body for Order model")
		}

		e := s.Contains(&models.Order{Model: models.Model{ID: uint(order.ID)}}, "orders")
		if e && order.ID != 0 {
			return c.Status(http.StatusNoContent).SendString("Order already exists within the database")
		}

		// Check if tabke has a bill
		var bill models.Bill
		err = s.Get(&bill, "bills", "table = ? AND paid = ?", order.TableRefer, false)
		if err != nil {
			// Create a new bill if one does not exist
			bill = models.Bill{
				TableRefer: order.TableRefer,
				Cost:       order.Cost,
				Paid:       false,
			}
			err = s.Add(&bill, "bills")
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString("Error adding new bill")
			}

		} else {
			// Update bill cost if order exists
			bill.Cost += order.Cost
			err = s.Update(&bill, "bills")
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString("Error updating bill")
			}
		}

		// Adds bill reference to order
		order.BillRefer = bill.ID

		err = s.Add(&order, "orders")
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// order godoc
//
//	@Summary		Patch an order
//	@Description	Updates the data of a currently existing order
//	@Tags			order
//	@Param			data	body	models.Order	true	"JSON object of an order to be updated"
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/orders [patch]
func PatchOrder(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var o models.Order
		err := c.BodyParser(&o)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error parsing body for Order model")
		}

		cpy := o
		e := s.Contains(&cpy, "orders")
		if !e {
			return c.Status(http.StatusNoContent).SendString("Order does not exist")
		}

		err = s.Update(&o, "orders")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error updating order")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// order godoc
//
//	@Summary		Updates the status of a specific order
//	@Description	Returns the status of the given order ID
//	@Tags			order
//	@Param			id		path	int	true	"id of the order"		format(int)
//	@Param			status	path	int	true	"status of the order"	format(int)
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/orders/{id}/{status} [patch]
func PatchOrderStatus(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get ID and status and validate their values
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error parsing id")
		}
		status, err := c.ParamsInt("status")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error parsing status")
		}
		if status > int(models.Cancelled) || status <= int(models.Unknown) {
			return c.Status(http.StatusBadRequest).SendString("Error invalid status")
		}

		e := s.Contains(&models.Order{Model: models.Model{ID: uint(id)}}, "orders")
		if !e {
			return c.Status(http.StatusNoContent).SendString("Order does not exist")
		}

		var o models.Order
		err = s.Get(&o, "orders", "id = ?", id)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error getting order")
		}
		o.Status = models.Status(status)

		err = s.Update(&o, "orders")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Error updating order")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// order godoc
//
//	@Summary		Put an order
//	@Description	Puts an order, if it does not exist it is added otherwise it updates the order of the same ID
//	@Tags			order
//	@Param			data	body	models.Order	true	"JSON object of an order to be added or updated"
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Router			/api/v1/orders [put]
func PutOrder(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var o models.Order
		err := c.BodyParser(&o)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error parsing body for Order model")
		}

		e := s.Contains(&models.Order{Model: models.Model{ID: uint(o.ID)}}, "orders")
		if !e {
			err = s.Add(&o, "orders")
			if err != nil {
				c.Status(http.StatusBadRequest).SendString("Error adding Order")
			}

			return c.SendStatus(http.StatusNoContent)
		}

		err = s.Update(&o, "orders")
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error updating Order")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// order godoc
//
//	@Summary		Delete a specific order
//	@Description	Returns a status, depending on if the Order was deleted successfully or not
//	@Tags			order
//	@Param			id	path	int	true	"id of the order"	format(int)
//	@Produce		json
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/orders/{id} [delete]
func DeleteOrder(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error parsing id parameter of type integer")
		}

		err = s.Delete(&models.Order{Model: models.Model{ID: uint(id)}}, "orders")
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error deleting order")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
