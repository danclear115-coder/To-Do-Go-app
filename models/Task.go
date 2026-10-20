package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Complexity  string `json:"complexity" gorm:"not null"`
	Priority    string `json:"priority" gorm:"not null"`
	Category    string `json:"category"`
	CreateTime  string `json:"сreateTime"`
	Deadline    string `json:"deadline"`
	IconImg     string `json:"iconImg"`
}
