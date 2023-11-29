package minimax

import (
	"math"

	"tictactoe/domain/board"
)

// Strategy is a computer strategy that implements the minimax algorithm.
type Strategy struct {
}

// NewStrategy returns a new Strategy.
func NewStrategy() *Strategy {
	return &Strategy{}
}

// String returns the string representation of the Strategy.
func (s *Strategy) String() string {
	return "Minimax"
}

// FindBestCellForNextTurn finds the best cell for the next turn.
func (s *Strategy) FindBestCellForNextTurn(b board.Board) board.Cell {
	bestVal := math.MinInt64
	var bestMove board.Cell
	boardSize := len(b)
	playerCellValue := b.CurrentTurnCellValue()
	opponentCellValue := b.OpponentCellValue()
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j].IsEmpty() {
				b[i][j] = playerCellValue
				moveVal := minimax(b, 0, false, playerCellValue, opponentCellValue)
				b[i][j] = board.EmptyValue

				if moveVal > bestVal {
					bestMove = board.MustNewCell(i, j)
					bestVal = moveVal
				}
			}
		}
	}
	return bestMove
}

// minimax is the minimax algorithm.
func minimax(b board.Board, depth int, isMax bool, playerCellValue, opponentCellValue board.CellValue) int {
	boardSize := len(b)
	if cellValue, ex := b.Winner(); ex {
		if cellValue == playerCellValue {
			return 10
		}
		return -10
	}
	if b.IsFull() {
		return 0
	}

	if isMax {
		best := math.MinInt64
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if b[i][j].IsEmpty() {
					b[i][j] = playerCellValue
					best = max(best, minimax(b, depth+1, !isMax, playerCellValue, opponentCellValue))
					b[i][j] = board.EmptyValue
				}
			}
		}
		return best
	} else {
		best := math.MaxInt64

		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if b[i][j].IsEmpty() {
					b[i][j] = opponentCellValue
					best = min(best, minimax(b, depth+1, !isMax, playerCellValue, opponentCellValue))
					b[i][j] = board.EmptyValue
				}
			}
		}
		return best
	}
}
