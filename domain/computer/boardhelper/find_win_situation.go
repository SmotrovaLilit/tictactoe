package boardhelper

import "tictactoe/domain/board"

// FindWinSituationsFor returns any cell that will create boardSize-1 in line and the count of win situations.
func FindWinSituationsFor(b board.Board, value board.CellValue) (winCell *board.Cell, count int) {
	boardSize := len(b)
	rowsSum := make([]int, boardSize)
	colsSum := make([]int, boardSize)
	diag1Sum := 0
	diag2Sum := 0
	neededCountInLine := boardSize - 1
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			rowsSum[i] += int(b[i][j])
			colsSum[i] += int(b[j][i])
			if i == j {
				diag1Sum += int(b[i][j])
			}
			if i+j == boardSize-1 {
				diag2Sum += int(b[i][j])
			}
		}
		if rowsSum[i] == int(value)*neededCountInLine {
			count++
			// Find the empty cell
			for j := 0; j < boardSize; j++ {
				if b[i][j].IsEmpty() {
					winCell = &board.Cell{RowNumber: i, ColumnNumber: j}
				}
			}
		}
		if colsSum[i] == int(value)*neededCountInLine {
			count++
			// Find the empty cell
			for j := 0; j < boardSize; j++ {
				if b[j][i].IsEmpty() {
					winCell = &board.Cell{RowNumber: j, ColumnNumber: i}
				}
			}
		}
	}
	if diag1Sum == int(value)*neededCountInLine {
		count++
		// Find the empty cell
		for j := 0; j < boardSize; j++ {
			if b[j][j].IsEmpty() {
				winCell = &board.Cell{RowNumber: j, ColumnNumber: j}
			}
		}
	}
	if diag2Sum == int(value)*neededCountInLine {
		count++
		// Find the empty cell
		for j := 0; j < boardSize; j++ {
			if b[j][boardSize-1-j].IsEmpty() {
				winCell = &board.Cell{RowNumber: j, ColumnNumber: boardSize - 1 - j}
			}
		}
	}
	return winCell, count
}
