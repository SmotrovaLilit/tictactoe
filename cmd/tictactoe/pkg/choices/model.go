package choices

import (
	tea "github.com/charmbracelet/bubbletea"

	"fmt"
)

// Model is the choices model.
type Model struct {
	choicesNames  []string
	choicesValues []any
	cursor        int
	selected      any
	questionTitle string
}

// NewModel creates a new choices model.
func NewModel(
	choicesNames []string,
	choicesValues []any,
	questionTitle string,
) Model {
	return Model{
		choicesNames:  choicesNames,
		choicesValues: choicesValues,
		questionTitle: questionTitle,
	}
}

// Update updates a choices model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choicesNames)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selected = m.choicesValues[m.cursor]
		}
	}

	return m, nil
}

// View returns a choices model view.
func (m Model) View() string {
	s := m.questionTitle + "\n\n"
	// Iterate over our choicesNames
	for i, choice := range m.choicesNames {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s [ ] %s\n", cursor, choice)
	}
	return s
}

// GetSelected returns the selected choice.
func (m Model) GetSelected() any {
	return m.selected
}

// Init initializes a choices model.
func (m Model) Init() tea.Cmd {
	return nil
}
