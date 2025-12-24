package rockpaperscissors

import (
	"github.com/alex-whitney/tictactoe/engine"
)

type Game struct {
	firstTo int
}

func NewGame(firstTo int) engine.Game {
	return &Game{
		firstTo,
	}
}

func (t *Game) Initialize() engine.GameState {
	return NewState()
}

func (t *Game) CurrentPlayer(state engine.GameState) int {
	s := state.(*State)

	if len(s.Player1History) == len(s.Player2History) {
		return 0
	}

	return 1
}

func (g *Game) IsFinished(state engine.GameState) (bool, int) {
	s := state.(*State)

	if s.PlayerWins[0] >= g.firstTo {
		return true, 0
	} else if s.PlayerWins[1] >= g.firstTo {
		return true, 1
	}

	return false, 0
}

func (g *Game) ValidActions(state engine.GameState) []engine.Action {
	return []engine.Action{
		NewAction(ShapePaper),
		NewAction(ShapeRock),
		NewAction(ShapeScissors),
	}
}

func determineWinner(player1 ShapeType, player2 ShapeType) int {
	if player1 == player2 {
		return -1
	}

	if (player1 == ShapeRock && player2 == ShapeScissors) ||
		(player1 == ShapeScissors && player2 == ShapePaper) ||
		(player1 == ShapePaper && player2 == ShapeRock) {
		return 0
	}

	return 1
}

func (g *Game) ApplyAction(state engine.GameState, player int, action engine.Action) (engine.GameState, error) {
	newState := state.DeepClone().(*State)
	a := action.(*Action)

	if player == 0 {
		newState.Player1History = append(newState.Player1History, a.shape)
	} else {
		newState.Player2History = append(newState.Player2History, a.shape)

		idx := len(newState.Player2History) - 1
		winner := determineWinner(newState.Player1History[idx], a.shape)

		if winner >= 0 {
			newState.PlayerWins[winner]++
		}
	}

	return newState, nil
}
