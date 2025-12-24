package engine

// The game state for a game.
//
// This interface does not make any internals of the state visible, but it does
// allow for producing a deep copy of a state.
type GameState interface {

	// Produces a string representation of the current state
	String() string

	// Generate a full copy of this game state
	DeepClone() GameState

	// Generate a full copy of this game state that represents what is visible
	// for the specified player
	VisibleForPlayer(player int) GameState
}
