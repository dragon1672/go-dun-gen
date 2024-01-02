package memai

import (
	"fmt"
	"github.com/dragon1672/go-dun-gen/dungeon"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

type MemAi struct {
	d              *dungeon.Dungeon
	roomsAndMoves  map[string]map[dungeon.Direction]struct{}
	pathToEntrance Stack[dungeon.Direction]
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

func (m *MemAi) logAndMoveForward(move dungeon.Direction) error {
	m.getCurrentRoomLog()[move] = struct{}{}
	m.pathToEntrance.Push(move.Reverse())
	return m.d.Move(move)
}

func (m *MemAi) logAndMoveBack() error {
	move, ok := m.pathToEntrance.Pop()
	if !ok {
		return fmt.Errorf("out of moves, discovered %d rooms", len(m.roomsAndMoves))
	}
	m.getCurrentRoomLog()[move] = struct{}{}
	return m.d.Move(move)
}

func (m *MemAi) Move() error {
	possibleMoves := m.d.ValidMoves()
	// prune already visited rooms
	possibleMoves = slices.DeleteFunc(possibleMoves, func(d dungeon.Direction) bool {
		if lastMove, ok := m.pathToEntrance.Peek(); ok && d == lastMove {
			return true // prune backtracking
		}
		if _, ok := m.getCurrentRoomLog()[d]; ok {
			return true //  prune visited rooms
		}
		return false // allow unique new rooms
	})

	if len(possibleMoves) == 0 {
		return m.logAndMoveBack() // backtrack if no new rooms to explore
	}
	move := possibleMoves[rand.Intn(len(possibleMoves))]
	return m.logAndMoveForward(move)
}

func (m *MemAi) Dungeon() *dungeon.Dungeon {
	return m.d
}
