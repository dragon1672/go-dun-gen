package main

import (
	"fmt"

	"github.com/dragon1672/go-dun-gen/dungeon"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

func AutoPlay(seed string, limit int) {
	d := dungeon.EnterTheDungeon(seed)
	fmt.Println(d.Description())
	var lastMove *dungeon.Direction
	for i := 0; i < limit; i++ {
		possibleMoves := d.ValidMoves()
		if lastMove != nil && len(possibleMoves) > 1 {
			slices.DeleteFunc(possibleMoves, func(d dungeon.Direction) bool {
				return d == lastMove.Reverse()
			})
		}
		move := possibleMoves[rand.Intn(len(possibleMoves))]
		lastMove = &move
		if err := d.Move(move); err != nil {
			fmt.Printf("Encountered error: %v", err)
			return
		}
		fmt.Println(d.Description())
	}
}

func main() {
	AutoPlay("yo ho", 1000)
}
