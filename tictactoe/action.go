package tictactoe

import (
	"fmt"

	"github.com/alex-whitney/tictactoe/engine"
)

type Action struct {
	Space int
}

func NewAction(space int) engine.Action {
	return &Action{
		space,
	}
}

func (a *Action) Description() string {
	return fmt.Sprintf("row=%d col=%d", a.Space/3, a.Space%3)
}

func (a *Action) Id() int {
	return a.Space
}

func (a *Action) IsHidden() bool {
	return false
}
