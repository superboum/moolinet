package parse

import (
	"regexp"
	"testing"
)

func TestGrammar(t *testing.T) {
	type tc struct {
		name, grammar, regex string
	}

	testCases := []tc{
		{"integer", "[int] | [int -100,0xff] | [int 1,2]2", `\d+ \| -?\d+ \| 12`},
		{"loop", "[loop 5,10][int],[/loop][int]", `(-?\d+,){5,9}-?\d+`},
		{"enum", "Good [enum afternoon,night][enum  Sir, Madam, my dearest friend,],", `Good (afternoon|night)( Sir| Madam| my dearest friend)?`},
	}

	for i, c := range testCases {
		func(i int, c tc) { // closure
			t.Run(c.name, func(t *testing.T) {
				g, err := NewGrammar(c.grammar)
				if err != nil {
					t.Error("Unable to compile grammar in case", i, ":", err)
					return
				}

				reg := regexp.MustCompile(c.regex)

				for j := 0; j < 1000; j++ {
					r, err := g.Render()
					if err != nil {
						t.Error("Unable to generate test case in case", i, ":", err)
						break
					}

					if !reg.Match(r) {
						t.Error("Invalid result for case", i, ":", string(r))
						break
					}
				}
			})
		}(i, c)
	}
}
