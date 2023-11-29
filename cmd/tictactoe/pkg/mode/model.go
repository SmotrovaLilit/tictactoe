package mode

import (
	"tictactoe/cmd/tictactoe/pkg/choices"
)

// Mode represents the game Mode.
type Mode int

const (
	// HumanVsHuman represents the human vs human game Mode.
	HumanVsHuman Mode = iota + 1
	// HumanVsComputer represents the human vs computer game Mode.
	HumanVsComputer
	// ComputerVsComputer represents the computer vs computer game Mode.
	ComputerVsComputer
)

var modeNames = map[Mode]string{
	HumanVsHuman:       "Human vs Human",
	HumanVsComputer:    "Human vs Computer",
	ComputerVsComputer: "Computer vs Computer",
}

// String returns the string representation of the GameMode
func (g Mode) String() string {
	return modeNames[g]
}

// NewModel creates a new game Mode model.
func NewModel() choices.Model {
	return choices.NewModel(
		[]string{
			HumanVsHuman.String(),
			HumanVsComputer.String(),
			ComputerVsComputer.String(),
		},
		[]any{
			HumanVsHuman,
			HumanVsComputer,
			ComputerVsComputer,
		},
		"Select game Mode:",
	)
}
