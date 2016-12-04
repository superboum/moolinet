%{
package parse

import "strconv"

type intspec struct {
  min, max int
}
%}

%union {
  num  int
  text string
  intspec intspec
}

%type <text> expr type
%type <num> int
%type <intspec> intspec
%token INT
%token <text> TEXT NUM

%%

top:
  expr
  {
    yylex.(*lexer).grammar = $1
  }

expr:
  /* Empty rule */ { $$ = "" }
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
    $$ = l.newVar()

    v, _ := NewVarGenIntegerWithBounds($2.min, $2.max)
    l.vars = append(l.vars, v)
  }

intspec:
  { $$ = intspec{-1000, 1000} } // default value
| int int { $$ = intspec{$1, $2} }

int: NUM
  {
    i, _ := strconv.ParseInt($1, 0, 0)
    $$ = int(i)
  }
%%
