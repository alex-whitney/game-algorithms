package rockpaperscissors

import (
	"fmt"
	"slices"

	"github.com/alex-whitney/tictactoe/engine"
	"github.com/thomaso-mirodin/intmath/intgr"
)

type State struct {
	Player1History []ShapeType
	Player2History []ShapeType
	PlayerWins     []int
}

func NewState() engine.GameState {
	b := &State{
		Player1History: make([]ShapeType, 0),
		Player2History: make([]ShapeType, 0),
		PlayerWins:     []int{0, 0},
	}

	return b
}

func (g *State) String() string {
	lastRound := intgr.Min(len(g.Player1History), len(g.Player2History))

	ret := ""

	if lastRound > 0 {
		winner := determineWinner(g.Player1History[lastRound-1], g.Player2History[lastRound-1])

		ret += fmt.Sprintf("Last round:\n\tPlayer 1 chose %s\n\tPlayer 2 chose %s\n",
			shapeName[g.Player1History[lastRound-1]],
			shapeName[g.Player2History[lastRound-1]])
		if winner < 0 {
			ret += "Result: TIE!\n"
		} else {
			ret += fmt.Sprintf("Result: Player %d wins!\n", winner+1)
		}
	}

	ret += fmt.Sprintf("\nPlayer 1: %d wins\nPlayer 2: %d wins\n", g.PlayerWins[0], g.PlayerWins[1])

	return ret
}

func (g *State) DeepClone() engine.GameState {
	return &State{
		Player1History: slices.Clone(g.Player1History),
		Player2History: slices.Clone(g.Player2History),
		PlayerWins:     slices.Clone(g.PlayerWins),
	}
}

func (g *State) VisibleForPlayer(_ int) engine.GameState {
	return g.DeepClone()
}
