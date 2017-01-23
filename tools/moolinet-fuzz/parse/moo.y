%{
package parse

import (
  "fmt"
  "strings"
)

type intspec struct {
  min, max string
}
%}

%union {
  num  string
  text string
  intspec intspec
}

%type <text> expr type loop
%type <num> int
%type <intspec> intspec
%token INT ENUM DEF STARTLOOP ENDLOOP
%token <text> TEXT NUM IDENTIFIER

%%

top:
  expr
  {
    yylex.(*lexer).grammar = $1
  }

expr:
  /* Empty rule */ { $$ = "" }
| loop definition expr ENDLOOP expr
  {
    $$ = $1 + $3 + "{{end}}" + $5
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
  INT definition intspec {
    l := yylex.(*lexer)
    $$ = fmt.Sprintf("{{index $.Vars %d}}", l.i)
    l.i++

    l.addVariable(NewVarGenIntegerWithBounds(l, $3.min, $3.max))
  }
| ENUM TEXT {
    l := yylex.(*lexer)
    $$ = fmt.Sprintf("{{index $.Vars %d}}", l.i)
    l.i++

    l.addVariable(NewVarGenEnum(strings.Split($2, ",")))
  }

intspec:
  { $$ = intspec{"-1000", "1000"} } // default value
| int { $$ = intspec{"0", $1} }
| int int { $$ = intspec{$1, $2} }

int:
  NUM { $$ = $1 }
| IDENTIFIER { $$ = "__" + $1 }

loop:
  STARTLOOP intspec {
    l := yylex.(*lexer)
    $$ = fmt.Sprintf("{{range (index $.Vars %d | $.GenRange)}}", l.i)
    l.i++

    v, _ := NewVarGenIntegerWithBounds(l, $2.min, $2.max)
    l.vars = append(l.vars, v)
  }

definition:
  {}
| DEF IDENTIFIER {
    l := yylex.(*lexer)
    l.local[$2] = l.i
  }
%%
