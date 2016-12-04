%{
package parse

import (
  "fmt"
  "strconv"
)

type intspec struct {
  min, max int
}
%}

%union {
  num  int
  text string
  intspec intspec
}

%type <text> expr type loop
%type <num> int
%type <intspec> intspec
%token INT STARTLOOP ENDLOOP
%token <text> TEXT NUM

%%

top:
  expr
  {
    yylex.(*lexer).grammar = $1
  }

expr:
  /* Empty rule */ { $$ = "" }
| loop expr ENDLOOP expr
  {
    $$ = $1 + $2 + "{{end}}" + $4
  }
| type expr
  {
    $$ = $1 + $2
  }
| TEXT expr
  {
    $$ = $1 + $2
  }

type:
  INT intspec {
    l := yylex.(*lexer)
    $$ = fmt.Sprintf("{{index $.Vars %d}}", l.i)
    l.i++

    v, _ := NewVarGenIntegerWithBounds($2.min, $2.max)
    l.vars = append(l.vars, v)
  }

intspec:
  { $$ = intspec{-1000, 1000} } // default value
| int { $$ = intspec{0, $1} }
| int int { $$ = intspec{$1, $2} }

int: NUM
  {
    i, _ := strconv.ParseInt($1, 0, 0)
    $$ = int(i)
  }

loop:
  STARTLOOP intspec
  {
    l := yylex.(*lexer)
    $$ = fmt.Sprintf("{{range (index $.Vars %d | $.GenRange)}}", l.i)
    l.i++

    v, _ := NewVarGenIntegerWithBounds($2.min, $2.max)
    l.vars = append(l.vars, v)
  }

%%
