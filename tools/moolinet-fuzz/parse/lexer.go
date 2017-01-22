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
	input   string // string being scanned
	start   int    // start position of item
	pos     int    // current position in input
	width   int    // width of last rune read, used for backup
	line    int    // current line number
	items   chan item
	keyword *lexKeyword

	// Output and local variables
	err     string
	grammar string
	vars    []VarGen
	i       int // used as a counters of variables
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	if r == '\n' {
		l.line++
	}

	l.pos += l.width
	return
}

func (l *lexer) backup() {
	l.pos -= l.width
	l.width = 0

	if l.pos < len(l.input) && l.input[l.pos] == '\n' {
		l.line--
	}
}

func (l *lexer) emit(t int) {
	i := item{t, l.input[l.start:l.pos]}
	l.items <- i
	l.start = l.pos
}

func (l *lexer) run() {
	lexKeywords = []*lexKeyword{
		{"[int", INT, lexInt},
		{"[enum ", ENUM, lexEnum},
		{"[loop", STARTLOOP, lexInt},
		{"[/loop]", ENDLOOP, lexEnd},
	}

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

func (l *lexer) acceptRunExcept(invalid string) {
	for {
		n := l.next()
		if n == eof {
			l.Error("unexpected EOF")
			return
		}

		if strings.ContainsRune(invalid, n) {
			break
		}
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
	l.err += fmt.Sprintf("%s at line %d\n", s, l.line+1)
}

func (l *lexer) addVariable(v VarGen, err error) {
	if err != nil {
		l.Error(err.Error())
		return
	}

	l.vars = append(l.vars, v)
}

// state transitions functions

type lexKeyword struct {
	tok string
	typ int
	sta stateFn
}

var lexKeywords []*lexKeyword

func lexText(l *lexer) stateFn {
	for {
		for _, keyword := range lexKeywords {
			if strings.HasPrefix(l.input[l.pos:], keyword.tok) {
				if l.pos > l.start {
					l.emit(TEXT)
				}
				l.keyword = keyword
				return keyword.sta
			}
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
	l.pos += len(l.keyword.tok)
	l.emit(l.keyword.typ)

	s := l.next()
	if s == ' ' {
		l.ignore()
		return lexNum
	} else if s == ']' {
		l.ignore()
	} else if s == eof {
		l.Error("unexpected EOF in integer definition")
	} else {
		l.backup()
	}

	return lexText
}

func lexEnum(l *lexer) stateFn {
	l.pos += len(l.keyword.tok)
	l.emit(l.keyword.typ)

	l.acceptRunExcept("]")
	l.emit(TEXT)

	l.next()
	l.ignore() // ignore trailing "]"
	return lexText
}

func lexEnd(l *lexer) stateFn {
	l.pos += len(l.keyword.tok)
	l.emit(l.keyword.typ)
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

	// Consume optional coma
	next := l.next()
	if next == ',' {
		l.ignore()
		return lexNum
	}

	// Consume optional number separator
	if next == ']' {
		l.ignore()
		return lexText
	}

	if next == eof {
		l.Error("unexpected EOF in number definition")
	}

	l.backup()
	return lexText
}
