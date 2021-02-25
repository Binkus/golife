package game

import (
	"fmt"
)

// Start game
func Start(height, width int) {
	b := CreateBoard(height, width)
	b.field[5][4] = Alive
	b.field[5][5] = Alive
	b.field[5][6] = Alive

	// stopWhen := func(int) bool {

	// }

	// c := func(x int) bool {
	// 	return x == Dead
	// }

	f := func() {
		//fmt.Println(linq.Any(b.field, c))
	}

	r := GenerateRuleset()

	b.Simulate(0, 3, r, 0, 0, true)
	f()
	fmt.Print(OutputTextField(b.field, true))
	b.Simulate(1, 3, r, 0, 0, true)
	f()
	fmt.Print(OutputTextField(b.field, true))
	b.Simulate(1, 3, r, 0, 0, true)
	f()
	fmt.Print(OutputTextField(b.field, true))
	b.Simulate(1, 3, r, 0, 0, true)
	f()
	fmt.Print(OutputTextField(b.field, true))
}

const aliveCell = "+"
const deadCell = "-"
const emptyCell = " "
const emptyNulCell = ""

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
	for i := 0; i < len(field); i++ {

		// Little don't repeat yourself fix
		printHashes := func(x int) {
			for j := 0; j < len(field[x])+2; j++ {
				// fmt.Print('#')
				s += "#"
			}
			// fmt.Println()
			s += "\n"
		}

		// Generate first line of ###s
		if enableUnnecessaryHashes && i == 0 {
			printHashes(i)
		}

		// fmt.Print(toprinthashes)
		s += toprinthashes

		for j := 0; j < len(field[i]); j++ {
			toprintchar := emptyCell
			switch field[i][j] {
			case 1:
				toprintchar = aliveCell
				break
			case -1:
				toprintchar = deadCell
				break
				// case 0:
				// 	toprintchar = emptyCell
			}

			// fmt.Print(toprintchar)
			s += string(toprintchar)

		}
		// fmt.Println(toprinthashes)
		s += toprinthashes
		s += "\n"

		// Generate last line of ###s
		if enableUnnecessaryHashes && i == len(field)-1 {
			printHashes(i)
		}
	}
	return
}
