package game

import "fmt"

type Board struct {
	cells     [3][3]Sign
	moveCount int
}

func (b *Board) IsFull() bool {
	return b.moveCount == 9
}

// HasWinner check if there is a winner
func (b *Board) HasWinner() bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if b.cells[i][0] != "" && b.cells[i][0] == b.cells[i][1] && b.cells[i][1] == b.cells[i][2] {
			return true
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if b.cells[0][j] != "" && b.cells[0][j] == b.cells[1][j] && b.cells[1][j] == b.cells[2][j] {
			return true
		}
	}

	// Check main diagonal
	if b.cells[0][0] != "" && b.cells[0][0] == b.cells[1][1] && b.cells[1][1] == b.cells[2][2] {
		return true
	}

	// Check anti-diagonal
	if b.cells[0][2] != "" && b.cells[0][2] == b.cells[1][1] && b.cells[1][1] == b.cells[2][0] {
		return true
	}

	return false
}

func (b *Board) PrintBoard() {
	fmt.Println("-------------") // Top border
	for i := 0; i < 3; i++ {
		fmt.Print("| ")
		for j := 0; j < 3; j++ {
			fmt.Print(b.cells[i][j], " | ") // Print cell value with separator
		}
		fmt.Println("\n-------------") // Row separator
	}
}

func NewBoard() *Board {
	return &Board{cells: [3][3]Sign{}}
}

func (b *Board) Place(r, c int, s Sign) {
	b.cells[r][c] = s
	b.moveCount++
}
