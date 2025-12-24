package rockpaperscissors

import (
	"github.com/alex-whitney/tictactoe/engine"
)

type ShapeType int

const (
	ShapeRock = iota
	ShapePaper
	ShapeScissors
)

var shapeName = map[ShapeType]string{
	ShapeRock:     "🪨",
	ShapePaper:    "📄",
	ShapeScissors: "✂️",
}

type Action struct {
	shape ShapeType
}

func NewAction(shape ShapeType) engine.Action {
	return &Action{
		shape,
	}
}

func (a *Action) Description() string {
	return shapeName[a.shape]
}

func (a *Action) Id() int {
	return int(a.shape)
}

func (a *Action) IsHidden() bool {
	return true
}
