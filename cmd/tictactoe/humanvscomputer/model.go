package humanvscomputer

import (
	tea "github.com/charmbracelet/bubbletea"
	"tictactoe/cmd/tictactoe/pkg/choices"
	"tictactoe/cmd/tictactoe/pkg/computerstrategy"

	cmdGame "tictactoe/cmd/tictactoe/game"
	"tictactoe/domain/computer"
	"tictactoe/domain/game"
	"tictactoe/domain/player"
)

type viewType int

const (
	viewTypeComputerStrategySelection viewType = iota + 1
	viewTypeChooseFirstPlayer
	viewTypeGame
)

// Model represents the human vs computer model.
type Model struct {
	gameModel              cmdGame.Model
	computerStrategyModel  choices.Model
	chooseFirstPlayerModel choices.Model
	currentView            viewType
	computer               game.Player
	player                 game.Player
}

// NewModel creates a new human vs computer model.
func NewModel() Model {
	return Model{
		currentView:           viewTypeComputerStrategySelection,
		computerStrategyModel: computerstrategy.NewModel("Choose computer strategy:"),
	}
}

// Update updates a human vs computer model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.currentView {
	case viewTypeComputerStrategySelection:
		child, _ := m.computerStrategyModel.Update(msg)
		m.computerStrategyModel = child.(choices.Model)
		s, ok := m.computerStrategyModel.GetSelected().(computer.Strategy)
		if ok {
			m.computer = computer.New(s)
			m.player = player.MustNew("Player")
			m.chooseFirstPlayerModel = chooseFirstPlayer(m.computer, m.player)
			m.currentView = viewTypeChooseFirstPlayer
		}
	case viewTypeChooseFirstPlayer:
		child, _ := m.chooseFirstPlayerModel.Update(msg)
		m.chooseFirstPlayerModel = child.(choices.Model)
		firstPlayer, ok := m.chooseFirstPlayerModel.GetSelected().(game.Player)
		if ok {
			p1 := firstPlayer
			p2 := m.computer
			if p1 == m.computer {
				p2 = m.player
			}
			g := game.New(p1, p2)
			m.gameModel = cmdGame.NewModel(*g)
			m.currentView = viewTypeGame
		}
	case viewTypeGame:
		child, _ := m.gameModel.Update(msg)
		m.gameModel = child.(cmdGame.Model)
	}
	return m, nil
}

// View returns a human vs computer model view.
func (m Model) View() string {
	switch m.currentView {
	case viewTypeComputerStrategySelection:
		return m.computerStrategyModel.View()
	case viewTypeChooseFirstPlayer:
		return m.chooseFirstPlayerModel.View()
	case viewTypeGame:
		return m.gameModel.View()
	}
	return ""
}

// chooseFirstPlayer creates a new choose first player model.
func chooseFirstPlayer(player1, player2 game.Player) choices.Model {
	return choices.NewModel(
		[]string{
			player1.Name(),
			player2.Name(),
		},
		[]any{
			player1,
			player2,
		},
		"Choose first player:",
	)
}

// Init initializes a choices model.
func (m Model) Init() tea.Cmd {
	return nil
}
