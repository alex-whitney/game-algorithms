package connectfour

import (
	"fmt"

	"github.com/alex-whitney/tictactoe/engine"
)

type Action struct {
	Column int
}

func NewAction(column int) engine.Action {
	return &Action{
		Column: column,
	}
}

func (a *Action) Description() string {
	return fmt.Sprintf("Column %d", a.Column+1)
}

func (a *Action) Id() int {
	return a.Column
}

func (a *Action) IsHidden() bool {
	return false
}
