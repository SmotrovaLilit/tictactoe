package player

import (
	"errors"
	"strings"
)

// ErrInvalidPlayer is returned when a player name is invalid.
var ErrInvalidPlayer = errors.New("invalid player name")

// Player represents a player.
type Player struct {
	name string
}

// New creates a new player with the given name.
func New(name string) (Player, error) {
	name = strings.Trim(name, " ")
	if name == "" {
		return Player{}, ErrInvalidPlayer
	}
	return Player{
		name: name,
	}, nil
}

// MustNew is like New but panics if name is invalid.
func MustNew(name string) Player {
	p, err := New(name)
	if err != nil {
		panic(err)
	}
	return p
}

// Name returns the player name.
func (p Player) Name() string {
	return p.name
}

// String returns the string representation of the player.
func (p Player) String() string {
	return p.Name()
}
