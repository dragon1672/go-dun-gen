package randai

import (
	"github.com/dragon1672/go-dun-gen/dungeon"
	"golang.org/x/exp/rand"
)

type Randy struct {
	d *dungeon.Dungeon
}

func Make(d *dungeon.Dungeon) *Randy {
	return &Randy{d: d}
}

func (r *Randy) Move() error {
	possibleMoves := r.d.ValidMoves()
	move := possibleMoves[rand.Intn(len(possibleMoves))]
	return r.d.Move(move)
}

func (r *Randy) Dungeon() *dungeon.Dungeon {
	return r.d
}
