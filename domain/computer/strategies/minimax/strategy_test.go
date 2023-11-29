package minimax

import (
	"github.com/stretchr/testify/require"
	"tictactoe/domain/board"

	"fmt"
	"testing"
)

func TestStrategy_findBestCellForNextTurn(t *testing.T) {
	tests := []struct {
		name  string
		board board.Board
		want  []int
	}{
		{
			name: "2. At third turn should return corner cell without win situation because it is too predictable",
			board: board.Board{
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.OValue},
			},
			want: []int{0, 0},
		},
		{
			name:  "At first turn should return corner cell",
			board: board.Board{},
			want:  []int{0, 0},
		},
		{
			name: "At second turn if center cell is empty should return center cell",
			board: board.Board{
				{board.XValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			want: []int{1, 1},
		},
		{
			name: "At second turn if center cell is not empty should return corner cell",
			board: board.Board{
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			want: []int{0, 0},
		},
		{
			name: "1. At 4th turn if opponent has win situation should return cell to block it",
			board: board.Board{
				{board.XValue, board.EmptyValue, board.OValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			want: []int{2, 2},
		},
		{
			name: "2. At 4th turn if opponent has win situation should return cell to block it",
			board: board.Board{
				{board.EmptyValue, board.EmptyValue, board.OValue},
				{board.EmptyValue, board.XValue, board.XValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			want: []int{1, 0},
		},
		{
			name: "3. At 4th turn if opponent has win situation should return cell to block it",
			board: board.Board{
				{board.XValue, board.EmptyValue, board.EmptyValue},
				{board.XValue, board.OValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			want: []int{2, 0},
		},
		{
			name: "4. At 4th turn if opponent has win situation even if center is empty should return cell to block it",
			board: board.Board{
				{board.XValue, board.EmptyValue, board.EmptyValue},
				{board.XValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.OValue},
			},
			want: []int{2, 0},
		},
		{
			name: "At 4th turn if opponent doesn't have win situation and center is empty and center can allow opponent to create a fork shouldn't return center cell",
			board: board.Board{
				{board.XValue, board.EmptyValue, board.OValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
			},
			want: []int{2, 2},
		},
		{
			name: "1. At 5th should create fork if possible",
			board: board.Board{
				{board.OValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.OValue},
				{board.EmptyValue, board.EmptyValue, board.XValue},
			},
			want: []int{2, 0},
		},
		{
			name: "2. At 5th should create fork if possible",
			board: board.Board{
				{board.OValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.OValue, board.XValue},
			},
			want: []int{0, 2},
		},
		{
			name: "Avoid fork",
			board: board.Board{
				{board.XValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.OValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.XValue},
			},
			want: []int{0, 1},
		},
		{
			name: "Last turns",
			board: board.Board{
				{board.XValue, board.OValue, board.XValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.OValue, board.XValue, board.OValue},
			},
			want: []int{1, 0},
		},
		{
			name: "If the opponent has two in a row, the player must play the third themselves to block the opponent.",
			board: board.Board{
				{board.OValue, board.EmptyValue, board.XValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			want: []int{2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := NewStrategy()
			got := str.FindBestCellForNextTurn(tt.board)

			expectedBoard := tt.board.MustSetCellValue(board.Cell{
				RowNumber:    tt.want[0],
				ColumnNumber: tt.want[1],
			})
			actualBoard := tt.board.MustSetCellValue(got)
			require.Equal(t, tt.want, []int{got.RowNumber, got.ColumnNumber},
				fmt.Sprintf("expected boardhelper:\n%vactual boardhelper:\n%v", expectedBoard, actualBoard),
			)
		})
	}
}
