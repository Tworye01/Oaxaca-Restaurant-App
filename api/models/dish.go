package models

import "errors"

type Course int

const (
	Starter Course = iota // 0
	Main                  // 1
	Dessert               // 2
	Side                  // 3
	Extra                 // 4
)

// Course to string map for getting the name of a course
var Courses = map[Course]string{
	Starter: "Starter",
	Main:    "Main",
	Dessert: "Dessert",
	Side:    "Side",
	Extra:   "Extra",
}

// Get the string value of a course
func (c Course) String() (string, error) {
	str, v := Courses[c]
	if !v {
		return "", errors.New("error, failed to find course ")
	}

	return str, nil
}

type Spice int

const (
	Mild     Spice = iota // 0
	Medium                // 1
	Hot                   // 2
	ExtraHot              // 3
)

// Spice to string map for getting the name of a spice level
var Spices = map[Spice]string{
	Mild:     "Mild",
	Medium:   "Medium",
	Hot:      "Hot",
	ExtraHot: "Extra Hot",
}

// Get a string value of the Spice
func (s Spice) String() (string, error) {
	str, v := Spices[s]
	if !v {
		return "", errors.New("error, failed to find course ")
	}

	return str, nil
}

// Dish item, inherits from the MenuItem struct
type Dish struct {
	Model
	ItemRefer  uint     `json:"itemRefer"`
	Item       MenuItem `gorm:"foreignKey:ItemRefer" json:"item"`
	Course     Course   `gorm:"not null;type:integer; check:(Course >= 0) and (Course <= 4)" json:"course"`
	Spice      Spice    `gorm:"not null;type:integer; check:(Spice >= 0) and (Spice <= 3)" json:"spice"`
	Vegetarian bool     `gorm:"not null;default:false" json:"vegetarian"`
}
