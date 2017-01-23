package parse

import (
	"bytes"
	"errors"
	"html/template"
	"strconv"
)

type tmplInput struct {
	Vars []VarGen
}

// GenRange is used in template geration for simple loop management.
func (t *tmplInput) GenRange(n *VarGenInteger) []bool {
	i, _ := strconv.Atoi(n.String())
	return make([]bool, i)
}

// Grammar is a representation of a parsed MOO grammar.
type Grammar struct {
	tmpl  *template.Template
	input *tmplInput
}

// NewGrammar returns a new grammar from a MOO string.
func NewGrammar(g string) (*Grammar, error) {
	yyErrorVerbose = true

	l := &lexer{
		input: g,
		items: make(chan item),
		local: make(map[string]int),
	}

	go l.run()
	yyParse(l)

	if len(l.err) > 0 {
		return nil, errors.New(l.err)
	}

	t, err := template.New("grammar").Parse(l.grammar)
	if err != nil {
		return nil, err
	}

	return &Grammar{
		tmpl: t,
		input: &tmplInput{
			Vars: l.vars,
		},
	}, nil
}

// Render returns a fuzzed generation fulfilling grammar specification.
func (g *Grammar) Render() ([]byte, error) {
	b := &bytes.Buffer{}
	err := g.tmpl.Execute(b, g.input)
	return b.Bytes(), err
}
