package models

// A drink item, inherits from the MenuItem struct
type Drink struct {
	Model
	ItemRefer uint     `json:"itemRefer"`
	Item      MenuItem `gorm:"foreignKey:ItemRefer" json:"item"`
	Alcoholic bool     `gorm:"type:boolean;not null" json:"alcoholic"`
}
