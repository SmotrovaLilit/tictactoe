package computer

import (
	"tictactoe/domain/board"

	"fmt"
)

const (
	computerPlayerName = "Computer"
)

// Strategy is the interface that wraps the basic FindBestCellForNextTurn method.
type Strategy interface {
	FindBestCellForNextTurn(b board.Board) board.Cell
	String() string
}

// Player is the computer player.
type Player struct {
	strategy Strategy
}

// New creates a new computer player.
func New(strategy Strategy) Player {
	if strategy == nil {
		panic("strategy is nil")
	}
	return Player{
		strategy: strategy,
	}
}

// Name returns the computer player name.
func (p Player) Name() string {
	return fmt.Sprintf(
		"%s/%s", computerPlayerName, p.strategy.String())
}

// String returns the string representation of the computer player.
func (p Player) String() string {
	return p.Name()
}

// GetNextCell returns the next turn cell.
func (p Player) GetNextCell(b board.Board) board.Cell {
	return p.strategy.FindBestCellForNextTurn(b)
}
