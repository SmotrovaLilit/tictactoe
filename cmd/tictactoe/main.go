package main

import (
	tea "github.com/charmbracelet/bubbletea"

	"tictactoe/cmd/tictactoe/computervscomputer"
	cmdGame "tictactoe/cmd/tictactoe/game"
	"tictactoe/cmd/tictactoe/humanvscomputer"
	"tictactoe/cmd/tictactoe/pkg/choices"
	"tictactoe/cmd/tictactoe/pkg/mode"
	"tictactoe/domain/game"
	"tictactoe/domain/player"
)

func main() {
	p := tea.NewProgram(newMainModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

type viewType int

const (
	viewTypeModeSelection viewType = iota
	viewTypeGame
)

type mainModel struct {
	chooseGameModeModel choices.Model
	gameModel           tea.Model
	currentView         viewType
}

func newMainModel() mainModel {
	return mainModel{
		chooseGameModeModel: mode.NewModel(),
		currentView:         viewTypeModeSelection,
	}
}

// Update updates a main model.
func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	switch m.currentView {
	case viewTypeModeSelection:
		child, _ := m.chooseGameModeModel.Update(msg)
		m.chooseGameModeModel = child.(choices.Model)
		gameMode, ok := m.chooseGameModeModel.GetSelected().(mode.Mode)
		if !ok {
			return m, nil
		}
		switch gameMode {
		case mode.HumanVsHuman:
			g := game.New(
				player.MustNew("Player 1"),
				player.MustNew("Player 2"),
			)
			m.gameModel = cmdGame.NewModel(*g)
		case mode.HumanVsComputer:
			m.gameModel = humanvscomputer.NewModel()
		case mode.ComputerVsComputer:
			m.gameModel = computervscomputer.NewModel()
		}

		m.currentView = viewTypeGame
	case viewTypeGame:
		child, cmd := m.gameModel.Update(msg)
		m.gameModel = child.(tea.Model)
		return m, cmd
	}
	return m, nil
}

// View returns a main model view.
func (m mainModel) View() string {
	footer := "\n\nPress q or ctrl + c to quit."
	header := "Tic Tac Toe\n\n"
	content := ""
	switch m.currentView {
	case viewTypeModeSelection:
		content = m.chooseGameModeModel.View()
	case viewTypeGame:
		content = m.gameModel.View()

	}
	return header + content + footer
}

// Init initializes a main model.
func (m mainModel) Init() tea.Cmd {
	return nil
}
