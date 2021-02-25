package game

// Ruleset :
type Ruleset func(Neighbours) int

// Apply :
func Apply(ruleset Ruleset, cellNeighbours Neighbours) int {
	return ruleset(cellNeighbours)
}

// GenerateRuleset generates the default ruleset for Conway's Game of Life
func GenerateRuleset() Ruleset {

	return func(n Neighbours) int {
		minAlive := 2
		maxAlive := 3

		if n.c.IsEmpty() {

			if n.alive == maxAlive {
				return Alive
			}
			return Empty

		} else if n.c.IsAlive() {

			if n.alive < minAlive {
				return Dead
			}
			if n.alive >= minAlive && n.alive <= maxAlive {
				return Alive
			}
			if n.alive > maxAlive {
				return Dead
			}
		} else if n.c.IsDead() {

			if n.alive == maxAlive {
				return Alive
			}
		}
		return Dead
	}
}
