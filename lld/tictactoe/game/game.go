package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Board       *Board
	Player1     *Player
	Player2     *Player
	currentTurn *Player
}

func (g *Game) Play() {

	for !g.Board.IsFull() && !g.Board.HasWinner() {
		g.Board.PrintBoard()
		fmt.Printf("%s place a valid move\n", g.currentTurn.GetName())
		row, col := getValidMove()
		g.Board.Place(row, col, g.currentTurn.GetSign())

		g.switchTurn()
	}
	if g.Board.HasWinner() {
		g.switchTurn()
		fmt.Println("Player wins ", g.currentTurn.GetName())
	} else {
		fmt.Println("its a draw")
	}
	return
}

func (g *Game) switchTurn() {
	if g.currentTurn == g.Player1 {
		g.currentTurn = g.Player2
	} else {
		g.currentTurn = g.Player1
	}
}

func NewGame(p1, p2 *Player) *Game {
	return &Game{Player2: p1, Player1: p2, Board: NewBoard(), currentTurn: p1}
}

func getValidMove() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.Split(scanner.Text(), ",")
	row, _ := strconv.Atoi(s[0])
	col, _ := strconv.Atoi(s[1])
	// add validations before returning
	return row, col
}
