package tictactoe

import (
	"errors"
	"fmt"
	"slices"

	"github.com/alex-whitney/tictactoe/engine"
)

type BoardMark int

const (
	MarkAvailable BoardMark = iota
	MarkX
	MarkO
)

var markToPlayerNumber = map[BoardMark]int{
	MarkX: 0,
	MarkO: 1,
}
var playerNumberToMark = map[int]BoardMark{
	0: MarkX,
	1: MarkO,
}

type Game struct {
}

func NewGame() engine.Game {
	return &Game{}
}

func (t *Game) Initialize() engine.GameState {
	return NewState()
}

func (t *Game) CurrentPlayer(state engine.GameState) int {
	tttState := state.(*State)

	return tttState.CurrentPlayer
}

var potentialWinningCombinations = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

func (g *Game) IsFinished(state engine.GameState) (bool, int) {
	tttState := state.(*State)

	if tttState.MarkCount >= 9 {
		return true, -1
	}

	for _, p := range potentialWinningCombinations {
		position := tttState.Board[p[0]]
		if position != MarkAvailable &&
			position == tttState.Board[p[1]] &&
			position == tttState.Board[p[2]] {
			return true, markToPlayerNumber[position]
		}
	}

	return false, 0
}

func (g *Game) ValidActions(state engine.GameState) []engine.Action {
	tttState := state.(*State)
	ret := []engine.Action{}

	if hasWinner, _ := g.IsFinished(state); hasWinner {
		return ret
	}

	for i, val := range tttState.Board {
		if val == MarkAvailable {
			ret = append(ret, NewAction(i))
		}
	}

	return ret
}

func (g *Game) ApplyAction(state engine.GameState, player int, action engine.Action) (engine.GameState, error) {
	newState := state.DeepClone().(*State)

	if newState.CurrentPlayer != player {
		return nil, fmt.Errorf("player %d is not current", player)
	}

	validActionIds := engine.CollectIds(g.ValidActions(state))
	if !slices.Contains(validActionIds, action.Id()) {
		return nil, errors.New("action is not valid")
	}

	newState.Board[action.Id()] = playerNumberToMark[newState.CurrentPlayer]
	newState.CurrentPlayer = (newState.CurrentPlayer + 1) % 2

	newState.MarkCount++

	return newState, nil
}
