package game

// Cell type
type Cell struct {
	state, x, y int
}

// Dead cell state
const Dead = -1

// Empty cell state
const Empty = 0

// Alive cell state
const Alive = 1

// SetState :
func (c *Cell) SetState(s int) int {
	oldState := c.state
	c.state = s
	return oldState
}

// GetPosition returns the coordinates x, y
func (c Cell) GetPosition() (x int, y int) {
	return c.x, c.y
}

// GetState returns the current cell state
func (c Cell) GetState() int {
	return c.state
}

// IsDead :
func (c Cell) IsDead() bool {
	return c.state == Dead
}

// IsEmpty :
func (c Cell) IsEmpty() bool {
	return c.state == Empty
}

// IsAlive :
func (c Cell) IsAlive() bool {
	return c.state == Alive
}

// IsDead :
func IsDead(state int) bool {
	return state == Dead
}

// IsEmpty :
func IsEmpty(state int) bool {
	return state == Empty
}

// IsAlive :
func IsAlive(state int) bool {
	return state == Alive
}
