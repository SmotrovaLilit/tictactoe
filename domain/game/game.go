package game

import (
	"fmt"
	"tictactoe/domain/board"
)

// Player represents a tic-tac-toe player.
type Player interface {
	Name() string
}

// Game represents a tic-tac-toe game.
// Game is responsible for managing the game state and
// the matching between players and board's cell values.
type Game struct {
	board             board.Board
	player1           Player
	player2           Player
	cellValuesPlayers map[board.CellValue]Player
}

// New creates a new game with the given players.
// The first player is X and the second player is O.
// The first player always starts the game.
func New(player1, player2 Player) *Game {
	g := &Game{
		player1: player1,
		player2: player2,
		cellValuesPlayers: map[board.CellValue]Player{
			board.XValue: player1,
			board.OValue: player2,
		},
	}
	return g
}

// Play plays the given cell.
func (g *Game) Play(cell board.Cell) error {
	b, err := g.board.SetCellValue(cell)
	if err != nil {
		return err
	}
	g.board = b
	return nil
}

// MustPlay is like Play but panics if the cell is invalid.
func (g *Game) MustPlay(cell board.Cell) {
	err := g.Play(cell)
	if err != nil {
		panic(err)
	}
}

// IsOver returns true if the game is over.
// The game is over when the board is full or there is a winner.
func (g *Game) IsOver() bool {
	return g.board.IsCompleted()
}

// GetBoard returns the board.
func (g *Game) GetBoard() board.Board {
	return g.board
}

// Winner returns the winner.
func (g *Game) Winner() Player {
	if winnCellValue, exist := g.board.Winner(); exist {
		p := g.cellValuesPlayers[winnCellValue]
		return p
	}
	return nil
}

// CurrentTurnPlayer returns the current turn player.
func (g *Game) CurrentTurnPlayer() Player {
	return g.cellValuesPlayers[g.board.CurrentTurnCellValue()]
}

// Sprint returns the string representation of the current state of a game.
func (g *Game) Sprint(cursor *board.Cell) string {
	if g.IsOver() {
		s := g.GetBoard().String() + "\n"

		w := g.Winner()
		if w == nil {
			return s + "\nGame is over, Draw"
		}
		return s + fmt.Sprintf("\nGame is over, winner is: %s\n", w)
	}

	result := g.GetBoard().Sprint(cursor) + "\n"
	result += fmt.Sprintf("Current player: %s", g.CurrentTurnPlayer().Name())
	return result
}
