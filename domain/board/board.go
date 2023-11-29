package board

import (
	"errors"
	"strings"
)

var (
	// ErrGameIsOver is returned when the game is over.
	ErrGameIsOver = errors.New("game is over")
	// ErrCellIsNotEmpty is returned when the cell is not empty.
	ErrCellIsNotEmpty = errors.New("cell is not empty")
)

const (
	boardSize = 3
)

// Board is a 3x3 matrix of CellValue.
// Board is the game board.
// Board is a value object so is immutable.
// Board is responsible for managing the board state and calculating the winner: X or 0 or the end of the game.
type Board [boardSize][boardSize]CellValue

// IsEmptyCell returns true if the cell is empty.
func (b Board) IsEmptyCell(cell Cell) bool {
	return b[cell.RowNumber][cell.ColumnNumber].IsEmpty()
}

// SetCellValue returns a new board with the set current turn cell value.
func (b Board) SetCellValue(cell Cell) (Board, error) {
	if !b.IsEmptyCell(cell) {
		return Board{}, ErrCellIsNotEmpty
	}
	if b.IsCompleted() {
		return Board{}, ErrGameIsOver
	}
	var r Board
	copy(r[:], b[:])
	r[cell.RowNumber][cell.ColumnNumber] = b.CurrentTurnCellValue()
	return r, nil
}

// String returns the string representation of the board.
func (b Board) String() string {
	return b.Sprint(nil)
}

// Sprint returns the string representation of the board.
func (b Board) Sprint(cursor *Cell) string {
	result := ""
	for i := 0; i < len(b); i++ {
		str := make([]string, len(b))
		for j := 0; j < len(b); j++ {
			v := b[i][j].String()
			if cursor != nil && i == cursor.RowNumber && j == cursor.ColumnNumber {
				str[j] = "[" + v + "]"
				continue
			}

			str[j] = " " + v + " "
		}
		result += strings.Join(str, " | ") + "\n"
	}
	return result
}

// MustSetCellValue returns a new board with the current turn cell value.
// It panics if the cell is not empty or the board is completed.
func (b Board) MustSetCellValue(cell Cell) Board {
	r, err := b.SetCellValue(cell)
	if err != nil {
		panic(err)
	}
	return r
}

// CurrentTurnCellValue returns the current turn cell value.
func (b Board) CurrentTurnCellValue() CellValue {
	c := b.FullCellsCount()
	if c%2 == 0 {
		return XValue
	}
	return OValue
}

// IsCompleted returns true if the game is completed.
// The game is completed when the board is full or there is a winner.
func (b Board) IsCompleted() bool {
	if _, exist := b.Winner(); exist {
		return true
	}
	if b.IsFull() {
		return true
	}
	return false
}

// FullCellsCount returns the number of full cells.
func (b Board) FullCellsCount() int {
	c := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if !b[i][j].IsEmpty() {
				c++
			}
		}
	}
	return c
}

// Winner returns the Winner CellValue if there is one, otherwise EmptyValue.
// It also returns true if there is a winner, otherwise false.
// The method support only 3x3 board.
func (b Board) Winner() (CellValue, bool) {
	for i := 0; i < boardSize; i++ {
		// cols
		if b[i][0] == b[i][1] && b[i][1] == b[i][2] && !b[i][0].IsEmpty() {
			return b[i][0], true
		}
		// rows
		if b[0][i] == b[1][i] && b[1][i] == b[2][i] && !b[0][i].IsEmpty() {
			return b[0][i], true
		}
	}
	// diagonals
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] && !b[0][0].IsEmpty() {
		return b[0][0], true
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] && !b[0][2].IsEmpty() {
		return b[0][2], true
	}
	return EmptyValue, false
}

// IsFull returns true if the board is full.
func (b Board) IsFull() bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j].IsEmpty() {
				return false
			}
		}
	}
	return true
}

// CellValue returns the cell value.
func (b Board) CellValue(cell Cell) CellValue {
	return b[cell.RowNumber][cell.ColumnNumber]
}

// MidCell returns the middle cell.
// It supports only 3x3 board.
func (b Board) MidCell() Cell {
	return Cell{RowNumber: 1, ColumnNumber: 1}
}

// Corners returns the corners of the board.
func (b Board) Corners() []Cell {
	return []Cell{
		{RowNumber: 0, ColumnNumber: 0},
		{RowNumber: 0, ColumnNumber: boardSize - 1},
		{RowNumber: boardSize - 1, ColumnNumber: 0},
		{RowNumber: boardSize - 1, ColumnNumber: boardSize - 1},
	}
}

// OpponentCellValue returns the opponent cell value.
func (b Board) OpponentCellValue() CellValue {
	curCellVal := b.CurrentTurnCellValue()
	if curCellVal == XValue {
		return OValue
	}
	return XValue
}

// SideCells returns the side cells of the board.
// It supports only 3x3 board.
func (b Board) SideCells() []Cell {
	return []Cell{
		{RowNumber: 0, ColumnNumber: 1},
		{RowNumber: 1, ColumnNumber: 0},
		{RowNumber: 1, ColumnNumber: 2},
		{RowNumber: 2, ColumnNumber: 1},
	}
}

// FindFirstEmptyCell returns the first empty cell.
func (b Board) FindFirstEmptyCell() *Cell {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j].IsEmpty() {
				return &Cell{RowNumber: i, ColumnNumber: j}
			}
		}
	}
	return nil
}
