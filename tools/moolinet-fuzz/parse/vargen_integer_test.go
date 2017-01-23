package parse

import (
	"strconv"
	"testing"
)

func TestVarGenInteger_Unbounded(t *testing.T) {
	v, err := NewVarGenInteger()
	if err != nil {
		t.Fatal("Expected nil value for error in NewVarGenInteger, got", err)
	}

	for i := 0; i < 1000; i++ {
		s := v.String()
		_, err := strconv.Atoi(s)
		if err != nil {
			t.Fatal("Expected to generate an integer, got", s, "with error", err)
		}
	}
}

func TestVarGenInteger_Bounded(t *testing.T) {
	type tc struct {
		min, max string
		err      error
	}

	cases := []tc{
		{"0", "5", nil},
		{"1", "2", nil},
		{"-10", "5", nil},
		{"5", "-5", ErrMinGreaterThanMax},
		{"a", "1", ErrCantConvert},
		{"1", "a", ErrCantConvert},
		{"", "", ErrCantConvert},
		{"1", "0.2", ErrCantConvert},
	}

	for i, c := range cases {
		v, err := NewVarGenIntegerWithBounds(nil, c.min, c.max)
		if err != c.err {
			t.Error("Expected", c.err, "error in case", i, "but got", err)
			continue
		}

		if c.err != nil {
			continue // do not continue test for errored test cases
		}

		encountered := make(map[int]bool)
		min, _ := strconv.Atoi(c.min)
		max, _ := strconv.Atoi(c.max)

		for j := 0; j < 10000; j++ {
			s := v.String()
			k, err := strconv.Atoi(s)
			if err != nil {
				t.Error("Unexpected error in case", i, "for string", s, ":", err)
				break
			}

			if k >= max || k < min {
				t.Error("Unexpected result in case", i, ":", k)
				break
			}

			encountered[k] = true
		}

		// Check validity of random
		for j := min; j < max; j++ {
			if !encountered[j] {
				t.Error("Not encountered value", j, "in case", i)
			}
		}
	}
}
