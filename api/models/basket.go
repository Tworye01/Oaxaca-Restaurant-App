package models

// Represents all items in the basket
type Basket struct {
	Items []BasketItem `json:"items"`
}
