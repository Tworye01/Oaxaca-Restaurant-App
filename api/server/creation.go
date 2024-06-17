package server

import (
	"errors"
	"fmt"
	"strconv"
	"team-project/database"
	"team-project/models"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var tables Tables

// Creates table
func (t *Tables) CreateTable(c *websocket.Conn, tableNo int) bool {
	if t.TableExists(tableNo) {
		return true
	}
	defer c.Close()
	table := &Table{TableNo: tableNo}
	t.Add(table)
	return false
}

// Handler for creation of the waiter/table
// websocket godoc
//
//	@Summary		Creates a waiter
//	@Description	Creates a waiter and assigns them to a table
//	@Tags			websocket
//	@Param			table	path	integer	true	"table number"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/staff/{table} [get]
func HandleCreationWaiter(s *database.Store) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		defer c.Close()
		tableNo := c.Params("table")
		number, err := strconv.Atoi(tableNo)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if !tables.CreateTable(c, number) {
			JoinTable(c, number, models.Role(1), s)
		}
		c.WriteMessage(websocket.TextMessage, []byte(CreatePacket("There is a waiter at this table")))
		JoinTable(c, number, models.Role(1), s)
	})
}

// Handler for the creation of a manager
// websocket godoc
//
//	@Summary		Creates a manager
//	@Description	Creates a manager and assigns them to all tables
//	@Tags			websocket
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/staff/manager [get]
func HandleCreationManager(s *database.Store) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		defer c.Close()
		number := 1000
		if !tables.CreateTable(c, number) {
			JoinTable(c, number, models.Role(3), s)
		}
		JoinTable(c, number, models.Role(3), s)
	})
}

// Handler for the creation of a manager
// websocket godoc
//
//	@Summary		Creates a customer
//	@Description	Creates a customer and assigns them to a table
//	@Tags			websocket
//	@Param			table	path	integer	true	"table number"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/order/{table} [get]
func HandleCreationCustomer(s *database.Store) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		defer c.Close()
		tableNo := c.Params("table")
		number, err := strconv.Atoi(tableNo)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		JoinTable(c, number, models.Role(0), s)
	})
}

// Adds user to the table array
func AddingUser(user *models.Users, tableNo int) error {
	table := tables.Get(tableNo)
	if table == nil {
		return errors.New("invalid table")
	}
	table.AddsUser(user)
	return nil
}

// Waiter/Customer joins table
func JoinTable(c *websocket.Conn, tableNo int, role models.Role, s *database.Store) {

	// Create a new user with the specified role
	user := models.Users{
		Conn: c,
		ListOfUsers: []models.Credentials{
			{Role: role}, // Add user with the specified role
		},
	}

	err := AddingUser(&user, tableNo)
	if err != nil {
		c.WriteMessage(websocket.TextMessage, []byte(CreatePacket("Invalid table!")))
		return
	}

	ta := tables.Get(tableNo)
	ta.BroadcastsEveryone([]byte(CreatePacket(fmt.Sprintf("%s has joined table number: %d", role.String(), tableNo))))
	UserWorker(c, &user, ta, s)
}
