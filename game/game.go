// Package game Game of Life
package game

import (
	"fmt"
)

// Start game
func Start(height, width int) {
	b := CreateBoard(height, width)

	b.GenRandomCells()

	b.field[5][4] = Alive
	b.field[5][5] = Alive
	b.field[5][6] = Alive

	r := GenerateRuleset()

	b.Simulate(0, 3, r, 0, 0, true)
	fmt.Print(OutputTextField(b.field, true))
	b.Simulate(1, 3, r, 0, 0, true)
	fmt.Print(OutputTextField(b.field, true))
	b.Simulate(1, 3, r, 0, 0, true)
	fmt.Print(OutputTextField(b.field, true))
	b.Simulate(1, 3, r, 0, 0, true)
	fmt.Print(OutputTextField(b.field, true))
}

var (
	// AliveCell char:
	AliveCell = "+"
	// DeadCell char:
	DeadCell = "-"
	// EmptyCell char:
	EmptyCell = " "
)

// OutputTextField :
func OutputTextField(field [][]int, enableUnnecessaryHashes bool) (s string) {
	minlength := 3
	if field == nil || len(field) < minlength || len(field[0]) < minlength {
		// fmt.Println("Bad Request: Your field doesn't exist or it is too small.")
		return
	}
	toprinthashes := ""
	if enableUnnecessaryHashes {
		toprinthashes = "#"
	}
	// Little don't repeat yourself fix
	printHashes := func(x int) {
		for j := 0; j < len(field[x])+2; j++ {
			s += "#"
		}
		s += "\n"
	}
	for i := 0; i < len(field); i++ {
		// Generate first line of ###s
		if enableUnnecessaryHashes && i == 0 {
			printHashes(i)
		}
		s += toprinthashes

		for j := 0; j < len(field[i]); j++ {
			toprintchar := EmptyCell
			switch field[i][j] {
			case 1:
				toprintchar = AliveCell
			case -1:
				toprintchar = DeadCell
			}
			s += string(toprintchar)
		}
		s += toprinthashes
		s += "\n"

		// Generate last line of ###s
		if enableUnnecessaryHashes && i == len(field)-1 {
			printHashes(i)
		}
	}
	return
}
