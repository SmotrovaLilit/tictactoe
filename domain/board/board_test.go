package board

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_board_isEmptyCell(t *testing.T) {
	type args struct {
		cell Cell
	}
	tests := []struct {
		name string
		b    Board
		args args
		want bool
	}{
		{
			name: "when cell is empty should return true",
			b: Board{
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, XValue, EmptyValue},
			},
			args: args{cell: Cell{RowNumber: 0, ColumnNumber: 2}},
			want: true,
		},
		{
			name: "when cell is not empty should return false",
			b: Board{
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, XValue, EmptyValue},
			},
			args: args{cell: Cell{RowNumber: 2, ColumnNumber: 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.IsEmptyCell(tt.args.cell)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_SetCellValue(t *testing.T) {
	tests := []struct {
		name    string
		b       Board
		cell    Cell
		want    Board
		wantErr error
	}{
		{
			name: "when cell is not empty should return error",
			cell: MustNewCell(0, 0),
			b: Board{
				{XValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
			wantErr: ErrCellIsNotEmpty,
			want:    Board{},
		},
		{
			name: "when games is over should return error",
			cell: MustNewCell(0, 2),
			b: Board{
				{XValue, EmptyValue, EmptyValue},
				{OValue, XValue, OValue},
				{OValue, EmptyValue, XValue},
			},
			wantErr: ErrGameIsOver,
			want:    Board{},
		},
		{
			name: "when 0 turn should update Board and set 0 value for the cell",
			cell: MustNewCell(0, 0),
			b: Board{
				{EmptyValue, OValue, XValue},
				{EmptyValue, XValue, EmptyValue},
				{OValue, XValue, OValue},
			},
			wantErr: nil,
			want: Board{
				{XValue, OValue, XValue},
				{EmptyValue, XValue, EmptyValue},
				{OValue, XValue, OValue},
			},
		},
		{
			name: "when X turn should update Board and set X value for the cell",
			cell: MustNewCell(0, 0),
			b: Board{
				{EmptyValue, OValue, EmptyValue},
				{XValue, XValue, OValue},
				{OValue, XValue, EmptyValue},
			},
			wantErr: nil,
			want: Board{
				{XValue, OValue, EmptyValue},
				{XValue, XValue, OValue},
				{OValue, XValue, EmptyValue},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.SetCellValue(tt.cell)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_Winner(t *testing.T) {
	tests := []struct {
		name          string
		b             Board
		wantCellValue CellValue
		wantExist     bool
	}{
		{
			name: "when Board has only 2 'X' or 2 'O' in row  should return empty CellValue and false",
			b: Board{
				{OValue, OValue, EmptyValue},
				{XValue, XValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
		},
		{
			name: "when Board has only 2 '0' in column or 2 'X' in row should return empty CellValue and false",
			b: Board{
				{EmptyValue, OValue, EmptyValue},
				{XValue, EmptyValue, XValue},
				{EmptyValue, OValue, EmptyValue},
			},
		},
		{
			name: "when Board has only 2 'X' in diagonal should return empty CellValue and false",
			b: Board{
				{XValue, OValue, OValue},
				{EmptyValue, XValue, EmptyValue},
				{OValue, OValue, EmptyValue},
			},
		},
		{
			name: "when Board has only 2 '0' in diagonal  should return empty CellValue and false",
			b: Board{
				{EmptyValue, EmptyValue, XValue},
				{EmptyValue, OValue, EmptyValue},
				{XValue, EmptyValue, OValue},
			},
		},
		{
			name: "when Board has 3 'X' in row  should return 'X' CellValue and true",
			b: Board{
				{EmptyValue, OValue, EmptyValue},
				{XValue, XValue, XValue},
				{EmptyValue, OValue, EmptyValue},
			},
			wantCellValue: XValue,
			wantExist:     true,
		},
		{
			name: "when Board has 3 '0' in column  should return '0' CellValue and true",
			b: Board{
				{EmptyValue, OValue, EmptyValue},
				{XValue, OValue, XValue},
				{EmptyValue, OValue, EmptyValue},
			},
			wantCellValue: OValue,
			wantExist:     true,
		},
		{
			name: "when Board has 3 'X' in diagonal should return 'X' CellValue and true",
			b: Board{
				{XValue, OValue, OValue},
				{EmptyValue, XValue, EmptyValue},
				{OValue, OValue, XValue},
			},
			wantCellValue: XValue,
			wantExist:     true,
		},
		{
			name: "when Board has 3 '0' in diagonal  should return '0' CellValue and true",
			b: Board{
				{XValue, EmptyValue, OValue},
				{EmptyValue, OValue, XValue},
				{OValue, EmptyValue, XValue},
			},
			wantCellValue: OValue,
			wantExist:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ex := tt.b.Winner()
			require.Equal(t, tt.wantCellValue, got)
			require.Equal(t, tt.wantExist, ex)
		})
	}
}

func TestBoard_FullCellsCount(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want int
	}{
		{
			name: "when Board is empty should return 0",
			b: Board{
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
			want: 0,
		},
		{
			name: "when Board is full should return 9",
			b: Board{
				{XValue, OValue, OValue},
				{XValue, OValue, XValue},
				{OValue, XValue, XValue},
			},
			want: 9,
		},
		{
			name: "when Board is almost full should return 8",
			b: Board{
				{XValue, OValue, OValue},
				{XValue, OValue, XValue},
				{EmptyValue, XValue, XValue},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.FullCellsCount()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_IsFull(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want bool
	}{
		{
			name: "when Board is full should return true",
			b: Board{
				{XValue, OValue, OValue},
				{XValue, OValue, XValue},
				{OValue, XValue, XValue},
			},
			want: true,
		},
		{
			name: "when Board is empty should return false",
			b: Board{
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
			want: false,
		},
		{
			name: "when Board is almost full should return false",
			b: Board{
				{XValue, OValue, OValue},
				{XValue, OValue, XValue},
				{EmptyValue, XValue, XValue},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.IsFull()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_MustSetCellValue1(t *testing.T) {
	t.Run("when SetCellValue return an error should panic", func(t *testing.T) {
		require.Panics(t, func() {
			b := Board{
				{XValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			}
			b.MustSetCellValue(MustNewCell(0, 0))
		})
	})
	t.Run("when cell is correct should set new value", func(t *testing.T) {
		b := Board{
			{XValue, EmptyValue, EmptyValue},
			{EmptyValue, EmptyValue, EmptyValue},
			{EmptyValue, EmptyValue, EmptyValue},
		}
		b.MustSetCellValue(MustNewCell(0, 1))
	})
}

func TestBoard_CurrentTurnCellValue(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want CellValue
	}{
		{
			name: "when Board is empty should return X",
			b: Board{
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
			want: XValue,
		},
		{
			name: "when Board has odd played cells should return 0",
			b: Board{
				{XValue, EmptyValue, EmptyValue},
				{EmptyValue, OValue, EmptyValue},
				{EmptyValue, EmptyValue, XValue},
			},
			want: OValue,
		},
		{
			name: "when Board has even played cells should return X",
			b: Board{
				{XValue, EmptyValue, OValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{OValue, EmptyValue, XValue},
			},
			want: XValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.CurrentTurnCellValue()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_IsCompleted(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want bool
	}{
		{
			name: "when Board is not full and there are no winner should return false",
			b: Board{
				{XValue, EmptyValue, OValue},
				{EmptyValue, XValue, OValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
			want: false,
		},
		{
			name: "when Board has winner should return true",
			b: Board{
				{XValue, EmptyValue, OValue},
				{EmptyValue, XValue, OValue},
				{EmptyValue, EmptyValue, XValue},
			},
			want: true,
		},
		{
			name: "when Board is full should return true",
			b: Board{
				{XValue, OValue, XValue},
				{XValue, OValue, XValue},
				{OValue, XValue, OValue},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.IsCompleted()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_MidCell(t *testing.T) {
	t.Run("should return middle cell", func(t *testing.T) {
		b := Board{
			{XValue, EmptyValue, OValue},
			{XValue, OValue, XValue},
			{OValue, EmptyValue, XValue},
		}
		got := b.MidCell()
		want := MustNewCell(1, 1)
		require.Equal(t, want, got)
	})
}

func TestBoard_Corners(t *testing.T) {
	t.Run("should return corners", func(t *testing.T) {
		b := Board{
			{XValue, EmptyValue, OValue},
			{XValue, OValue, XValue},
			{OValue, EmptyValue, XValue},
		}
		got := b.Corners()
		want := []Cell{
			MustNewCell(0, 0),
			MustNewCell(0, 2),
			MustNewCell(2, 0),
			MustNewCell(2, 2),
		}
		require.Equal(t, want, got)
	})
}

func TestBoard_SideCells(t *testing.T) {
	t.Run("should return side cells", func(t *testing.T) {
		b := Board{
			{XValue, EmptyValue, OValue},
			{XValue, OValue, XValue},
			{OValue, EmptyValue, XValue},
		}
		got := b.SideCells()
		want := []Cell{
			MustNewCell(0, 1),
			MustNewCell(1, 0),
			MustNewCell(1, 2),
			MustNewCell(2, 1),
		}
		require.Equal(t, want, got)
	})
}

func TestBoard_CellValue(t *testing.T) {
	t.Run("should return cell value", func(t *testing.T) {
		b := Board{
			{XValue, EmptyValue, OValue},
			{XValue, OValue, XValue},
			{OValue, EmptyValue, XValue},
		}
		require.Equal(t, OValue, b.CellValue(MustNewCell(0, 2)))
		require.Equal(t, XValue, b.CellValue(MustNewCell(1, 2)))
	})
}

func TestBoard_OpponentCellValue(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want CellValue
	}{
		{
			name: "when Board is empty should return O",
			b: Board{
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
			want: OValue,
		},
		{
			name: "when Board has odd played cells should return X",
			b: Board{
				{XValue, EmptyValue, EmptyValue},
				{EmptyValue, OValue, EmptyValue},
				{EmptyValue, EmptyValue, XValue},
			},
			want: XValue,
		},
		{
			name: "when Board has even played cells should return O",
			b: Board{
				{XValue, EmptyValue, OValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{OValue, EmptyValue, XValue},
			},
			want: OValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.OpponentCellValue()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_FindFirstEmptyCell(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want *Cell
	}{
		{
			name: "when Board is empty should return first cell",
			b: Board{
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
			want: &Cell{RowNumber: 0, ColumnNumber: 0},
		},
		{
			name: "when Board's  second cell is empty and first cell is not empty should return the second cell",
			b: Board{
				{XValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
				{EmptyValue, EmptyValue, EmptyValue},
			},
			want: &Cell{RowNumber: 0, ColumnNumber: 1},
		},
		{
			name: "when Board is full should return nil",
			b: Board{
				{XValue, OValue, XValue},
				{XValue, OValue, XValue},
				{OValue, XValue, OValue},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.FindFirstEmptyCell()
			require.Equal(t, tt.want, got)
		})
	}
}
