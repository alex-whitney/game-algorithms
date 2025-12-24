package tictactoe

import (
	"fmt"
	"slices"

	"github.com/alex-whitney/tictactoe/engine"
)

type State struct {
	Board []BoardMark

	CurrentPlayer int
	MarkCount     int
}

func NewState() engine.GameState {
	b := &State{}
	b.Board = make([]BoardMark, 9)

	return b
}

func (g *State) String() string {
	s := map[BoardMark]string{
		MarkAvailable: ".",
		MarkX:         "X",
		MarkO:         "O",
	}

	ret := "\t    0 1 2  \n"
	ret += "\t  ---------\n"
	ret += fmt.Sprintf("\t0 | %s %s %s |\n", s[g.Board[0]], s[g.Board[1]], s[g.Board[2]])
	ret += fmt.Sprintf("\t1 | %s %s %s |\n", s[g.Board[3]], s[g.Board[4]], s[g.Board[5]])
	ret += fmt.Sprintf("\t2 | %s %s %s |\n", s[g.Board[6]], s[g.Board[7]], s[g.Board[8]])
	ret += "\t  ---------\n"

	return ret
}

func (g *State) DeepClone() engine.GameState {
	ret := &State{
		Board:         slices.Clone(g.Board),
		CurrentPlayer: g.CurrentPlayer,
		MarkCount:     g.MarkCount,
	}

	return ret
}

func (g *State) VisibleForPlayer(_ int) engine.GameState {
	return g.DeepClone()
}
