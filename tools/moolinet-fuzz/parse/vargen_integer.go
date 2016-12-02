package parse

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// Maximum and minimum values for given type
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// ErrMinGreaterThanMax is an error returned when someone provide a value of min greater than max
var ErrMinGreaterThanMax = errors.New("The minimum value is greater than the max value")

// VarGenInteger is a struct storing integer parameters needed for the generation
type VarGenInteger struct {
	min, max int
	rnd      *rand.Rand
}

// NewVarGenInteger creates a new integer generator
func NewVarGenInteger() (*VarGenInteger, error) {
	return NewVarGenIntegerWithBounds(MinInt, MaxInt)
}

// NewVarGenIntegerWithBounds creates a new integer generator with given bounds
func NewVarGenIntegerWithBounds(min, max int) (*VarGenInteger, error) {
	if min > max {
		return nil, ErrMinGreaterThanMax
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &VarGenInteger{min, max, rnd}, nil
}

// GetValue return the generated value according to the VarGen interface
func (v *VarGenInteger) GetValue() string {
	// We use a float64 because if we use Intn as follow: Intn(max - min) we'll have an integer overflow
	// We don't really care about distribution or precision of the generated numbers so float64 should be enough
	gen := int(v.rnd.Int63n(int64(v.max)-int64(v.min)) + int64(v.min))
	return strconv.Itoa(gen)
}
