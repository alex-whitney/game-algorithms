package connectfour

import (
	"fmt"
	"slices"

	"github.com/alex-whitney/tictactoe/engine"
)

const numRows = 6
const numCols = 7

type State struct {
	// value = which player has put a piece in that position
	//  0 = empty
	//  1 = player 1
	//  2 = player 2
	//
	// row 0 = bottom row, row 6 = top row
	Board [][]int

	// A counter of all of the pieces that are on the board
	TurnCounter int
}

func NewState() engine.GameState {
	b := &State{}
	b.Board = make([][]int, numRows)
	for idx := range b.Board {
		b.Board[idx] = make([]int, numCols)
	}

	return b
}

func (s *State) String() string {
	dict := map[int]string{
		0: ".",
		1: "R",
		2: "Y",
	}

	ret := "\t--1-2-3-4-5-6-7--\n"
	ret += "\t-----------------\n"

	// Top to bottom
	for row := numRows - 1; row >= 0; row-- {
		ret += "\t| "

		for col := range s.Board[row] {
			ret += fmt.Sprintf("%s ", dict[s.Board[row][col]])
		}

		ret += "|\n"
	}

	ret += "\t-----------------\n"

	return ret
}

func (s *State) DeepClone() engine.GameState {
	ret := &State{
		TurnCounter: s.TurnCounter,
		Board:       make([][]int, numRows),
	}

	for row := range ret.Board {
		ret.Board[row] = slices.Clone(s.Board[row])
	}

	return ret
}
