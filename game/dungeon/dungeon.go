package dungeon

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/dragon1672/go-dun-gen/game/generator"
	"github.com/sethvargo/go-diceware/diceware"
	"golang.org/x/exp/slices"
)

type Direction rune

const (
	North            = 'N'
	South            = 'S'
	East             = 'E'
	West             = 'W'
	invalidDirection = '?'
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	}
	return "?" + hex.EncodeToString([]byte{byte(d)}) + "?"
}

func (d Direction) Reverse() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	}
	log.Fatalf("direction %b is invalid", d)
	return invalidDirection
}

func (d Direction) Rune() rune {
	return rune(d)
}

func directionFromIndex(i int) Direction {
	switch i {
	case 0:
		return North
	case 1:
		return South
	case 2:
		return East
	case 3:
		return West
	}
	log.Fatalf("direction index got %d want between 0 and 3", i)
	return invalidDirection
}

type Dungeon struct {
	room *generator.Room[Direction]
}

func (d *Dungeon) ValidMoves() []Direction {
	var validMoves []Direction
	if lastMove, ok := d.room.LastMove(); ok {
		validMoves = append(validMoves, lastMove.Reverse())
	}
	r := d.room.GetRand()
	// Generate directions in random order
	roomChance := 100 - int(math.Min(80, float64(d.room.RoomDepth())*3))
	for _, roomIndex := range r.Perm(4) {
		move := directionFromIndex(roomIndex)
		if lastMove, ok := d.room.LastMove(); (!ok || move != lastMove.Reverse()) && r.Intn(100) < roomChance {
			validMoves = append(validMoves, move)
		}
	}
	sort.Slice(validMoves, func(i, j int) bool {
		return validMoves[i] < validMoves[j]
	})
	return validMoves
}

func (d *Dungeon) Move(move Direction) error {
	if lastMove, ok := d.room.LastMove(); ok && lastMove.Reverse() == move {
		d.room = d.room.MoveBack()
		return nil
	}
	if !slices.Contains(d.ValidMoves(), move) {
		return fmt.Errorf("move %s is not among valid moves %v", move, d.ValidMoves())
	}
	d.room = d.room.MoveForward(move)
	return nil
}

func (d *Dungeon) Name() string {
	gen, err := diceware.NewGenerator(&diceware.GeneratorInput{
		RandReader: d.room.GetRand(),
	})
	if err != nil {
		log.Fatal(err)
	}
	list, err := gen.Generate(3)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Join(list, "-")
}

func (d *Dungeon) Description() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%s: %d Moves %s, Name: %s", d.room.String(), len(d.ValidMoves()), d.ValidMoves(), d.Name()))

	return sb.String()
}

func EnterTheDungeon(seed string) *Dungeon {
	return &Dungeon{
		room: generator.Start[Direction](seed),
	}
}