//go:generate -command yacc go tool yacc
//go:generate yacc -o moo.go moo.y

package parse

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const eof = 0

type stateFn func(*lexer) stateFn

type item struct {
	t int
	v string
}

// lexer is inspired by Rob Pike's regexp lexer.
type lexer struct {
	input string // string being scanned
	start int    // start position of item
	pos   int    // current position in input
	width int    // width of last rune read, used for backup
	items chan item

	// Output and local variables
	grammar string
	vars    []VarGen
	i       int // used as a counter of variables
}

func (l *lexer) newVar() string {
	s := fmt.Sprintf("{{index . %d}}", l.i)
	l.i++
	return s
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return
}

func (l *lexer) backup() {
	l.pos -= l.width
	l.width = 0
}

func (l *lexer) emit(t int) {
	i := item{t, l.input[l.start:l.pos]}
	l.items <- i
	l.start = l.pos
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func (l *lexer) Lex(yylval *yySymType) int {
	i, ok := <-l.items
	if !ok {
		return eof
	}

	yylval.text = i.v
	return i.t
}

func (l *lexer) Error(s string) {
	panic(s)
}

// state transitions functions

func lexText(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], "int") {
			if l.pos > l.start {
				l.emit(TEXT)
			}
			return lexInt
		}
		if l.next() == eof {
			break
		}
	}

	// Reached eof
	if l.pos > l.start {
		l.emit(TEXT)
	}
	return nil
}

func lexInt(l *lexer) stateFn {
	l.pos += 3
	l.emit(INT)

	if l.next() == '#' {
		l.ignore()
		return lexNum
	}

	l.backup()
	return lexText
}

func lexNum(l *lexer) stateFn {
	l.accept("+-")
	digits := "0123456789"
	// Is it hex?
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	l.emit(NUM)

	// Consume optionnal coma
	if l.next() == ',' {
		l.ignore()
		return lexNum
	}
	l.backup()
	return lexText
}
