package models

import "time"

// Lesson represents a guitar lesson
type Lesson struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Level     string    `json:"level"` // beginner, intermediate, advanced
	Content   string    `json:"content" gorm:"type:text"`
	Exercises []Exercise `json:"exercises" gorm:"type:jsonb"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Exercise represents an exercise within a lesson
type Exercise struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// TableName specifies the table name for Lesson model
func (Lesson) TableName() string {
	return "lessons"
}
