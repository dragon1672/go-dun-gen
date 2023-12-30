package main

import (
	"fmt"
	"go-dun-gen/game"
	"go-dun-gen/game/gen"
	"golang.org/x/exp/rand"
)

func PrePlay(seed string, moves []gen.Direction) {
	d := game.EnterDungeon(seed)
	fmt.Println(d.Describe())
	for i, move := range moves {
		fmt.Printf("Move %d: %s\n", i, move.String())
		if err := d.Move(move); err != nil {
			fmt.Printf("Encountered error: %v", err)
			return
		}
		fmt.Println(d.Describe())
	}
}

func AutoPlay(seed string) {
	d := game.EnterDungeon(seed)
	fmt.Println(d.Describe())
	for {
		possibleMoves := d.ValidMoves()
		move := possibleMoves[rand.Intn(len(possibleMoves))]
		if err := d.Move(move); err != nil {
			fmt.Printf("Encountered error: %v", err)
			return
		}
		fmt.Println(d.Describe())
	}
}

func main() {
	/*
		PrePlay("yar", []gen.Direction{
			gen.South,
			gen.South,
			gen.North,
			gen.South,
			gen.East,
			gen.South,
			gen.South,
			gen.West,
			gen.South,
			gen.East,
			gen.South,
			gen.South,
			gen.South,
			gen.South,
		})
		//*/
	AutoPlay("yo ho")
}
