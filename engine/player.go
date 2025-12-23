package engine

import "math/rand"

// A player
type Player interface {
	// Returns the player's name
	Name() string

	// Given a list of potentially valid actions and a game state, chooses an
	// action
	ChooseAction(state GameState, validActions []Action) Action
}

type randomPlayer struct {
	name string
}

// Creates a player that chooses an action at random
func NewRandomPlayer(name string) Player {
	return &randomPlayer{
		name,
	}
}

func (p *randomPlayer) Name() string {
	return p.name
}

func (p *randomPlayer) ChooseAction(state GameState, validActions []Action) Action {
	if len(validActions) == 1 {
		return validActions[0]
	}

	return validActions[rand.Intn(len(validActions))]
}
