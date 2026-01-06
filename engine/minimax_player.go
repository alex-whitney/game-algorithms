package engine

import (
	"math/rand"
	"slices"
)

type StateScoringFunc func(game Game, state GameState, playerNumber int, opponentNumber int) float64

type minimaxPlayer struct {
	game        Game
	maxDepth    int
	scoringFunc StateScoringFunc

	playerNumber   int
	opponentNumber int
}

// Creates a player that chooses an action at random
func NewMinimaxPlayer(game Game, maxDepth int, scoringFunc StateScoringFunc) Player {
	return &minimaxPlayer{
		game:        game,
		maxDepth:    maxDepth,
		scoringFunc: scoringFunc,
	}
}

func (p *minimaxPlayer) Initialize(playerNumber int, _ GameState) {
	p.playerNumber = playerNumber
	p.opponentNumber = (playerNumber + 1) % 2
}

func (p *minimaxPlayer) Name() string {
	return "Minimax"
}

func (p *minimaxPlayer) ChooseAction(state GameState, validActions []Action) Action {
	if len(validActions) == 1 {
		return validActions[0]
	}

	root := p.walk(state, validActions, true, 1)

	maxActions := []Action{}
	for action, node := range root.children {
		if root.value == node.value {
			maxActions = append(maxActions, action)
		}
	}

	if len(maxActions) > 0 {
		return maxActions[rand.Intn(len(maxActions))]
	} else {
		panic("something broke - action score mismatch")
	}
}

func (p *minimaxPlayer) walk(state GameState, validActions []Action, myMove bool, depth int) *node {
	ret := &node{
		state:  state,
		myMove: myMove,
		value:  p.scoringFunc(p.game, state, p.playerNumber, p.opponentNumber),
	}

	isOver, _ := p.game.IsFinished(state)
	if isOver || depth > p.maxDepth {
		return ret
	}

	ret.children = map[Action]*node{}
	scores := []float64{}

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

		child := p.walk(result, actions, !myMove, depth+1)
		ret.children[a] = child
		scores = append(scores, child.value)
	}

	if myMove {
		ret.value = slices.Max(scores)
	} else {
		ret.value = slices.Min(scores)
	}

	return ret
}

type node struct {
	state    GameState
	myMove   bool
	value    float64
	children map[Action]*node
}
