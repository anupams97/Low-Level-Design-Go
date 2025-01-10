package main

import "github.com/anupams97/Low-Level-Design-Go/tictactoe/game"

func main() {
	p1 := game.NewPlayer("Anupam", game.Cross)
	p2 := game.NewPlayer("Samila", game.Circle)

	newGame := game.NewGame(p1, p2)
	newGame.Play()
}
