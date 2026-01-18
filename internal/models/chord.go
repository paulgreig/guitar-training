package models

import "time"

// Chord represents a guitar chord
type Chord struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;uniqueIndex"`
	Notes     []string  `json:"notes" gorm:"type:jsonb"`
	Frets     []int     `json:"frets" gorm:"type:jsonb"`
	Fingers   []int     `json:"fingers" gorm:"type:jsonb"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Chord model
func (Chord) TableName() string {
	return "chords"
}
