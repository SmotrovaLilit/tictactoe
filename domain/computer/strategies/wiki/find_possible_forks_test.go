package wiki

import (
	"github.com/stretchr/testify/require"
	"tictactoe/domain/board"

	"fmt"
	"testing"
)

func TestFindPossibleFork(t *testing.T) {
	tests := []struct {
		name string
		b    board.Board
		want *board.Cell
	}{
		{
			name: "1. when Board don't have possible fork should return nil",
			b: board.Board{
				{board.OValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			want: nil,
		},
		{
			name: "2. when Board don't have possible fork should return nil",
			b: board.Board{
				{board.OValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.XValue},
			},
			want: nil,
		},
		{
			name: "1. when Board has possible fork should return cell",
			b: board.Board{
				{board.OValue, board.OValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.XValue},
			},
			want: &board.Cell{
				RowNumber:    0,
				ColumnNumber: 2,
			},
		},
		{
			name: "2. when Board has possible fork should return cell",
			b: board.Board{
				{board.OValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.OValue},
				{board.EmptyValue, board.EmptyValue, board.XValue},
			},
			want: &board.Cell{
				RowNumber:    2,
				ColumnNumber: 0,
			},
		},
		{
			name: "3. when Board has possible fork should return cell",
			b: board.Board{
				{board.OValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.EmptyValue, board.OValue, board.XValue},
			},
			want: &board.Cell{
				RowNumber:    0,
				ColumnNumber: 2,
			},
		},
		{
			name: "4. when Board has possible fork should return cell",
			b: board.Board{
				{board.OValue, board.XValue, board.XValue},
				{board.EmptyValue, board.OValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.XValue},
			},
			want: &board.Cell{
				RowNumber:    1,
				ColumnNumber: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPossibleFork(tt.b)
			require.Equal(t, tt.want, got, fmt.Sprintf("\nboardhelper:\n%v", tt.b))
		})
	}
}
