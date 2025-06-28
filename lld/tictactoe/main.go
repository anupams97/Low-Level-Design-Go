package main

import (
	game2 "github.com/anupams97/Low-Level-Design-Go/lld/tictactoe/game"
)

func main() {
	p1 := game2.NewPlayer("Anupam", game2.Cross)
	p2 := game2.NewPlayer("Samila", game2.Circle)

	newGame := game2.NewGame(p1, p2)
	newGame.Play()
}
