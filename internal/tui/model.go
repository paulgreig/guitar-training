package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	// Current view state
	view string // "menu", "scales", "lessons", "scale-detail", "lesson-detail"
	
	// Data
	scales  []Scale
	lessons []Lesson
	
	// Navigation
	selectedIndex int
	cursor        int
	
	// Styles
	styles Styles
}

type Styles struct {
	Title    lipgloss.Style
	Menu     lipgloss.Style
	Selected lipgloss.Style
	Text     lipgloss.Style
}

func NewModel() Model {
	return Model{
		view:          "menu",
		selectedIndex: 0,
		cursor:        0,
		styles:        defaultStyles(),
	}
}

func defaultStyles() Styles {
	return Styles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("63")).
			Padding(1, 2),
		Menu: lipgloss.NewStyle().
			PaddingLeft(2),
		Selected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true),
		Text: lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")),
	}
}

func (m Model) Init() tea.Cmd {
	// Load scales and lessons data
	return tea.Batch(loadScales(), loadLessons())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			maxItems := m.getMaxItems()
			if m.cursor < maxItems-1 {
				m.cursor++
			}
		case "enter":
			return m.handleEnter()
		case "esc":
			if m.view != "menu" {
				m.view = "menu"
				m.cursor = 0
			}
		}
	case ScalesLoadedMsg:
		m.scales = msg.Scales
	case LessonsLoadedMsg:
		m.lessons = msg.Lessons
	}

	return m, nil
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	switch m.view {
	case "menu":
		switch m.cursor {
		case 0:
			m.view = "scales"
			m.cursor = 0
		case 1:
			m.view = "lessons"
			m.cursor = 0
		case 2:
			return m, tea.Quit
		}
	case "scales":
		if len(m.scales) > 0 && m.cursor < len(m.scales) {
			m.view = "scale-detail"
			m.selectedIndex = m.cursor
		}
	case "lessons":
		if len(m.lessons) > 0 && m.cursor < len(m.lessons) {
			m.view = "lesson-detail"
			m.selectedIndex = m.cursor
		}
	}
	return m, nil
}

func (m Model) View() string {
	switch m.view {
	case "menu":
		return m.renderMenu()
	case "scales":
		return m.renderScalesList()
	case "lessons":
		return m.renderLessonsList()
	case "scale-detail":
		return m.renderScaleDetail()
	case "lesson-detail":
		return m.renderLessonDetail()
	default:
		return "Unknown view"
	}
}

func (m Model) renderMenu() string {
	title := m.styles.Title.Render("🎸 Guitar Training")
	
	menuItems := []string{
		"View Scales",
		"View Lessons",
		"Quit",
	}
	
	var menu string
	for i, item := range menuItems {
		if i == m.cursor {
			menu += m.styles.Selected.Render("> " + item) + "\n"
		} else {
			menu += m.styles.Menu.Render("  " + item) + "\n"
		}
	}
	
	help := m.styles.Text.Render("\nUse ↑/↓ to navigate, Enter to select, q to quit")
	
	return lipgloss.JoinVertical(lipgloss.Left, title, menu, help)
}

func (m Model) renderScalesList() string {
	title := m.styles.Title.Render("Guitar Scales")
	
	if len(m.scales) == 0 {
		return title + "\n\nNo scales loaded."
	}
	
	var list string
	for i, scale := range m.scales {
		if i == m.cursor {
			list += m.styles.Selected.Render(fmt.Sprintf("> %s", scale.Name)) + "\n"
		} else {
			list += m.styles.Menu.Render(fmt.Sprintf("  %s", scale.Name)) + "\n"
		}
	}
	
	help := m.styles.Text.Render("\nUse ↑/↓ to navigate, Enter to view details, Esc to go back")
	
	return lipgloss.JoinVertical(lipgloss.Left, title, list, help)
}

func (m Model) renderLessonsList() string {
	title := m.styles.Title.Render("Guitar Lessons")
	
	if len(m.lessons) == 0 {
		return title + "\n\nNo lessons loaded."
	}
	
	var list string
	for i, lesson := range m.lessons {
		if i == m.cursor {
			list += m.styles.Selected.Render(fmt.Sprintf("> %s [%s]", lesson.Title, lesson.Level)) + "\n"
		} else {
			list += m.styles.Menu.Render(fmt.Sprintf("  %s [%s]", lesson.Title, lesson.Level)) + "\n"
		}
	}
	
	help := m.styles.Text.Render("\nUse ↑/↓ to navigate, Enter to view details, Esc to go back")
	
	return lipgloss.JoinVertical(lipgloss.Left, title, list, help)
}

func (m Model) renderScaleDetail() string {
	if m.selectedIndex >= len(m.scales) {
		return "Scale not found"
	}
	
	scale := m.scales[m.selectedIndex]
	title := m.styles.Title.Render(scale.Name)
	
	content := m.styles.Text.Render(fmt.Sprintf("Notes: %v\n\n", scale.Notes))
	
	// Render scale positions on fretboard
	content += m.renderFretboard(scale)
	
	help := m.styles.Text.Render("\nPress Esc to go back")
	
	return lipgloss.JoinVertical(lipgloss.Left, title, content, help)
}

func (m Model) renderLessonDetail() string {
	if m.selectedIndex >= len(m.lessons) {
		return "Lesson not found"
	}
	
	lesson := m.lessons[m.selectedIndex]
	title := m.styles.Title.Render(lesson.Title)
	
	level := m.styles.Text.Render(fmt.Sprintf("Level: %s\n", lesson.Level))
	content := m.styles.Text.Render(lesson.Content)
	
	help := m.styles.Text.Render("\nPress Esc to go back")
	
	return lipgloss.JoinVertical(lipgloss.Left, title, level, content, help)
}

func (m Model) renderFretboard(scale Scale) string {
	// Simple text-based fretboard representation
	// E|--0--1--2--3--4--5--
	// B|--0--1--2--3--4--5--
	// etc.
	
	fretboard := "Fretboard:\n"
	strings := []string{"E", "A", "D", "G", "B", "e"}
	
	for i, str := range strings {
		line := fmt.Sprintf("%s|", str)
		for fret := 0; fret <= 12; fret++ {
			// Check if this position is in the scale
			marked := false
			for _, pos := range scale.Positions {
				if pos.Fret == fret && contains(pos.Strings, i) {
					line += "--●--"
					marked = true
					break
				}
			}
			if !marked {
				line += "-----"
			}
		}
		fretboard += line + "\n"
	}
	
	return m.styles.Text.Render(fretboard)
}

func (m Model) getMaxItems() int {
	switch m.view {
	case "menu":
		return 3 // Menu has 3 items
	case "scales":
		return len(m.scales)
	case "lessons":
		return len(m.lessons)
	default:
		return 1
	}
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
