package parse

import (
	"bytes"
	"html/template"
)

// Grammar is a representation of a parsed MOO grammar.
type Grammar struct {
	tmpl *template.Template
	vars []VarGen
}

// NewGrammar returns a new grammar from a MOO string.
func NewGrammar(g string) (*Grammar, error) {
	l := &lexer{
		input: g,
		items: make(chan item),
	}

	go l.run()
	yyParse(l)

	t, err := template.New("grammar").Parse(l.grammar)
	if err != nil {
		return nil, err
	}

	return &Grammar{
		tmpl: t,
		vars: l.vars,
	}, nil
}

// Render returns a fuzzed generation fulfilling grammar specification.
func (g *Grammar) Render() ([]byte, error) {
	b := &bytes.Buffer{}
	err := g.tmpl.Execute(b, g.vars)
	return b.Bytes(), err
}
