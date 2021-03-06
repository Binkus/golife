package game

import (
	"gameoflife/helper"
	"gameoflife/helper/linq"
)

// Board type of Life
type Board struct {
	field,
	prevBoard,
	prev2Board,
	prev3Board [][]int
	iterationStack [][][]int
}

// CreateBoard :
func CreateBoard(height, width int) (b Board) {
	b.field = make([][]int, height)
	for i := range b.field {
		b.field[i] = make([]int, width)
	}
	return
}

// TypeIterationStack enum
type TypeIterationStack int

// TypeIterationStackIntEnum
const (
	Disabled TypeIterationStack = iota
	Enabled
	EnabledButClear
)

func (b *Board) setPreviousFields(field [][]int) {
	b.prev3Board = b.prev2Board
	b.prev2Board = b.prevBoard
	b.prevBoard = field
}

// GenRandomCells :
func (b *Board) GenRandomCells() {
	for i := 0; i < len(b.field); i++ {
		for j := 0; j < len(b.field[i]); j++ {
			b.field[i][j] = helper.RandInt(0, 2)
		}
	}
}

// Simulate Conway's Game of Life with dynamic ruleset
func (b *Board) Simulate(t, viewedRadius int, ruleset func(Neighbours) int,
	defaultStateWhenRuleNull int, typeIterationStack TypeIterationStack,
	stopWhenNoChange bool) [][]int {

	if b.field == nil || len(b.field) < 1 || t < 1 {
		return b.field
	}
	var nField = linq.Copy2DimSliceByValue(b.field)
	if typeIterationStack != 0 {
		if typeIterationStack == 2 {
			b.iterationStack = nil
		}
		if len(b.iterationStack) == 0 {
			b.iterationStack = append(b.iterationStack, linq.Copy2DimSliceByValue(b.field))
		}
	}

	for ; t > 0; t-- {
		field := linq.Copy2DimSliceByValue(nField)
		if stopWhenNoChange {
			b.setPreviousFields(field)
		}
		for i := 0; i < len(b.field); i++ {
			for j := 0; j < len(b.field[i]); j++ {
				nField[i][j] = ruleset(Neighbours{radius: viewedRadius, c: Cell{field[i][j], j, i}}.Count(field))
			}
		}
		b.field = nField
		if stopWhenNoChange && t > 0 {
			if linq.EqualsByValue(field, nField) {
				break
			} else if linq.EqualsByValue(b.prev2Board, nField) {
				t = t % 2
			} else if t > 1 && linq.EqualsByValue(b.prev3Board, nField) {
				t = t % 3
			}
		}
		if typeIterationStack != 0 {
			b.iterationStack = append(b.iterationStack, linq.Copy2DimSliceByValue(nField))
		}
	}
	return b.field
}
