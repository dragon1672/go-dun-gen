package gen

import (
	"encoding/hex"
	"fmt"
	"github.com/sethvargo/go-diceware/diceware"
	"golang.org/x/exp/rand"
	"hash/fnv"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
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

func (d Direction) ToRune() rune {
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

type DungeonSeed struct {
	seed string
	// Should always have at least 1 value
	previousMovements []Direction
}

func NewSeed(seed string) *DungeonSeed {
	return &DungeonSeed{
		seed: seed,
	}
}

func (d *DungeonSeed) formatRootSeed() string {
	sb := strings.Builder{}
	seedSize := strconv.Itoa(len(d.seed))
	sb.Grow(len(seedSize) + 1 + len(d.seed))
	sb.WriteString(seedSize)
	sb.WriteRune('_')
	sb.WriteString(d.seed)
	return sb.String()
}

func (d *DungeonSeed) LastMove() *Direction {
	if len(d.previousMovements) > 0 {
		return &d.previousMovements[len(d.previousMovements)-1]
	}
	return nil
}

func (d *DungeonSeed) RoomDepth() int {
	return len(d.previousMovements)
}

func (d *DungeonSeed) MoveForward(dir Direction) *DungeonSeed {
	newMoves := make([]Direction, 0, len(d.previousMovements)+1)
	if len(d.previousMovements) > 0 {
		newMoves = append(newMoves, d.previousMovements...)
	}
	newMoves = append(newMoves, dir)
	return &DungeonSeed{
		seed:              d.seed,
		previousMovements: newMoves,
	}
}

func (d *DungeonSeed) MoveBack() *DungeonSeed {
	return &DungeonSeed{
		seed:              d.seed,
		previousMovements: d.previousMovements[:len(d.previousMovements)-1],
	}
}

func (d *DungeonSeed) String() string {
	sb := strings.Builder{}
	rootSeed := d.formatRootSeed()
	sb.Grow(len(rootSeed) + len(d.previousMovements))
	sb.WriteString(rootSeed)
	for _, m := range d.previousMovements {
		sb.WriteRune(m.ToRune())
	}
	return sb.String()
}

func (d *DungeonSeed) getRandom() *rand.Rand {
	h := fnv.New32a()
	_, _ = h.Write([]byte(d.String()))
	return rand.New(rand.NewSource(uint64(h.Sum32())))
}

func (d *DungeonSeed) GetValidMoves() []Direction {
	var validMoves []Direction
	if d.LastMove() != nil {
		validMoves = append(validMoves, d.LastMove().Reverse())
	}
	r := d.getRandom()
	// Generate directions in random order
	roomChance := 100 - int(math.Min(80, float64(d.RoomDepth())*3))
	for _, roomIndex := range r.Perm(4) {
		move := directionFromIndex(roomIndex)
		if (d.LastMove() == nil || move != d.LastMove().Reverse()) && r.Intn(100) < roomChance {
			validMoves = append(validMoves, move)
		}
	}
	sort.Slice(validMoves, func(i, j int) bool {
		return validMoves[i] < validMoves[j]
	})
	return validMoves
}

func (d *DungeonSeed) GetName() string {
	gen, err := diceware.NewGenerator(&diceware.GeneratorInput{
		RandReader: d.getRandom(),
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

func (d *DungeonSeed) GetDescription() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%s: %d Moves %s, Name: %s", d.String(), len(d.GetValidMoves()), d.GetValidMoves(), d.GetName()))

	return sb.String()

}
