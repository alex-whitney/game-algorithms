package main

import (
	"fmt"
	"math/rand"

	"github.com/alex-whitney/tictactoe/connectfour"
	"github.com/alex-whitney/tictactoe/engine"
	"github.com/alex-whitney/tictactoe/tictactoe"
)

var games = struct {
	TicTacToe   engine.Game
	ConnectFour engine.Game
}{
	tictactoe.NewGame(),
	connectfour.NewGame(),
}

var players = struct {
	Random  engine.Player
	Console engine.Player
}{
	engine.NewRandomPlayer("RandomBot"),
	engine.NewConsolePlayer("Human"),
}

func main() {
	game := games.TicTacToe

	players := []engine.Player{
		players.Random,
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
