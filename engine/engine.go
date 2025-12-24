package engine

import (
	"fmt"
	"slices"
)

type Engine struct {
	players      []Player
	game         Game
	currentState GameState
}

func NewEngine(game Game, player1 Player, player2 Player) *Engine {
	ret := &Engine{
		game: game,
		players: []Player{
			player1,
			player2,
		},
		currentState: game.Initialize(),
	}

	return ret
}

// Runs the game until it reaches its conclusion
//
// In the case of ties, nil is returned for Player. Otherwise, the winner is
// is returned.
func (e *Engine) RunToCompletion() (Player, error) {
	var winner int
	var err error
	for {
		var done bool
		done, winner = e.game.IsFinished(e.currentState)
		if done {
			break
		}

		activePlayer := e.game.CurrentPlayer(e.currentState)
		actions := e.game.ValidActions(e.currentState)

		action := e.players[activePlayer].ChooseAction(e.currentState, actions)
		if !slices.Contains(actions, action) {
			return nil, fmt.Errorf("player %d is trying to cheat - selected action not valid", activePlayer)
		}

		if action.IsHidden() {
			fmt.Printf("%s has selected an action\n", e.players[activePlayer].Name())
		} else {
			fmt.Printf("%s action: %s\n", e.players[activePlayer].Name(), action.Description())
		}

		e.currentState, err = e.game.ApplyAction(e.currentState, activePlayer, action)
		if err != nil {
			return nil, err
		}
	}

	fmt.Printf("\n\nFinal Game State:\n%s\n", e.currentState.String())

	if winner >= 0 {
		return e.players[winner], err
	}

	return nil, err
}
