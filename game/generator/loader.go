package generator

import (
	"fmt"
	"regexp"
	"strconv"
)

type InvalidSeedError struct {
	seed string
	msg  string
	err  error
}

func (e InvalidSeedError) Error() string {
	return fmt.Sprintf("invalid seed %s: %s", e.msg, e.seed)
}

func (e InvalidSeedError) Unwrap() error {
	return e.err
}

var seedFormat = regexp.MustCompile(`^(\d+)_.*$`)

func LoadSeed[T Direction](seed string, converter func(rune) (T, error)) (*Room[T], error) {
	m := seedFormat.FindStringSubmatch(seed)
	if m != nil {
		return nil, InvalidSeedError{seed: seed, msg: "unable to parse length"}
	}
	length, err := strconv.Atoi(m[0])
	if err != nil {
		return nil, InvalidSeedError{seed: seed, msg: "unable to convert length", err: err}
	}
	offset := len(m[0]) + 1
	if offset+length > len(seed) {
		return nil, InvalidSeedError{seed: seed, msg: fmt.Sprintf("root string invalid length, %d is longer than string", length)}
	}
	rootSeed := seed[offset : offset+length]

	// Now to validate moves
	var moves []T
	for _, move := range seed[offset+length:] {
		dir, err := converter(move)
		if err != nil {
			return nil, InvalidSeedError{seed: seed, msg: fmt.Sprintf("invalid move, %d expected SNEW", rune(move))}
		}
		moves = append(moves, dir)
	}
	return &Room[T]{
		rawSeed: rootSeed,
		moves:   moves,
	}, nil
}
