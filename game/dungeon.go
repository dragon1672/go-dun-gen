package game

import (
	"fmt"
	"go-dun-gen/game/gen"
	"golang.org/x/exp/slices"
)

type Dungeon struct {
	seed *gen.DungeonSeed
}

func EnterDungeon(seed string) *Dungeon {
	return &Dungeon{
		seed: gen.NewSeed(seed),
	}
}

func (d *Dungeon) Describe() string {
	return d.seed.GetDescription()
}

func (d *Dungeon) ValidMoves() []gen.Direction {
	return d.seed.GetValidMoves()
}

func (d *Dungeon) Move(move gen.Direction) error {
	if d.seed.LastMove() != nil && d.seed.LastMove().Reverse() == move {
		d.seed = d.seed.MoveBack()
		return nil
	}
	if !slices.Contains(d.seed.GetValidMoves(), move) {
		return fmt.Errorf("move %s is not among valid moves %v", move, d.seed.GetValidMoves())
	}
	d.seed = d.seed.MoveForward(move)

	// Validate move exists
	// Pop if moving backwards
	// Boundary check for moving out of entrance
	return nil
}
