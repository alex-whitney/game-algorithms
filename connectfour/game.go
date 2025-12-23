package connectfour

import (
	"errors"
	"fmt"
	"slices"

	"github.com/alex-whitney/tictactoe/engine"
)

type Game struct {
}

func NewGame() engine.Game {
	return &Game{}
}

func (t *Game) Initialize() engine.GameState {
	return NewState()
}

func (t *Game) CurrentPlayer(state engine.GameState) int {
	s := state.(*State)

	return s.TurnCounter % 2
}

func (g *Game) IsFinished(state engine.GameState) (bool, int) {
	s := state.(*State)

	// Columns
	for row := range numRows - 3 {
		for col := range numCols {
			p := s.Board[row][col]
			if p == 0 {
				continue
			}

			if p == s.Board[row+1][col] &&
				p == s.Board[row+2][col] &&
				p == s.Board[row+3][col] {
				return true, p - 1
			}
		}
	}

	// Rows
	for row := range numRows {
		for col := range numCols - 3 {
			p := s.Board[row][col]
			if p == 0 {
				continue
			}

			if p == s.Board[row][col+1] &&
				p == s.Board[row][col+2] &&
				p == s.Board[row][col+3] {
				return true, p - 1
			}
		}
	}

	// Diagonals - right
	for row := range numRows - 3 {
		for col := range numCols - 3 {
			p := s.Board[row][col]
			if p == 0 {
				continue
			}

			if p == s.Board[row+1][col+1] &&
				p == s.Board[row+2][col+2] &&
				p == s.Board[row+3][col+3] {
				return true, p - 1
			}
		}
	}

	// Diagonals - left
	for row := 3; row < numRows; row++ {
		for col := range numCols - 3 {
			p := s.Board[row][col]
			if p == 0 {
				continue
			}

			if p == s.Board[row-1][col+1] &&
				p == s.Board[row-2][col+2] &&
				p == s.Board[row-3][col+3] {
				return true, p - 1
			}
		}
	}

	// Tie
	if s.TurnCounter == numCols*numRows {
		return true, -1
	}

	return false, 0
}

func (g *Game) ValidActions(state engine.GameState) []engine.Action {
	s := state.(*State)

	ret := []engine.Action{}

	// If the top position of each column is empty, that column can be played
	for col, val := range s.Board[numRows-1] {
		if val == 0 {
			ret = append(ret, NewAction(col))
		}
	}

	return ret
}

func (g *Game) ApplyAction(state engine.GameState, player int, action engine.Action) (engine.GameState, error) {
	newState := state.DeepClone().(*State)

	if g.CurrentPlayer(state) != player {
		return nil, fmt.Errorf("player %d is not current", player)
	}

	validActionIds := engine.CollectIds(g.ValidActions(state))
	if !slices.Contains(validActionIds, action.Id()) {
		return nil, errors.New("action is not valid")
	}

	topRow := numRows
	for r := topRow - 1; r >= 0; r-- {
		if newState.Board[r][action.Id()] > 0 {
			break
		}
		topRow = r
	}

	newState.Board[topRow][action.Id()] = player + 1
	newState.TurnCounter++

	return newState, nil
}
