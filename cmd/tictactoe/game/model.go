package game

import (
	tea "github.com/charmbracelet/bubbletea"

	"tictactoe/domain/board"
	"tictactoe/domain/computer"
	"tictactoe/domain/game"

	"fmt"
)

// Model is the game model.
type Model struct {
	game   game.Game
	cursor *board.Cell
	err    error
}

// NewModel creates a new game model.
func NewModel(game game.Game) Model {
	// Play computer turn if it's the first player
	if p, ok := game.CurrentTurnPlayer().(computer.Player); ok {
		game.MustPlay(p.GetNextCell(game.GetBoard()))
	}
	cursor := game.GetBoard().FindFirstEmptyCell()

	return Model{
		game:   game,
		cursor: cursor,
	}
}

// Update handles messages from the Bubble Tea runtime.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	boardSize := len(m.game.GetBoard())
	m.err = nil
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		// move the cursor up to the next empty cell
		case "up", "k":
			if m.cursor.RowNumber > 0 {
				m.cursor.RowNumber--
			}

		// move the cursor down to the next empty cell
		case "down", "j":
			if m.cursor.RowNumber < boardSize-1 {
				m.cursor.RowNumber++
			}
		// move the cursor left to the next empty cell
		case "left", "h":
			if m.cursor.ColumnNumber > 0 {
				m.cursor.ColumnNumber--
			}
		// move the cursor right to the next empty cell
		case "right", "l":
			if m.cursor.ColumnNumber < boardSize-1 {
				m.cursor.ColumnNumber++
			}
		// play the current turn player
		case "enter":
			if m.game.IsOver() {
				return m, nil
			}
			err := m.game.Play(*m.cursor)
			if err != nil {
				m.err = err
				return m, nil
			}

			// Play computer turn
			if p, ok := m.game.CurrentTurnPlayer().(computer.Player); ok && !m.game.IsOver() {
				m.game.MustPlay(p.GetNextCell(m.game.GetBoard()))
			}
			// Move the cursor to the first empty cell
			emptyCell := m.game.GetBoard().FindFirstEmptyCell()
			if emptyCell != nil {
				m.cursor = emptyCell
			}
		}
	}

	return m, nil
}

// View renders the game model.
func (m Model) View() string {
	result := m.game.Sprint(m.cursor)
	if m.err != nil {
		result += fmt.Sprintf("\nError: %s", m.err)
	} else {
		result += "\n"
	}
	return result
}

// Init initializes the model before the game loop starts.
func (m Model) Init() tea.Cmd {
	return nil
}
