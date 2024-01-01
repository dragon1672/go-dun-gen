package main

import (
	"github.com/dragon1672/go-dun-gen/ai"
	"github.com/dragon1672/go-dun-gen/ai/memai"
	"github.com/dragon1672/go-dun-gen/dungeon"
)

func main() {
	//ai.RunAi(randai.Make(dungeon.EnterTheDungeon("yo ho")), 1000)
	ai.RunAi(memai.Make(dungeon.EnterTheDungeon("yo ho")), 20000)
}
