package player

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Player
		wantErr error
	}{
		{
			name:    "Empty name",
			args:    args{name: ""},
			want:    Player{},
			wantErr: ErrInvalidPlayer,
		},
		{
			name:    "Blank name",
			args:    args{name: "   "},
			want:    Player{},
			wantErr: ErrInvalidPlayer,
		},
		{
			name:    "Valid name",
			args:    args{name: "John"},
			want:    Player{name: "John"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.name)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMustNewPlayer(t *testing.T) {
	t.Run("When name is valid should return player", func(t *testing.T) {
		got := MustNew("John")
		require.Equal(t, Player{name: "John"}, got)
	})
	t.Run("When name is invalid should panic", func(t *testing.T) {
		require.Panics(t, func() {
			MustNew("")
		})
	})
}
