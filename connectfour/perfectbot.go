package connectfour

import (
	"slices"

	"github.com/alex-whitney/tictactoe/engine"
)

func scoringFunc(game engine.Game, s engine.GameState, playerNumber int, opponentNumber int) float64 {
	state := s.(*State)

	isOver, winner := game.IsFinished(state)
	if isOver {
		if winner == playerNumber {
			return 1000
		} else if winner == opponentNumber {
			return -1000
		}

		return 0
	}

	// +1 for a potential winning line, -1 for a potential losing line
	// don't bother with which lines are actually possible
	score := 0

	// Columns
	for row := range numRows - 3 {
		for col := range numCols {
			vals := []int{
				state.Board[row][col],
				state.Board[row+1][col],
				state.Board[row+2][col],
				state.Board[row+3][col],
			}

			if !slices.Contains(vals, playerNumber) {
				score--
			}
			if !slices.Contains(vals, opponentNumber) {
				score++
			}
		}
	}

	// Rows
	for row := range numRows {
		for col := range numCols - 3 {
			vals := []int{
				state.Board[row][col],
				state.Board[row][col+1],
				state.Board[row][col+2],
				state.Board[row][col+3],
			}

			if !slices.Contains(vals, playerNumber) {
				score--
			}
			if !slices.Contains(vals, opponentNumber) {
				score++
			}
		}
	}

	// Diagonals - right
	for row := range numRows - 3 {
		for col := range numCols - 3 {
			vals := []int{
				state.Board[row][col],
				state.Board[row+1][col+1],
				state.Board[row+2][col+2],
				state.Board[row+3][col+3],
			}

			if !slices.Contains(vals, playerNumber) {
				score--
			}
			if !slices.Contains(vals, opponentNumber) {
				score++
			}
		}
	}

	// Diagonals - left
	for row := 3; row < numRows; row++ {
		for col := range numCols - 3 {
			vals := []int{
				state.Board[row][col],
				state.Board[row-1][col+1],
				state.Board[row-2][col+2],
				state.Board[row-3][col+3],
			}

			if !slices.Contains(vals, playerNumber) {
				score--
			}
			if !slices.Contains(vals, opponentNumber) {
				score++
			}
		}
	}

	return float64(score)
}

// Creates a player that chooses an action at random
func NewPerfectBot(maxDepth int) engine.Player {
	return engine.NewMinimaxPlayer(NewGame(), maxDepth, scoringFunc)
}
