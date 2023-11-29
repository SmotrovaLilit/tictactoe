package game

import (
	"github.com/stretchr/testify/require"
	"testing"
	"tictactoe/domain/board"
	"tictactoe/domain/player"
)

func TestGame_Play(t *testing.T) {
	tests := []struct {
		name    string
		board   board.Board
		cell    board.Cell
		wantErr bool
	}{
		{
			name: "when cell is not empty or board is completed  should return error",
			board: board.Board{
				{board.XValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			cell:    board.MustNewCell(0, 0),
			wantErr: true,
		},
		{
			name: "when cell is empty and board is not completed should update Board and set X value for the cell",
			board: board.Board{
				{board.EmptyValue, board.OValue, board.XValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.OValue, board.XValue, board.OValue},
			},
			cell:    board.MustNewCell(0, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(player.MustNew("John"), player.MustNew("Jane"))
			g.board = tt.board
			err := g.Play(tt.cell)
			require.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				require.Equal(t, board.XValue, g.board.CellValue(tt.cell))
			} else {
				require.Equal(t, tt.board, g.board)
			}
		})
	}
}

func TestGame_IsOver(t *testing.T) {
	tests := []struct {
		name string
		b    board.Board
		want bool
	}{
		{
			name: "when board is not completed should return false",
			b: board.Board{
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.XValue},
				{board.EmptyValue, board.EmptyValue, board.OValue},
			},
		},
		{
			name: "when board is completed should return true",
			b: board.Board{
				{board.XValue, board.OValue, board.XValue},
				{board.XValue, board.XValue, board.OValue},
				{board.OValue, board.XValue, board.OValue},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(player.MustNew("John"), player.MustNew("Jane"))
			g.board = tt.b
			got := g.IsOver()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewGame(t *testing.T) {
	p1 := player.MustNew("John")
	p2 := player.MustNew("Jane")
	g := New(p1, p2)
	require.Equal(t, p1, g.player1)
	require.Equal(t, p2, g.player2)
	require.Equal(t, p1, g.cellValuesPlayers[board.XValue])
	require.Equal(t, p2, g.cellValuesPlayers[board.OValue])
	require.Equal(t, board.Board{}, g.board)
}

func TestGame_CurrentTurnPlayer(t *testing.T) {
	t.Run("when game is started should return first player", func(t *testing.T) {
		p1 := player.MustNew("John")
		g := New(p1, player.MustNew("Jane"))
		require.Equal(t, p1, g.CurrentTurnPlayer())
	})
	t.Run("on second turn should return second player", func(t *testing.T) {
		p1 := player.MustNew("John")
		p2 := player.MustNew("Jane")
		g := New(p1, p2)
		g.board = board.Board{
			{board.XValue, board.EmptyValue, board.EmptyValue},
			{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			{board.EmptyValue, board.EmptyValue, board.EmptyValue},
		}
		require.Equal(t, p2, g.CurrentTurnPlayer())
	})
}

func TestGame_Winner(t *testing.T) {
	t.Run("when game is started should return nil", func(t *testing.T) {
		g := New(player.MustNew("John"), player.MustNew("Jane"))
		require.Nil(t, g.Winner())
	})
	t.Run("when game is over in draw should return nil", func(t *testing.T) {
		g := New(player.MustNew("John"), player.MustNew("Jane"))
		g.board = board.Board{
			{board.XValue, board.OValue, board.XValue},
			{board.XValue, board.XValue, board.OValue},
			{board.OValue, board.XValue, board.OValue},
		}
		require.Nil(t, g.Winner())
	})
	t.Run("when game is over with winner should return winner", func(t *testing.T) {
		p1 := player.MustNew("John")
		g := New(p1, player.MustNew("Jane"))
		g.board = board.Board{
			{board.XValue, board.OValue, board.XValue},
			{board.XValue, board.OValue, board.OValue},
			{board.XValue, board.XValue, board.OValue},
		}
		require.Equal(t, p1, g.Winner())
	})
}

func TestGame_MustPlay(t *testing.T) {
	t.Run("when cell is not empty or board is completed  should panic", func(t *testing.T) {
		g := New(player.MustNew("John"), player.MustNew("Jane"))
		g.board = board.Board{
			{board.XValue, board.EmptyValue, board.EmptyValue},
			{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			{board.EmptyValue, board.EmptyValue, board.EmptyValue},
		}
		require.Panics(t, func() {
			g.MustPlay(board.MustNewCell(0, 0))
		})
	})
	t.Run("when cell is empty and board is not completed should update Board and set X value for the cell", func(t *testing.T) {
		g := New(player.MustNew("John"), player.MustNew("Jane"))
		g.board = board.Board{
			{board.EmptyValue, board.OValue, board.XValue},
			{board.EmptyValue, board.XValue, board.EmptyValue},
			{board.OValue, board.XValue, board.OValue},
		}
		g.MustPlay(board.MustNewCell(0, 0))
		require.Equal(t, board.XValue, g.board.CellValue(board.MustNewCell(0, 0)))
	})
}

func TestGame_Sprint(t *testing.T) {
	t.Run("when cursor is nil should return string representation of the current state of a game", func(t *testing.T) {
		g := New(player.MustNew("John"), player.MustNew("Jane"))
		g.board = board.Board{
			{board.EmptyValue, board.OValue, board.XValue},
			{board.EmptyValue, board.XValue, board.EmptyValue},
			{board.OValue, board.XValue, board.OValue},
		}
		want := " -  |  O  |  X \n -  |  X  |  - \n O  |  X  |  O \n\nCurrent player: John"
		require.Equal(t, want, g.Sprint(nil))
	})
	t.Run("when game is over with draw should return string representation of the current state of a game", func(t *testing.T) {
		g := New(player.MustNew("John"), player.MustNew("Jane"))
		g.board = board.Board{
			{board.XValue, board.OValue, board.XValue},
			{board.OValue, board.XValue, board.XValue},
			{board.OValue, board.XValue, board.OValue},
		}
		want := " X  |  O  |  X \n O  |  X  |  X \n O  |  X  |  O \n\n\nGame is over, Draw"
		require.Equal(t, want, g.Sprint(nil))
	})
	t.Run("when game is over with a winner should return string representation of the current state of a game", func(t *testing.T) {
		g := New(player.MustNew("John"), player.MustNew("Jane"))
		g.board = board.Board{
			{board.XValue, board.OValue, board.XValue},
			{board.OValue, board.OValue, board.XValue},
			{board.OValue, board.XValue, board.XValue},
		}
		want := " X  |  O  |  X \n O  |  O  |  X \n O  |  X  |  X \n\n\nGame is over, winner is: John\n"
		require.Equal(t, want, g.Sprint(nil))
	})
}
