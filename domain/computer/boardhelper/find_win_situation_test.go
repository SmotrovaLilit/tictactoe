package boardhelper

import (
	"github.com/stretchr/testify/require"
	"tictactoe/domain/board"

	"fmt"
	"testing"
)

func TestFindWinSituationFor(t *testing.T) {
	tests := []struct {
		name        string
		b           board.Board
		value       board.CellValue
		wantWinCell board.Cell
		wantCount   int
	}{
		{
			name: "For X turn: when Board don't have 2 'X in lines should return nil and 0",
			b: board.Board{
				{board.EmptyValue, board.OValue, board.OValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
			},
			value:     board.XValue,
			wantCount: 0,
		},
		{
			name: "For X turn: when Board has 2 'X' in column should return the empty in this line cell and 1",
			b: board.Board{
				{board.XValue, board.OValue, board.OValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.XValue, board.EmptyValue, board.EmptyValue},
			},
			value:       board.XValue,
			wantWinCell: board.MustNewCell(1, 0),
			wantCount:   1,
		},
		{
			name: "For X turn: when Board has 2 'X' in row should return the empty in this line cell and 1",
			b: board.Board{
				{board.EmptyValue, board.OValue, board.OValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.XValue, board.XValue, board.EmptyValue},
			},
			value:       board.XValue,
			wantWinCell: board.MustNewCell(2, 2),
			wantCount:   1,
		},
		{
			name: "For X turn: when Board has 2 'X' in diagonal should return the empty in this line cell and 1",
			b: board.Board{
				{board.OValue, board.OValue, board.EmptyValue},
				{board.EmptyValue, board.XValue, board.EmptyValue},
				{board.XValue, board.EmptyValue, board.EmptyValue},
			},
			value:       board.XValue,
			wantWinCell: board.MustNewCell(0, 2),
			wantCount:   1,
		},
		{
			name: "For O turn: when Board has 2 'O' in column should return the empty in this line cell and 1",
			b: board.Board{
				{board.OValue, board.XValue, board.XValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.OValue, board.EmptyValue, board.EmptyValue},
			},
			value:       board.OValue,
			wantWinCell: board.MustNewCell(1, 0),
			wantCount:   1,
		},
		{
			name: "For O turn: when Board has 2 'O' in row should return the empty in this line cell and 1",
			b: board.Board{
				{board.EmptyValue, board.XValue, board.XValue},
				{board.EmptyValue, board.EmptyValue, board.EmptyValue},
				{board.OValue, board.OValue, board.EmptyValue},
			},
			value:       board.OValue,
			wantWinCell: board.MustNewCell(2, 2),
			wantCount:   1,
		},
		{
			name: "For O turn: when Board has 2 'O' in diagonal should return the empty in this line cell and 1",
			b: board.Board{
				{board.OValue, board.XValue, board.XValue},
				{board.EmptyValue, board.OValue, board.XValue},
				{board.XValue, board.EmptyValue, board.EmptyValue},
			},
			value:       board.OValue,
			wantWinCell: board.MustNewCell(2, 2),
			wantCount:   1,
		},
		{
			name: "For O turn: when Board has 2 'O' in 2 lines should return the empty cell in the any found line with 2 'O' and 2",
			b: board.Board{
				{board.OValue, board.XValue, board.XValue},
				{board.EmptyValue, board.OValue, board.XValue},
				{board.OValue, board.XValue, board.EmptyValue},
			},
			value:       board.OValue,
			wantWinCell: board.MustNewCell(2, 2),
			wantCount:   2,
		},
		{
			name: "For X turn: when Board has 2 'X' in 2 lines should return the empty cell in the any found line with 2 'X' and 2",
			b: board.Board{
				{board.OValue, board.EmptyValue, board.EmptyValue},
				{board.OValue, board.XValue, board.EmptyValue},
				{board.XValue, board.EmptyValue, board.XValue},
			},
			value:       board.XValue,
			wantWinCell: board.MustNewCell(0, 2),
			wantCount:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWinCell, gotCount := FindWinSituationsFor(tt.b, tt.value)
			if gotCount > 0 {
				require.Equal(t, &tt.wantWinCell, gotWinCell, fmt.Sprintf("\nboardhelper:\n%v", tt.b))
			}
			require.Equal(t, tt.wantCount, gotCount, fmt.Sprintf("\nboardhelper:\n%v", tt.b))
		})
	}
}
