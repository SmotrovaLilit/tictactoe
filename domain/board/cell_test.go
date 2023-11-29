package board

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCell(t *testing.T) {
	type args struct {
		rowN  int
		collN int
	}
	tests := []struct {
		name    string
		args    args
		want    Cell
		wantErr error
	}{
		{
			name:    "small rowN",
			args:    args{rowN: -1, collN: 0},
			want:    Cell{},
			wantErr: ErrInvalidCell,
		},
		{
			name:    "small collN",
			args:    args{rowN: 0, collN: -1},
			want:    Cell{},
			wantErr: ErrInvalidCell,
		},
		{
			name:    "big rowN",
			args:    args{rowN: boardSize, collN: 0},
			want:    Cell{},
			wantErr: ErrInvalidCell,
		},
		{
			name:    "big collN",
			args:    args{rowN: 0, collN: boardSize},
			want:    Cell{},
			wantErr: ErrInvalidCell,
		},
		{
			name:    "valid",
			args:    args{rowN: 0, collN: 0},
			want:    Cell{RowNumber: 0, ColumnNumber: 0},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCell(tt.args.rowN, tt.args.collN)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMustNewCell(t *testing.T) {
	t.Run("when row and col are correct should return cell", func(t *testing.T) {
		require.NotPanics(t, func() {
			c := MustNewCell(1, 1)
			require.Equal(t, Cell{RowNumber: 1, ColumnNumber: 1}, c)
		})
	})
	t.Run("when row or col are incorrect should panic", func(t *testing.T) {
		require.Panics(t, func() {
			MustNewCell(-1, 1)
		})
	})
}
