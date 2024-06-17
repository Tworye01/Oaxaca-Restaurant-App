package models

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type Allergens []string

// Scan in a string to produce an allergens type, when reading from database
func (a *Allergens) Scan(src any) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("src value cannot cast to string")
	}

	*a = strings.Split(string(str), ",")

	return nil
}

// Get an sql value of the Allergens, when writing to database
func (a Allergens) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}

	return strings.Join(a, ","), nil
}

// A menu item, inherits from the Model struct
type MenuItem struct {
	Model
	MenuRefer uint `json:"menuRefer"`
	Menu      Menu `gorm:"foreignKey:MenuRefer" json:"menu"`

	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"size:255;not null" json:"description"`
	Cost        float32   `gorm:"real;not null" json:"cost"`
	Calories    int       `gorm:"smallint;not null" json:"calories"`
	Allergens   Allergens `gorm:"type:varchar(255);default:'';" json:"allergens"`
	Available   bool      `gorm:"boolean;not null" json:"available"`
	Image       string    `gorm:"type:text" json:"image"`
}
