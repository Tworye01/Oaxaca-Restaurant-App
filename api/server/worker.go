package server

import (
	"fmt"
	"strconv"
	"team-project/database"
	"team-project/models"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type UserData struct {
	Type any    `json:"type"`
	Data string `json:"body"`
}

func UserWorker(c *websocket.Conn, p *models.Users, t *Table, s *database.Store) {
	// Continuously poll for messages from the user
	for {
		var o models.Order
		var data UserData
		err := c.ReadJSON(&data)
		if err != nil {
			// Disconnect if no data
			fmt.Printf("User has disconnected from table %d\n", t.TableNo)
			break
		}

		// If data retrieved is an alert then ping the waiter
		if data.Type == "PING" && data.Data == "ALERT" {
			t.BroadcastsWaiter([]byte(CreatePacket(fmt.Sprintf("Table %d needs you", t.TableNo))))
		}

		// If data retrieved is about the table it works on the Status
		if data.Type == "ORDER" {
			b, err := strconv.Atoi(data.Data)
			if err != nil {
				fmt.Println("Can't convert this to an int!")
			}

			// Keep track of the previous status
			o.Status = 0
			var prevStatus = o.Status

			// Loops every 5 seconds
			ticker := time.NewTicker(1 * time.Second)
			quit := make(chan struct{})
			go func() {
				for {
					select {
					case <-ticker.C:
						err = s.Get(&o, "orders", "id = ?", b)
						if err != nil {
							fmt.Printf("User has disconnected from table %d\n", t.TableNo)
							break
						}

						// Check if the status has changed
						if o.Status != prevStatus {
							switch o.Status {
							case 1:
								t.BroadcastsCustomer([]byte(CreatePacket(fmt.Sprintf("Order: %d is being prepared", b))))
								tables.BroadcastsManagers([]byte(CreatePacket(fmt.Sprintf("Order: %d is being prepared for table: %d", b, t.TableNo))))
								t.BroadcastsWaiter([]byte(CreatePacket(fmt.Sprintf("Order: %d is being prepared for table: %d", b, t.TableNo))))
							case 2:
								t.BroadcastsCustomer([]byte(CreatePacket(fmt.Sprintf("Order: %d is cooking", b))))
								tables.BroadcastsManagers([]byte(CreatePacket(fmt.Sprintf("Order: %d is being cooked for table: %d", b, t.TableNo))))
								t.BroadcastsWaiter([]byte(CreatePacket(fmt.Sprintf("Order: %d is being cooked for table: %d", b, t.TableNo))))
							case 3:
								t.BroadcastsCustomer([]byte(CreatePacket(fmt.Sprintf("Order: %d is waiting on server", b))))
								tables.BroadcastsManagers([]byte(CreatePacket(fmt.Sprintf("Order: %d is waiting on server for table: %d", b, t.TableNo))))
								t.BroadcastsWaiter([]byte(CreatePacket(fmt.Sprintf("Order: %d is ready for table: %d", b, t.TableNo))))
							case 4:
								t.BroadcastsCustomer([]byte(CreatePacket(fmt.Sprintf("Order: %d is being served", b))))
								tables.BroadcastsManagers([]byte(CreatePacket(fmt.Sprintf("Order: %d is being served for table: %d", b, t.TableNo))))
								t.BroadcastsWaiter([]byte(CreatePacket(fmt.Sprintf("Order: %d is being served for table: %d", b, t.TableNo))))
							case 5:
								t.BroadcastsCustomer([]byte(CreatePacket(fmt.Sprintf("Order: %d has been completed", b))))
								tables.BroadcastsManagers([]byte(CreatePacket(fmt.Sprintf("Order: %d has been completed for table: %d", b, t.TableNo))))
								t.BroadcastsWaiter([]byte(CreatePacket(fmt.Sprintf("Order: %d  is done for table: %d", b, t.TableNo))))
								close(quit)
							case 6:
								t.BroadcastsCustomer([]byte(CreatePacket(fmt.Sprintf("Order: %d has been cancelled", b))))
								tables.BroadcastsManagers([]byte(CreatePacket(fmt.Sprintf("Order: %d has been cancelled for table: %d", b, t.TableNo))))
								t.BroadcastsWaiter([]byte(CreatePacket(fmt.Sprintf("Order: %d  is cancelled for table: %d", b, t.TableNo))))
								close(quit)
							}

							// Previous status to current
							prevStatus = o.Status
						}
					case <-quit:
						ticker.Stop()
						return
					}
				}
			}()
		}
	}
	// Remove waiter from list and if no waiter, then it removes the session
	t.RemoveWaiter(p)
	if len(t.UserList) == 0 {
		tables.RemoveTable(t)
	}
}
