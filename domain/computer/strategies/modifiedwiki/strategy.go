package modifiedwiki

import (
	"tictactoe/domain/board"
	"tictactoe/domain/computer/boardhelper"
	"tictactoe/domain/computer/strategies/wiki"
)

// Strategy is a strategies that is based on the wiki article:
// https://en.wikipedia.org/wiki/Tic-tac-toe#Strategy
// but modified by me to win more often because
// wiki algorithm is too predictable for opponent in the first 3 turns,
// and it leads to the draw.
type Strategy struct {
	wiki.Strategy
}

// NewStrategy returns a new Strategy.
func NewStrategy() *Strategy {
	return &Strategy{
		Strategy: *wiki.NewStrategy(),
	}
}

// String returns the string representation of the Strategy.
func (s *Strategy) String() string {
	return "ModifiedWiki"
}

// FindBestCellForNextTurn finds the best cell for the next turn.
func (s *Strategy) FindBestCellForNextTurn(b board.Board) board.Cell {
	turnNumber := b.FullCellsCount() + 1
	if turnNumber <= 3 { // It is my own Strategy for first 3 turns,
		// I think it is more win Strategy than wiki algorithm.
		if b.IsEmptyCell(b.MidCell()) {
			return b.MidCell()
		}
		// Return corner that do not create win situation for current player,
		// Because it is too predictable for opponent, and it leads to the draw.
		if cell := findEmptyCornerThatNotCreateWinSituation(b); cell != nil {
			return *cell
		}
	}
	return s.Strategy.FindBestCellForNextTurn(b)
}

// findEmptyCornerThatNotCreateWinSituation returns the empty corner that not create 2 in a line.
// Because it is too predictable for opponent, and it leads to the draw.
func findEmptyCornerThatNotCreateWinSituation(b board.Board) *board.Cell {
	corners := b.Corners()
	for _, corner := range corners {
		if !b.IsEmptyCell(corner) {
			continue
		}
		nextBoard := b.MustSetCellValue(corner)

		// win situation on this step is too predictable, so we will not use it
		if c, _ := boardhelper.FindWinSituationsFor(nextBoard, b.CurrentTurnCellValue()); c != nil {
			continue
		}
		return &corner
	}
	return nil
}
