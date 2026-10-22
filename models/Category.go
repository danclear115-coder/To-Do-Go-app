package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Category    string  `json:"category"`
	Description string 	`json:"description"`
}
