package game

// Neighbours type
type Neighbours struct {
	c Cell
	empty,
	dead,
	alive,
	radius int
}

// Count the number of neighbours
func (n Neighbours) Count(f [][]int) Neighbours {
	if f == nil || len(f) < 1 || n.c.y < 0 || n.c.x < 0 ||
		n.c.y >= len(f) || n.c.x >= len(f[0]) || n.radius < 3 || n.radius%2 == 0 {
		//return Neighbours{}, fmt.Errorf("Bad boy panics")
		// fmt.Println("ERRORERROR")
		return n
	}
	halfRadius := n.radius / 2
	for i := 0; i < n.radius; i++ {
		for j := 0; j < n.radius; j++ {
			viewedY := n.c.y + i - halfRadius
			viewedX := n.c.x + j - halfRadius
			if !(viewedY == n.c.y && viewedX == n.c.x) {
				currentCellState := 0
				// if !(viewedY < 0 || viewedX < 0 || viewedY >= len(f) || viewedX >= len(f[0])) {
				if viewedY >= 0 && viewedX >= 0 && viewedY < len(f) && viewedX < len(f[viewedY]) {
					// if viewedY < 0 && viewedX < 0 && viewedY >= len(f) && viewedX >= len(f[0]) {
					currentCellState = f[viewedY][viewedX]
					// if IsAlive(currentCellState) {
					// 	n.alive++
					// }

					if viewedY < 0 || viewedX < 0 || viewedY >= len(f) || viewedX >= len(f[0]) || IsEmpty(currentCellState) {
						n.empty++
					} else if IsAlive(currentCellState) {
						n.alive++
					} else if IsDead(currentCellState) {
						n.dead++
					} else {
						n.empty++
					}
				}

				// if viewedY < 0 || viewedX < 0 || viewedY >= len(f) || viewedX >= len(f[0]) || IsEmpty(currentCellState) {
				// 	n.empty++
				// } else if IsAlive(currentCellState) {
				// 	n.alive++
				// } else if IsDead(currentCellState) {
				// 	n.dead++
				// } else {
				// 	n.empty++
				// }
			}
		}
	}
	return n
	// if r*r-1 != empty+alive+dead {
	// 	fmt.Println(":( Woups... ' should have gone ECC Memory, MemTest recommended.")
	// }

	// return Neighbours{c, empty, dead, alive, r} //, nil
	// return Neighbours{cell: c}, nil
}
