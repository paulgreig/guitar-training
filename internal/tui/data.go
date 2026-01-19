package tui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

type Scale struct {
	Name      string    `json:"name"`
	Notes     []string  `json:"notes"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Fret    int   `json:"fret"`
	Strings []int `json:"strings"`
}

type Lesson struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Level   string `json:"level"`
	Content string `json:"content"`
}

type ScalesLoadedMsg struct {
	Scales []Scale
}

type LessonsLoadedMsg struct {
	Lessons []Lesson
}

func loadScales() tea.Cmd {
	return func() tea.Msg {
		scales, err := loadScalesFromFile()
		if err != nil {
			// Return empty scales if file doesn't exist yet
			return ScalesLoadedMsg{Scales: []Scale{}}
		}
		return ScalesLoadedMsg{Scales: scales}
	}
}

func loadLessons() tea.Cmd {
	return func() tea.Msg {
		lessons, err := loadLessonsFromFile()
		if err != nil {
			// Return empty lessons if file doesn't exist yet
			return LessonsLoadedMsg{Lessons: []Lesson{}}
		}
		return LessonsLoadedMsg{Lessons: lessons}
	}
}

func loadScalesFromFile() ([]Scale, error) {
	// Try to load from data/scales.json
	dataPath := filepath.Join("data", "scales.json")
	
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("could not read scales file: %w", err)
	}
	
	var scales []Scale
	if err := json.Unmarshal(data, &scales); err != nil {
		return nil, fmt.Errorf("could not parse scales: %w", err)
	}
	
	return scales, nil
}

func loadLessonsFromFile() ([]Lesson, error) {
	// Try to load from data/lessons.json
	dataPath := filepath.Join("data", "lessons.json")
	
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("could not read lessons file: %w", err)
	}
	
	var lessons []Lesson
	if err := json.Unmarshal(data, &lessons); err != nil {
		return nil, fmt.Errorf("could not parse lessons: %w", err)
	}
	
	return lessons, nil
}
