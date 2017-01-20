package parse

import "testing"

func TestVarGenEnum(t *testing.T) {
	type tc struct {
		data []string
		err  error
	}

	cases := []tc{
		{[]string{"Hello"}, nil},
		{[]string{"A", "B", "C"}, nil},
		{[]string{}, ErrNotEnoughData},
		{nil, ErrNotEnoughData},
	}

	for i, c := range cases {
		v, err := NewVarGenEnum(c.data)
		if err != c.err {
			t.Error("Expected", c.err, "error in case", i, "but got", err)
		}

		if c.err != nil {
			continue // do not continue test for errored test cases
		}

		encountered := make(map[string]bool)

		for j := 0; j < 10000; j++ {
			s := v.String()
			var contains bool
			for _, e := range c.data {
				if s == e {
					contains = true
					break
				}
			}

			if !contains {
				t.Error("Unexpected result in case", i, ":", s)
				break
			}

			encountered[s] = true
		}

		// Check validity of random
		for _, e := range c.data {
			if !encountered[e] {
				t.Error("Not encountered value", e, "in case", i)
			}
		}
	}
}
