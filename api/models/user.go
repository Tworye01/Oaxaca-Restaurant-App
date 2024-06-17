package models

import (
	"github.com/gofiber/contrib/websocket"
)

type Role int

const (
	Customer Role = iota // 0
	Waiter               // 1
	Kitchen              // 2
	Manager              // 3
	Admin                // 4
)

func (r Role) String() string {
	switch r {
	case Customer:
		return "Customer"
	case Waiter:
		return "Waiter"
	case Kitchen:
		return "Kitchen"
	case Manager:
		return "Manager"
	case Admin:
		return "Admin"
	default:
		return "Unknown"
	}
}

// Stores user details
type Credentials struct {
	Model
	Username string `gorm:"size:100;not null;unique" json:"username"`
	Password string `gorm:"size:100;not null" json:"password"`
	Role     Role   `gorm:"type:integer;not null; check:(Role >= 0) and (Role <= 4)" json:"role"`
}

// Stores list of users
type Users struct {
	Conn        *websocket.Conn `json:"-"`
	ListOfUsers []Credentials   `json:"credentials"`
}
