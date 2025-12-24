package main

import (
	"fmt"
	"math/rand"

	"github.com/alex-whitney/tictactoe/connectfour"
	"github.com/alex-whitney/tictactoe/engine"
	"github.com/alex-whitney/tictactoe/rockpaperscissors"
	"github.com/alex-whitney/tictactoe/tictactoe"
)

var games = struct {
	TicTacToe         engine.Game
	ConnectFour       engine.Game
	RockPaperScissors engine.Game
}{
	tictactoe.NewGame(),
	connectfour.NewGame(),
	rockpaperscissors.NewGame(7),
}

var players = struct {
	Random       engine.Player
	Console      engine.Player
	C4PerfectBot engine.Player
}{
	engine.NewRandomPlayer("RandomBot"),
	engine.NewConsolePlayer("Human"),
	connectfour.NewPerfectBot(5),
}

func main() {
	game := games.ConnectFour

	players := []engine.Player{
		players.C4PerfectBot,
		players.Console,
	}
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	eng := engine.NewEngine(game, players[0], players[1])

	winner, err := eng.RunToCompletion()
	if err != nil {
		panic(err)
	}

	if winner == nil {
		fmt.Println("\n\nThe game has ended in a tie!")
	} else {
		fmt.Printf("\n\n%s is the winner!\n", winner.Name())
	}
}
