package engine

// The rules engine for a particular game
type Game interface {
	// Initializes a new GameState for this game
	Initialize() GameState

	// Returns the current player number
	CurrentPlayer(state GameState) int

	// Returns a list of valid actions for the active player, given a game state
	ValidActions(state GameState) []Action

	// Applies a particular action by the specified player to the game state,
	// returning the new game state. The original state is not mutated.
	ApplyAction(state GameState, player int, action Action) (GameState, error)

	// Returns whether or not the game is finished, as well as the player number
	// of the winner. A winner of -1 indicates a tie.
	IsFinished(state GameState) (gameIsOver bool, winner int)
}
