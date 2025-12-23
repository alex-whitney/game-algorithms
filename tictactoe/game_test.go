package tictactoe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGameInvalidMove(t *testing.T) {
	game := NewGame()
	state := game.Initialize()

	_, err := game.ApplyAction(state, 1, NewAction(0))
	require.Error(t, err)

	state, err = game.ApplyAction(state, 0, NewAction(0))
	require.NoError(t, err)

	_, err = game.ApplyAction(state, 1, NewAction(0))
	require.Error(t, err)
}

func TestWinner(t *testing.T) {
	game := NewGame()
	state := game.Initialize()

	state, _ = game.ApplyAction(state, 0, NewAction(0))
	state, _ = game.ApplyAction(state, 1, NewAction(4))
	state, _ = game.ApplyAction(state, 0, NewAction(1))
	state, _ = game.ApplyAction(state, 1, NewAction(5))
	state, _ = game.ApplyAction(state, 0, NewAction(2))

	completed, winner := game.IsFinished(state)
	assert.True(t, completed)
	assert.Equal(t, 0, winner)
}
