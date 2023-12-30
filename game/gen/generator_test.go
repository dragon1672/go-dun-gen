package gen

import (
	"golang.org/x/exp/rand"
	"reflect"
	"testing"
)

func TestDungeonSeed_GetDescription(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.GetDescription(); got != tt.want {
				t.Errorf("GetDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_GetName(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_GetValidMoves(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   []Direction
	}{
		{
			name:   "example",
			fields: fields{seed: "yar"},
			want:   []Direction{East, South, West},
		},
		{
			name:   "example_stable",
			fields: fields{seed: "yar"},
			want:   []Direction{East, South, West},
		},
		{
			name:   "example_stable2",
			fields: fields{seed: "yar"},
			want:   []Direction{East, South, West},
		},
		{
			name:   "example_stable3",
			fields: fields{seed: "yar"},
			want:   []Direction{East, South, West},
		},
		{
			name:   "example_stable4",
			fields: fields{seed: "yar"},
			want:   []Direction{East, South, West},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.GetValidMoves(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValidMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_LastMove(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   *Direction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.LastMove(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_MoveBack(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   *DungeonSeed
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.MoveBack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_MoveForward(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	type args struct {
		dir Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DungeonSeed
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.MoveForward(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_RoomDepth(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.RoomDepth(); got != tt.want {
				t.Errorf("RoomDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_String(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_formatRootSeed(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.formatRootSeed(); got != tt.want {
				t.Errorf("formatRootSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDungeonSeed_getRandom(t *testing.T) {
	type fields struct {
		seed              string
		previousMovements []Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   *rand.Rand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DungeonSeed{
				seed:              tt.fields.seed,
				previousMovements: tt.fields.previousMovements,
			}
			if got := d.getRandom(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSeed(t *testing.T) {
	type args struct {
		seed string
	}
	tests := []struct {
		name string
		args args
		want *DungeonSeed
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeed(tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directionFromIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := directionFromIndex(tt.args.i); got != tt.want {
				t.Errorf("directionFromIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
