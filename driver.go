package main

import (
	"fmt"

	"github.com/dragon1672/go-dun-gen/dungeon"
	"golang.org/x/exp/rand"
)

func AutoPlay(seed string, limit int) {
	d := dungeon.EnterTheDungeon(seed)
	fmt.Println(d.Description())
	for i := 0; i < limit; i++ {
		possibleMoves := d.ValidMoves()
		move := possibleMoves[rand.Intn(len(possibleMoves))]
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
