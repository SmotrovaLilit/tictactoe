package wiki

import (
	"tictactoe/domain/board"
	"tictactoe/domain/computer/boardhelper"
)

// findPossibleFork returns the cell that will create fork against opponent.
func findPossibleFork(b board.Board) *board.Cell {
	boardSize := len(b)
	count := b.FullCellsCount()
	if count > boardSize*boardSize-3 { // it is already end  part of the game, so the player can't create a fork
		return nil
	}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if !b[i][j].IsEmpty() {
				continue
			}
			nextBoard := b.MustSetCellValue(board.Cell{RowNumber: i, ColumnNumber: j})
			if _, count := boardhelper.FindWinSituationsFor(
				nextBoard,
				b.CurrentTurnCellValue(),
			); count > 1 {
				return &board.Cell{RowNumber: i, ColumnNumber: j}
			}
		}
	}
	return nil
}

// findCellToBlockPossibleOpponentForks returns a cell to block possible opponent forks.
// If there are no forks, returns nil.
// If there is only one possible fork for the opponent, the player should block it.
// Otherwise, the player should block all forks in any way that simultaneously allows them to make two in a row.
// Otherwise, the player should make a two in a row to force the opponent into defending, as long as it does not result in them producing a fork.
func findCellToBlockPossibleOpponentForks(b board.Board) *board.Cell {
	boardSize := len(b)
	count := b.FullCellsCount()
	if count > boardSize*boardSize-3 { // it is already end part of the game, so the opponent can't create a fork
		return nil
	}
	opponentCanCreateFork := false
	var winCell, possibleCell *board.Cell
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if !b[i][j].IsEmpty() {
				continue
			}
			cell := board.Cell{RowNumber: i, ColumnNumber: j}
			opponentBoard := b.MustSetCellValue(cell)
			opponentForkCell := findPossibleFork(opponentBoard)
			if opponentForkCell != nil {
				opponentCanCreateFork = true
			} else { // opponent can't create a fork, remember this cell
				possibleCell = &cell
			}
			if winCell == nil {
				// try to find a cell that will create a win situation(2 in line) for current player
				// and force opponent to block it
				// and opponent's turn will not create a fork against current player
				c, _ := boardhelper.FindWinSituationsFor(opponentBoard, b.CurrentTurnCellValue())
				if c == nil {
					continue
				}
				nextCurrentPlayerBoard := opponentBoard.MustSetCellValue(*c)
				if _, count := boardhelper.FindWinSituationsFor(
					nextCurrentPlayerBoard,
					b.OpponentCellValue(),
				); count < 2 { // opponent can't create a fork, remember this cell
					winCell = &cell
				}
			}

		}
	}
	if !opponentCanCreateFork {
		return nil
	}
	// Opponent can create fork, we should block it
	if winCell != nil {
		return winCell
	}
	if possibleCell != nil {
		return possibleCell
	}
	panic("opponent can create fork but we can't block it. It is unpredictable situation, fix the code")
}
