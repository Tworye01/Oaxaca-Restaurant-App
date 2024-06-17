package models

// Represents a tables bill
type Bill struct {
	Model
	TableRefer int     `gorm:"not null" json:"table"`
	Cost       float32 `gorm:"real;not null" json:"cost"`
	Paid       bool    `gorm:"default:false;not null" json:"paid"`
}
