package models

import "time"

// Scale represents a guitar scale
type Scale struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;uniqueIndex"`
	Notes     []string  `json:"notes" gorm:"type:jsonb"`
	Positions []Position `json:"positions" gorm:"type:jsonb"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Position represents a scale position on the fretboard
type Position struct {
	Fret    int   `json:"fret"`
	Strings []int `json:"strings"`
}

// TableName specifies the table name for Scale model
func (Scale) TableName() string {
	return "scales"
}
