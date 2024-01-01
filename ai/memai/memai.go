package memai

import (
	"github.com/dragon1672/go-dun-gen/dungeon"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

type MemAi struct {
	d             *dungeon.Dungeon
	roomsAndMoves map[string]map[dungeon.Direction]struct{}
	lastMove      *dungeon.Direction
}

func Make(d *dungeon.Dungeon) *MemAi {
	return &MemAi{
		d:             d,
		roomsAndMoves: make(map[string]map[dungeon.Direction]struct{}),
	}
}

func (m *MemAi) getCurrentRoomLog() map[dungeon.Direction]struct{} {
	if m.roomsAndMoves[m.d.Name()] == nil {
		m.roomsAndMoves[m.d.Name()] = make(map[dungeon.Direction]struct{})
	}
	return m.roomsAndMoves[m.d.Name()]
}

func (m *MemAi) logAndMove(move dungeon.Direction) error {
	m.lastMove = &move
	m.getCurrentRoomLog()[move] = struct{}{}
	return m.d.Move(move)
}

func (m *MemAi) Move() error {
	possibleMoves := m.d.ValidMoves()
	// prune already visited rooms
	possibleMoves = slices.DeleteFunc(possibleMoves, func(d dungeon.Direction) bool {
		if m.lastMove != nil && d == m.lastMove.Reverse() {
			return true // prune backtracking
		}
		if _, ok := m.getCurrentRoomLog()[d]; ok {
			return true //  prune visited rooms
		}
		return false // allow unique new rooms
	})

	// Add backtracking if no new unique rooms to explore
	if m.lastMove != nil && len(possibleMoves) == 0 {
		possibleMoves = []dungeon.Direction{m.lastMove.Reverse()}
	}

	move := possibleMoves[rand.Intn(len(possibleMoves))]
	return m.logAndMove(move)
}

func (m *MemAi) Dungeon() *dungeon.Dungeon {
	return m.d
}
