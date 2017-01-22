package parse

import (
	"errors"
	"math/rand"
	"time"
)

// ErrNotEnoughData is emitted when a user asks an enum type with no possibility.
var ErrNotEnoughData = errors.New("enum type requires at least one possibility")

// VarGenEnum is a struct storing enumeration parameters needed for the generation.
type VarGenEnum struct {
	data []string
	rnd  *rand.Rand
}

// NewVarGenEnum returns a new VarGenEnum with provided possibilities.
func NewVarGenEnum(data []string) (*VarGenEnum, error) {
	if len(data) == 0 || len(data) == 1 && data[0] == "" {
		return nil, ErrNotEnoughData
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &VarGenEnum{data, rnd}, nil
}

// String returns the generated value according to the VarGen interface.
func (v *VarGenEnum) String() string {
	return v.data[v.rnd.Intn(len(v.data))]
}
