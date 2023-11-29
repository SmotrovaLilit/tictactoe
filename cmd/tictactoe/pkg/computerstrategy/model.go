package computerstrategy

import (
	"tictactoe/cmd/tictactoe/pkg/choices"
	"tictactoe/domain/computer/strategies/minimax"
	"tictactoe/domain/computer/strategies/modifiedwiki"
	"tictactoe/domain/computer/strategies/wiki"
)

// NewModel creates a new computer strategy model.
func NewModel(questionTitle string) choices.Model {
	return choices.NewModel(
		[]string{
			"Wiki",
			"Modified Wiki",
			"Minimax",
		},
		[]any{
			wiki.NewStrategy(),
			modifiedwiki.NewStrategy(),
			minimax.NewStrategy(),
		},
		questionTitle,
	)
}
