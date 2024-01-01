package ai

import (
	"fmt"
	"github.com/dragon1672/go-dun-gen/dungeon"
	"strconv"
)

type AI interface {
	Move() error
	Dungeon() *dungeon.Dungeon
}

func RunAi(ai AI, limit int) {
	pad := len(strconv.Itoa(limit))
	fmt.Printf("%*d - %s\n", pad, 0, ai.Dungeon().Description())
	for i := 0; i < limit; i++ {
		if err := ai.Move(); err != nil {
			fmt.Printf("Encountered error: %v", err)
			return
		}
		fmt.Printf("%*d - %s\n", pad, i+1, ai.Dungeon().Description())
	}
	fmt.Println("Yay!")
}
