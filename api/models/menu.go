package models

// Stores a menu consisting of dishes and drinks
type Menu struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
