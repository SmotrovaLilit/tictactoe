package wiki

import (
	"tictactoe/domain/board"
	"tictactoe/domain/computer/boardhelper"
)

// Strategy is a strategies that is based on the wiki article:
// https://en.wikipedia.org/wiki/Tic-tac-toe#Strategy
type Strategy struct{}

// NewStrategy returns a new Strategy.
func NewStrategy() *Strategy {
	return &Strategy{}
}

// String returns the string representation of the Strategy.
func (s *Strategy) String() string {
	return "Wiki"
}

// FindBestCellForNextTurn finds the best cell for the next turn.
func (s *Strategy) FindBestCellForNextTurn(b board.Board) board.Cell {
	// Win: If the player has two in a row, they can place a third to get three in a row.
	if cell, _ := boardhelper.FindWinSituationsFor(b, b.CurrentTurnCellValue()); cell != nil {
		return *cell
	}
	// Block: If the opponent has two in a row, the player must play the third themselves to block the opponent.
	if cell, _ := boardhelper.FindWinSituationsFor(b, b.OpponentCellValue()); cell != nil {
		return *cell
	}
	// Fork: Cause a scenario where the player has two ways to win (two non-blocked lines of 2).
	if cell := findPossibleFork(b); cell != nil {
		return *cell
	}

	// Blocking an opponent's fork.
	if cell := findCellToBlockPossibleOpponentForks(b); cell != nil {
		return *cell
	}

	// Center: A player marks the center.
	// As long as it does not result in them producing a fork against current player.
	if mid := b.MidCell(); b.IsEmptyCell(mid) {
		return mid
	}

	// Opposite corner: If the opponent is in the corner, the player plays the opposite corner.
	if cell := findOppositeEmptyCorner(b); cell != nil {
		return *cell
	}

	// Empty corner: The player plays in a corner square.
	if cell := findEmptyCorner(b); cell != nil {
		return *cell
	}

	// Empty side: The player plays in a middle square on any of the 4 sides.
	if cell := findEmptySide(b); cell != nil {
		return *cell
	}
	panic("can't find best cell for next turn, it is impossible, fix the code")
}

// findOppositeEmptyCorner returns the opposite empty corner to opponent corner.
// If there are no opposite empty corner, returns nil.
func findOppositeEmptyCorner(b board.Board) *board.Cell {
	boardSize := len(b)
	cellBottomRight := board.MustNewCell(boardSize-1, boardSize-1)
	cellTopLeft := board.MustNewCell(0, 0)
	if b.CellValue(cellTopLeft) == b.OpponentCellValue() &&
		b.CellValue(cellBottomRight).IsEmpty() {
		return &cellBottomRight
	}
	if b.CellValue(cellTopLeft).IsEmpty() &&
		b.CellValue(cellBottomRight) == b.OpponentCellValue() {
		return &cellTopLeft
	}
	cellTopRight := board.MustNewCell(0, boardSize-1)
	cellBottomLeft := board.MustNewCell(boardSize-1, 0)
	if b.CellValue(cellTopRight) == b.OpponentCellValue() &&
		b.CellValue(cellBottomLeft).IsEmpty() {
		return &cellBottomLeft
	}
	if b.CellValue(cellTopRight).IsEmpty() &&
		b[boardSize-1][0] == b.OpponentCellValue() {
		return &cellTopRight
	}
	return nil
}

// findEmptyCorner returns the empty corner.
func findEmptyCorner(b board.Board) *board.Cell {
	corners := b.Corners()
	for _, cell := range corners {
		if b[cell.RowNumber][cell.ColumnNumber].IsEmpty() {
			return &cell
		}
	}
	return nil
}

// findEmptySide returns the empty side.
func findEmptySide(b board.Board) *board.Cell {
	cells := b.SideCells()
	for _, cell := range cells {
		if b[cell.RowNumber][cell.ColumnNumber].IsEmpty() {
			return &cell
		}
	}
	return nil
}
