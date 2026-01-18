package models

import "time"

// PracticeSession represents a practice session record
type PracticeSession struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	UserID       string    `json:"user_id" gorm:"index"` // Optional for multi-user
	Date         time.Time `json:"date" gorm:"index"`
	Duration     int       `json:"duration"` // in minutes
	ExerciseType string    `json:"exercise_type"`
	Score        float64   `json:"score"`
	Rating       int       `json:"rating"` // 1-5
	Notes        string    `json:"notes" gorm:"type:text"`
	CreatedAt    time.Time `json:"created_at"`
}

// TableName specifies the table name for PracticeSession model
func (PracticeSession) TableName() string {
	return "practice_sessions"
}
