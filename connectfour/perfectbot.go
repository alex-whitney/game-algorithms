package connectfour

import (
	"math/rand"
	"slices"

	"github.com/alex-whitney/tictactoe/engine"
)

type perfectBot struct {
	game     engine.Game
	maxDepth int

	playerNumber   int
	opponentNumber int
}

// Creates a player that chooses an action at random
func NewPerfectBot(maxDepth int) engine.Player {
	return &perfectBot{
		game:     NewGame(),
		maxDepth: maxDepth,
	}
}

func (p *perfectBot) Initialize(playerNumber int, _ engine.GameState) {
	p.playerNumber = playerNumber
	p.opponentNumber = (playerNumber + 1) % 2
}

func (p *perfectBot) Name() string {
	return "PerfectBot"
}

func (p *perfectBot) ChooseAction(state engine.GameState, validActions []engine.Action) engine.Action {
	if len(validActions) == 1 {
		return validActions[0]
	}

	root := p.walk(state, validActions, true, 1)

	maxScore := p.score(root)
	maxActions := []engine.Action{}
	for action, node := range root.children {
		if maxScore == node.value {
			maxActions = append(maxActions, action)
		}
	}

	if len(maxActions) > 0 {
		return maxActions[rand.Intn(len(maxActions))]
	} else {
		panic("something broke - action score mismatch")
	}
}

func (p *perfectBot) walk(state engine.GameState, validActions []engine.Action, myMove bool, depth int) *node {
	ret := &node{
		state:  state,
		myMove: myMove,
	}

	isOver, winner := p.game.IsFinished(state)
	if isOver {
		if winner == p.playerNumber {
			ret.value = 100
		} else if winner == p.opponentNumber {
			ret.value = -100
		}

		return ret
	}

	if depth > p.maxDepth {
		return ret
	}

	ret.children = map[engine.Action]*node{}

	num := p.playerNumber
	if !myMove {
		num = p.opponentNumber
	}

	for _, a := range validActions {
		result, err := p.game.ApplyAction(state, num, a)
		if err != nil {
			panic(err)
		}

		actions := p.game.ValidActions(result)
		ret.children[a] = p.walk(result, actions, !myMove, depth+1)
	}

	return ret
}

func (p *perfectBot) score(root *node) int {
	if len(root.children) > 0 {
		scores := []int{}
		for _, child := range root.children {
			scores = append(scores, p.score(child))
		}

		if root.myMove {
			root.value = slices.Max(scores)
		} else {
			root.value = slices.Min(scores)
		}
	}

	return root.value
}

type node struct {
	state    engine.GameState
	myMove   bool
	value    int
	children map[engine.Action]*node
}
