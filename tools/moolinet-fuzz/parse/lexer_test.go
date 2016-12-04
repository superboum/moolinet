package parse

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	l := &lexer{
		input: "int | int#-100,0xff",
		items: make(chan item),
	}

	go l.run()

	p := &yyParserImpl{}
	p.Parse(l)
	fmt.Println(l.grammar)
	fmt.Println(l.vars)
}
