package models

import (
	"errors"
)

// What is the current status of an order?
type Status int

const (
	Unknown     Status = iota // 0
	Preparation               // 1
	Cooking                   // 2
	Waiting                   // 3
	Serving                   // 4
	Done                      // 5
	Cancelled                 // 6

)

var Statuses = map[Status]string{
	Unknown:     "Unknown",
	Preparation: "Preparation",
	Cooking:     "Cooking",
	Waiting:     "Waiting",
	Serving:     "Serving", // slay
	Done:        "Done",
	Cancelled:   "Cancelled",
}

func (s Status) String() (string, error) {
	str, v := Statuses[s]
	if !v {
		return "", errors.New("error, failed to find status")
	}

	return str, nil
}

type Order struct {
	Model
	TableRefer int     `gorm:"not null" json:"table"`
	Status     Status  `gorm:"not null;type:integer; check:(Status >= 0) and (Status <= 6)" json:"status"`
	Cost       float32 `gorm:"real;not null;default:0" json:"cost"`
	Note       string  `gorm:"type:text" json:"note"`

	ItemRefer uint     `json:"itemRefer"`
	Item      MenuItem `gorm:"foreignKey:ItemRefer" json:"item"`
	BillRefer uint     `json:"billRefer"`
	Bill      Bill     `gorm:"foreignKey:BillRefer" json:"bill"`
}
