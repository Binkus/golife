package main

import (
	"gameoflife/game"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// fmt.Println(helper.RandInt(0, 2))

	game.Start(10, 10)
}
