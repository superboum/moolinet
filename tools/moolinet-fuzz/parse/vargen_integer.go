package parse

import (
	"errors"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

// Maximum and minimum values for given type
const (
	MaxUint = ^uint64(0)
	MinUint = 0
	MaxInt  = int64(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// ErrMinGreaterThanMax is an error returned when someone provide a value of min greater than max
var ErrMinGreaterThanMax = errors.New("the minimum value is greater than the max value")

// ErrCantConvert is returned while encountering a conversion error
var ErrCantConvert = errors.New("can't convert the given string to a bigint")

// VarGenInteger is a struct storing integer parameters needed for the generation
type VarGenInteger struct {
	min, max *big.Int
	rnd      *rand.Rand
}

// NewVarGenInteger creates a new integer generator in the int64
func NewVarGenInteger() (*VarGenInteger, error) {
	return NewVarGenIntegerWithBounds(strconv.FormatInt(MinInt, 10), strconv.FormatInt(MaxInt, 10))
}

// NewVarGenIntegerWithBounds creates a new integer generator with given bounds
func NewVarGenIntegerWithBounds(minStr, maxStr string) (*VarGenInteger, error) {
	// Parse min and max
	min, success := (&big.Int{}).SetString(minStr, 0)
	if !success {
		return nil, ErrCantConvert
	}
	max, success := (&big.Int{}).SetString(maxStr, 0)
	if !success {
		return nil, ErrCantConvert
	}

	// Check than min is smaller than max
	if min.Cmp(max) == 1 {
		return nil, ErrMinGreaterThanMax
	}

	// Init random
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &VarGenInteger{min, max, rnd}, nil
}

// GetValue return the generated value according to the VarGen interface
func (v *VarGenInteger) String() string {
	// Calculate the range between max and min (ex: max=10, min=-10, range = 20)
	gen := (&big.Int{}).Sub(v.max, v.min)

	// Generate a random number in this range (ex: [0, 20])
	gen.Rand(v.rnd, gen)

	// Shift the range to its original position (ex: [-10, 10])
	gen.Add(gen, v.min)

	// Return the number, use decimal format
	return gen.Text(10)
}
