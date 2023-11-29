package computervscomputer

import (
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"tictactoe/cmd/tictactoe/pkg/choices"
	"tictactoe/cmd/tictactoe/pkg/computerstrategy"

	"tictactoe/domain/computer"
	"tictactoe/domain/game"

	"time"
)

type viewType int

const (
	viewTypeComputer1StrategySelection viewType = iota + 1
	viewTypeComputer2StrategySelection viewType = iota + 1
	viewTypeGame
)

// Model represents the computer  vs computer model.
type Model struct {
	game                   *game.Game
	computerStrategyModel  choices.Model
	chooseFirstPlayerModel choices.Model
	currentView            viewType
	computer1              game.Player
	computer2              game.Player
	timer                  timer.Model
}

// NewModel creates a new computer vs computer model.
func NewModel() Model {
	return Model{
		timer:                 timer.NewWithInterval(time.Minute*5, 500*time.Millisecond),
		currentView:           viewTypeComputer1StrategySelection,
		computerStrategyModel: computerstrategy.NewModel("Choose strategy for first computer:"),
	}
}

// Update updates a computer vs computer  model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Computer wil play its turn when the timer ticks
	case timer.TickMsg:
		var cmd tea.Cmd
		if m.game.IsOver() {
			cmd := m.timer.Stop()
			return m, cmd
		}
		p := m.game.CurrentTurnPlayer().(computer.Player)
		m.game.MustPlay(p.GetNextCell(m.game.GetBoard()))
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	}
	switch m.currentView {
	case viewTypeComputer1StrategySelection:
		child, _ := m.computerStrategyModel.Update(msg)
		m.computerStrategyModel = child.(choices.Model)
		s, ok := m.computerStrategyModel.GetSelected().(computer.Strategy)
		if ok {
			m.computer1 = computer.New(s)
			m.computerStrategyModel = computerstrategy.NewModel("Choose strategy for second computer:")
			m.currentView = viewTypeComputer2StrategySelection
		}
	case viewTypeComputer2StrategySelection:
		child, _ := m.computerStrategyModel.Update(msg)
		m.computerStrategyModel = child.(choices.Model)
		s, ok := m.computerStrategyModel.GetSelected().(computer.Strategy)
		if ok {
			m.computer2 = computer.New(s)
			m.game = game.New(m.computer1, m.computer2)
			cmd := m.timer.Start()
			m.currentView = viewTypeGame
			return m, cmd
		}
	}
	return m, nil
}

// View returns a computer vs computer model view.
func (m Model) View() string {
	switch m.currentView {
	case viewTypeComputer1StrategySelection, viewTypeComputer2StrategySelection:
		return m.computerStrategyModel.View()
	case viewTypeGame:
		return m.game.Sprint(nil)
	}
	return ""
}

func (m Model) Init() tea.Cmd {
	return m.timer.Init()
}
