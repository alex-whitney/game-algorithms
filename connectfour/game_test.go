package connectfour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFinished(t *testing.T) {
	cases := []struct {
		name   string
		state  *State
		isOver bool
		winner int
	}{{
		"empty",
		&State{
			Board: [][]int{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			TurnCounter: 0,
		},
		false,
		0,
	}, {
		"verticalBottom",
		&State{
			Board: [][]int{
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			TurnCounter: 4,
		},
		true,
		0,
	}, {
		"verticalTop",
		&State{
			Board: [][]int{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 2, 0, 0, 0},
			},
			TurnCounter: 4,
		},
		true,
		1,
	}, {
		"horizontalBottomRow",
		&State{
			Board: [][]int{
				{0, 0, 0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			TurnCounter: 4,
		},
		true,
		0,
	}, {
		"horizontalTopRow",
		&State{
			Board: [][]int{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{2, 2, 2, 2, 0, 0, 0},
			},
			TurnCounter: 4,
		},
		true,
		1,
	}, {
		"rightDiagBottom",
		&State{
			Board: [][]int{
				{2, 0, 0, 0, 0, 0, 0},
				{0, 2, 0, 0, 0, 0, 0},
				{0, 0, 2, 0, 0, 0, 0},
				{0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			TurnCounter: 4,
		},
		true,
		1,
	}, {
		"rightDiagTop",
		&State{
			Board: [][]int{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			TurnCounter: 4,
		},
		true,
		0,
	}, {
		"leftDiagBottom",
		&State{
			Board: [][]int{
				{0, 0, 0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			TurnCounter: 4,
		},
		true,
		0,
	}, {
		"leftDiagTop",
		&State{
			Board: [][]int{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 2, 0, 0, 0},
				{0, 0, 2, 0, 0, 0, 0},
				{0, 2, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 0},
			},
			TurnCounter: 4,
		},
		true,
		1,
	}, {
		"tie",
		&State{
			Board: [][]int{
				{2, 1, 2, 1, 2, 1, 2},
				{1, 2, 1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1, 2, 1},
				{2, 1, 2, 1, 2, 1, 2},
				{2, 1, 2, 1, 2, 1, 2},
			},
			TurnCounter: 6 * 7,
		},
		true,
		-1,
	}, {
		"oneMoreMove",
		&State{
			Board: [][]int{
				{2, 1, 2, 1, 2, 1, 2},
				{1, 2, 1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1, 2, 1},
				{2, 1, 2, 1, 2, 1, 2},
				{2, 1, 2, 1, 2, 1, 0},
			},
			TurnCounter: 6*7 - 1,
		},
		false,
		0,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			game := NewGame()
			isOver, winner := game.IsFinished(tc.state)

			assert.Equal(t, tc.isOver, isOver, "isOver")
			assert.Equal(t, tc.winner, winner, "winner")
		})
	}
}
