package models

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Complexity  string `json:"complexity" gorm:"not null"`
	Priority    string `json:"priority" gorm:"not null"`
	Category    string `json:"category"`
	CreateTime  string `json:"createTime"`
	Deadline    string `json:"deadline"`
	IconImg     string `json:"iconImg"`
}