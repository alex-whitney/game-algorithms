package engine

import "fmt"

type consolePlayer struct {
	name string
}

// Creates a new Player that can be controlled via the the console
func NewConsolePlayer(name string) Player {
	return &consolePlayer{
		name,
	}
}

func (p *consolePlayer) Initialize(playerNumber int, state GameState) {
	// Do nothing
}

func (p *consolePlayer) Name() string {
	return p.name
}

func (p *consolePlayer) ChooseAction(state GameState, actions []Action) Action {
	fmt.Println()
	fmt.Println(state.String())

	for ind, a := range actions {
		fmt.Printf("%d: %s\n", ind+1, a.Description())
	}

	var err error
	var actionNum int

	for actionNum <= 0 || actionNum > len(actions) || err != nil {
		fmt.Println("\nChoose an action:")
		_, err = fmt.Scanln(&actionNum)
	}

	return actions[actionNum-1]
}
