package parse

import (
	"regexp"
	"testing"
)

func TestGrammar(t *testing.T) {
	type tc struct {
		name, grammar, regex string
		err                  bool
	}

	testCases := []tc{
		// Standard cases
		{"integer", "[int] | [int -100,0xff] | [int 10] | [int 1,2]2", `-?\d+ \| -?\d+ \| \d \| 12`, false},
		{"loop", "[loop 5,10][int],[/loop][int]", `(-?\d+,){5,9}-?\d+`, false},
		{"enum", "Good [enum afternoon,night][enum  Sir, Madam, my dearest friend,],", `Good (afternoon|night)( Sir| Madam| my dearest friend)?`, false},
		{"newline", "[int 10]\n[int 10,100]", `\d\n\d\d`, false},
		{"local_var", "[int:var 4,7] [int var,var]", `(4 4|5 5|6 6)`, false},
		{"local_var_2", "[int:var 3] [int 0,var]", `(0 0|1 0|2 0|2 1)`, false},
		{"local_var_loop", "[loop 2,2][int:x 2][int x,x] [/loop]", `(00 00 |00 11 |11 00 |11 11 )`, false},
		// Syntax error for integers
		{"err_int_0_no_space", "[int", "", true},
		{"err_int_0_other", "[int\n", `\n`, false},
		{"err_int_0_with_space", "[int ", "", true},
		{"err_int_1", "[int @]", "", true},
		{"err_int_2", "[int @,@]", "", true},
		{"err_int_3", "[int 0,1,2]", "", true},
		{"err_int_float", "[int 0.5]", "", true},
		{"err_int_exp", "[int 1e5]", "", true},
		// Syntax error for loops
		{"err_loop_unclosed", "[loop]", "", true},
		{"err_loop_unclosed_2", "[loop][loop][/loop]", "", true},
		{"err_loop_unopened", "[/loop]", "", true},
		// Syntax error for enumerations
		{"err_enum_space", "[enum ", "", true},
		{"err_enum_none", "[enum ]", "", true},
		// Syntax error for variables
		{"err_var_not_defined", "[int var]", "", true},
		// Misc errors
		{"err_tmpl", "{{ invalid }}", "", true},
	}

	for i, c := range testCases {
		func(i int, c tc) { // closure
			t.Run(c.name, func(t *testing.T) {
				g, err := NewGrammar(c.grammar)
				if c.err {
					t.Log(err)
					if err == nil {
						t.Error("Expected error in case", i)
					}
					return
				}
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
