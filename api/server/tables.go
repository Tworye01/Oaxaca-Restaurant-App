package server

import (
	"sync"
	"team-project/models"

	"github.com/gofiber/contrib/websocket"
)

type Table struct {
	mu       sync.Mutex      `json:"-"`
	TableNo  int             `json:"tableNo"`
	UserList []*models.Users `json:"customers"`
}

type Tables struct {
	mu     sync.Mutex
	Tables []*Table
}

// Check for waiter
func (t *Table) HasWaiter() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, t := range t.UserList {
		for _, j := range t.ListOfUsers {
			if j.Role == 1 {
				return true
			}
		}
	}
	return false
}

// Broadcast to everyone
func (t *Table) BroadcastsEveryone(m []byte) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, t := range t.UserList {
		t.Conn.WriteMessage(websocket.TextMessage, m)
	}
}

// Broadcast to waiter
func (t *Table) BroadcastsWaiter(m []byte) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, t := range t.UserList {
		for _, j := range t.ListOfUsers {
			if j.Role == 1 {
				t.Conn.WriteMessage(websocket.TextMessage, m)
			}
		}
	}
}

// Broadcasts to manager
func (t *Table) BroadcastsManager(m []byte) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, t := range t.UserList {
		for _, j := range t.ListOfUsers {
			if j.Role == 3 {
				t.Conn.WriteMessage(websocket.TextMessage, m)
			}
		}
	}
}

// Broadcasts to manager from all tables
func (t *Tables) BroadcastsManagers(m []byte) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, table := range t.Tables {
		table.BroadcastsManager(m)
	}
}

// Broadcast tocustomer
func (t *Table) BroadcastsCustomer(m []byte) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, t := range t.UserList {
		for _, j := range t.ListOfUsers {
			if j.Role == 0 {
				t.Conn.WriteMessage(websocket.TextMessage, m)
			}
		}
	}
}

// Remove customer
func (t *Table) RemoveWaiter(user *models.Users) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for i := range t.UserList {

		t.UserList[i] = t.UserList[len(t.UserList)-1]
		t.UserList = t.UserList[:len(t.UserList)-1]
		return
	}

}

// Adds customer to list
func (t *Table) AddsUser(user *models.Users) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.UserList = append(t.UserList, user)
}

// Checks if table exists
func (t *Tables) TableExists(tableNo int) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, ta := range t.Tables {
		if ta.TableNo == tableNo {
			return true
		}
	}
	return false
}

// Gets table number for tables
func (t *Tables) Get(tableNo int) *Table {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, ta := range t.Tables {
		if ta.TableNo == tableNo {
			return ta
		}
	}
	return nil
}

// Add table to tables
func (t *Tables) Add(ta *Table) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Tables = append(t.Tables, ta)
}

// Removes table from tables
func (t *Tables) RemoveTable(ta *Table) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for i, tab := range t.Tables {
		if ta.TableNo == tab.TableNo {
			t.Tables[i] = t.Tables[len(t.Tables)-1]
			t.Tables = t.Tables[:len(t.Tables)-1]
			return
		}
	}
}
