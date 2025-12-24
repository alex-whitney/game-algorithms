package engine

// Represents a distinct action in a game
type Action interface {
	// Returns a game-specific id for this specific action. This Id has no
	// meaning for the game engine
	Id() int

	// A string description of this action
	Description() string

	// Whether or not this action is hidden from other players
	IsHidden() bool
}

func CollectIds(actions []Action) []int {
	ret := make([]int, len(actions))
	for idx, v := range actions {
		ret[idx] = v.Id()
	}
	return ret
}
