package generator

import (
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/rand"
	"reflect"
	"testing"
)

type testMove rune

func (t testMove) Rune() rune {
	return rune(t)
}

func expandRand(t *testing.T, r *rand.Rand) []float32 {
	t.Helper()
	var ret []float32
	for i := 0; i < 100; i++ {
		ret = append(ret, r.Float32())
	}
	return ret
}

func TestRoom_GetRand(t *testing.T) {
	t.Run("SameWithSameSeeds", func(t *testing.T) {
		room1 := &Room[testMove]{rawSeed: "sameSeed"}
		room2 := &Room[testMove]{rawSeed: "sameSeed"}
		got1 := expandRand(t, room1.GetRand())
		got2 := expandRand(t, room2.GetRand())
		if diff := cmp.Diff(got1, got2); diff != "" {
			t.Errorf("GetRand() mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("SameWithSameSeedsNullMoves", func(t *testing.T) {
		room1 := &Room[testMove]{rawSeed: "sameSeed", moves: nil}
		room2 := &Room[testMove]{rawSeed: "sameSeed", moves: []testMove{}}
		got1 := expandRand(t, room1.GetRand())
		got2 := expandRand(t, room2.GetRand())
		if diff := cmp.Diff(got1, got2); diff != "" {
			t.Errorf("GetRand() mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("AbleToCallMultipleTimes", func(t *testing.T) {
		room1 := &Room[testMove]{rawSeed: "testSeed"}
		got1 := expandRand(t, room1.GetRand())
		got2 := expandRand(t, room1.GetRand())
		if diff := cmp.Diff(got1, got2); diff != "" {
			t.Errorf("GetRand() mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("DifferentBetweenSeeds", func(t *testing.T) {
		room1 := &Room[testMove]{rawSeed: "test1"}
		room2 := &Room[testMove]{rawSeed: "test2"}
		got1 := expandRand(t, room1.GetRand())
		got2 := expandRand(t, room2.GetRand())
		if diff := cmp.Diff(got1, got2); diff == "" {
			t.Errorf("GetRand() expected to generate different results and did not")
		}
	})
	t.Run("DifferentWithMoves", func(t *testing.T) {
		room1 := &Room[testMove]{rawSeed: "testSeed"}
		room2 := &Room[testMove]{rawSeed: "testSeed", moves: []testMove{'X'}}
		got1 := expandRand(t, room1.GetRand())
		got2 := expandRand(t, room2.GetRand())
		if diff := cmp.Diff(got1, got2); diff == "" {
			t.Errorf("GetRand() expected to generate different results and did not")
		}
	})
	t.Run("DifferentWithSeedAndMoveOverlap", func(t *testing.T) {
		room1 := &Room[testMove]{rawSeed: "testSeedX"}
		room2 := &Room[testMove]{rawSeed: "testSeed", moves: []testMove{'X'}}
		got1 := expandRand(t, room1.GetRand())
		got2 := expandRand(t, room2.GetRand())
		if diff := cmp.Diff(got1, got2); diff == "" {
			t.Errorf("GetRand() expected to generate different results and did not")
		}
	})
}

func TestRoom_LastMove(t *testing.T) {
	tests := []struct {
		name   string
		r      *Room[testMove]
		want   testMove
		wantOk bool
	}{
		{
			name: "empty nil",
			r: &Room[testMove]{
				moves: nil,
			},
			wantOk: false,
		},
		{
			name: "empty slice",
			r: &Room[testMove]{
				moves: []testMove{},
			},
			wantOk: false,
		},
		{
			name:   "start",
			r:      Start[testMove]("testSeed"),
			wantOk: false,
		},
		{
			name: "singleEntry",
			r: &Room[testMove]{
				moves: []testMove{'R'},
			},
			want:   'R',
			wantOk: true,
		},
		{
			name: "multiEntry",
			r: &Room[testMove]{
				moves: []testMove{'A', 'B', 'C'},
			},
			want:   'C',
			wantOk: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, gotOk := tc.r.LastMove()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LastMove() got = %v, want %v", got, tc.want)
			}
			if gotOk != tc.wantOk {
				t.Errorf("LastMove() gotOk = %v, want %v", gotOk, tc.wantOk)
			}
		})
	}
}

func TestRoom_MoveBack(t *testing.T) {
	tests := []struct {
		name string
		r    *Room[testMove]
		want *Room[testMove]
	}{
		{
			name: "no moves",
			r:    &Room[testMove]{},
			want: nil,
		},
		{
			name: "start",
			r:    Start[testMove]("testSeed"),
			want: nil,
		},
		{
			name: "one move",
			r: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'m'},
			},
			want: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{},
			},
		},
		{
			name: "multi move",
			r: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'A', 'B', 'C'},
			},
			want: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'A', 'B'},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.r.MoveBack(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MoveBack() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRoom_MoveForward(t *testing.T) {
	tests := []struct {
		name string
		r    *Room[testMove]
		move testMove
		want *Room[testMove]
	}{
		{
			name: "First Move",
			r: &Room[testMove]{
				rawSeed: "testSeed",
			},
			move: 'T',
			want: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'T'},
			},
		},
		{
			name: "start",
			r:    Start[testMove]("testSeed"),
			move: 'T',
			want: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'T'},
			},
		},
		{
			name: "Second Move",
			r: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'A'},
			},
			move: 'B',
			want: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'A', 'B'},
			},
		},
		{
			name: "Multi Move",
			r: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'A', 'B'},
			},
			move: 'C',
			want: &Room[testMove]{
				rawSeed: "testSeed",
				moves:   []testMove{'A', 'B', 'C'},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.r.MoveForward(tc.move); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MoveForward() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRoom_Moves(t *testing.T) {
	tests := []struct {
		name string
		r    *Room[testMove]
		want []testMove
	}{
		{
			name: "nil",
			r:    &Room[testMove]{moves: nil},
			want: nil,
		},
		{
			name: "empty array",
			r:    &Room[testMove]{moves: []testMove{}},
			want: []testMove{},
		},
		{
			name: "start",
			r:    Start[testMove]("testSeed"),
			want: nil,
		},
		{
			name: "single value",
			r:    &Room[testMove]{moves: []testMove{'A'}},
			want: []testMove{'A'},
		},
		{
			name: "multi value",
			r:    &Room[testMove]{moves: []testMove{'A', 'B'}},
			want: []testMove{'A', 'B'},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.r.Moves(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Moves() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRoom_MovesIsCopy(t *testing.T) {
	room := &Room[testMove]{moves: []testMove{'A', 'B'}}
	got := room.Moves()
	got[0] = 'X'

	wantModified := []testMove{'X', 'B'}
	if !reflect.DeepEqual(got, wantModified) {
		t.Errorf("Moves() = %v, want %v", got, wantModified)
	}

	wantRoom := []testMove{'A', 'B'}
	if got := room.Moves(); !reflect.DeepEqual(got, wantRoom) {
		t.Errorf("Moves() = %v, want %v", got, wantRoom)
	}
	if !reflect.DeepEqual(room.moves, wantRoom) {
		t.Errorf("Moves() = %v, want %v", room.moves, wantRoom)
	}
}

func TestRoom_RoomDepth(t *testing.T) {
	tests := []struct {
		name string
		r    *Room[testMove]
		want int
	}{
		{
			name: "nil",
			r:    &Room[testMove]{moves: nil},
			want: 0,
		},
		{
			name: "empty array",
			r:    &Room[testMove]{moves: []testMove{}},
			want: 0,
		},
		{
			name: "start",
			r:    Start[testMove]("testSeed"),
			want: 0,
		},
		{
			name: "one move",
			r:    &Room[testMove]{moves: []testMove{'A'}},
			want: 1,
		},
		{
			name: "multi move",
			r:    &Room[testMove]{moves: []testMove{'A', 'B', 'C'}},
			want: 3,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.r.RoomDepth(); got != tc.want {
				t.Errorf("RoomDepth() = %v, want %v", got, tc.want)
			}
		})
	}
}
