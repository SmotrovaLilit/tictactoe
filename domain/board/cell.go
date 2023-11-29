package board

import "errors"

var (
	// ErrInvalidCell is returned when a cell is invalid.
	ErrInvalidCell = errors.New("invalid Cell, Cell must be between 0 and 2, inclusive")
)

// Cell represents a cell in the board.
type Cell struct {
	RowNumber    int
	ColumnNumber int
}

// NewCell creates a Cell.
func NewCell(row, coll int) (Cell, error) {
	if row < 0 || coll < 0 || row > boardSize-1 || coll > boardSize-1 {
		return Cell{}, ErrInvalidCell
	}
	return Cell{
		RowNumber:    row,
		ColumnNumber: coll,
	}, nil
}

// MustNewCell returns a new Cell or panics.
func MustNewCell(row, coll int) Cell {
	c, err := NewCell(row, coll)
	if err != nil {
		panic(err)
	}
	return c
}
