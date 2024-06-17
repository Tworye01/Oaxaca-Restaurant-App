package models

// Represents an item in the basket
type BasketItem struct {
	ID     int    `json:"id"`
	ItemID int    `json:"itemId"`
	Note   string `json:"note"`
}
