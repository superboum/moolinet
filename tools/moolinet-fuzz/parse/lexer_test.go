package parse

import "testing"

func TestGrammar(t *testing.T) {
	testCases := []string{
		"[int] | [int -100,0xff] | [int 1,2]2",
		"[loop 5,10][int],[/loop][int]",
	}

	for _, c := range testCases {
		g, err := NewGrammar(c)
		if err != nil {
			t.Fatal(err)
		}

		r, err := g.Render()
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%s", r)
	}
}
