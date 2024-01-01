package generator

import (
	"fmt"
	"golang.org/x/exp/slices"
	"hash/fnv"
	"strings"

	"golang.org/x/exp/rand"
)

type Direction interface {
	Rune() rune
}

type Room[T Direction] struct {
	rawSeed string
	moves   []T
}

func Start[T Direction](rawSeed string) *Room[T] {
	return &Room[T]{
		rawSeed: rawSeed,
	}
}

func (r *Room[T]) Moves() []T {
	return slices.Clone(r.moves)
}

func (r *Room[T]) LastMove() (T, bool) {
	if len(r.moves) > 0 {
		return r.moves[len(r.moves)-1], true
	}
	var t T
	return t, false
}

func (r *Room[T]) RoomDepth() int {
	return len(r.moves)
}

func (r *Room[T]) MoveForward(move T) *Room[T] {
	newMoves := make([]T, 0, len(r.moves)+1)
	if len(r.moves) > 0 {
		newMoves = append(newMoves, r.moves...)
	}
	newMoves = append(newMoves, move)
	return &Room[T]{
		rawSeed: r.rawSeed,
		moves:   newMoves,
	}
}

func (r *Room[T]) MoveBack() *Room[T] {
	if len(r.moves) == 0 {
		return nil
	}
	return &Room[T]{
		rawSeed: r.rawSeed,
		moves:   r.moves[:len(r.moves)-1],
	}
}

// formatRawSeed will prefix seed with the length of the seed. This will make the raw seed parsable from the movements.
func (r *Room[T]) formatRawSeed() string {
	return fmt.Sprintf("%d-%s", len(r.rawSeed), r.rawSeed)
}

func (r *Room[T]) String() string {
	sb := strings.Builder{}
	rootSeed := r.formatRawSeed()
	sb.Grow(len(rootSeed) + len(r.moves))
	sb.WriteString(rootSeed)
	for _, m := range r.moves {
		sb.WriteRune(m.Rune())
	}
	return sb.String()
}

// GetRand returns a freshly initialized randomizer seeded based on this room.
func (r *Room[T]) GetRand() *rand.Rand {
	h := fnv.New32a()
	_, _ = h.Write([]byte(r.String()))
	return rand.New(rand.NewSource(uint64(h.Sum32())))
}
