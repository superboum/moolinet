moolinet-fuzz
=============

This tool can be used to automatically test a user submission in moolinet, given the fact that a solution (oracle) already exists.

Given a grammar file, `moolinet-fuzz` is able to generate many random input files, send them to the oracle and the program to be tested, and compare the results.

Install
-------

```
go generate ./...
go get ./...
```

**Note**: Go 1.7 is strongly recommended for this project, because we need the GoYacc tool to generate a grammar parser.

Usage
-----

```
Usage: moolinet-fuzz [OPTIONS] GRAMMAR ORACLE SUBJECT

A fuzzer for algorithmic competition programs

Arguments:
  GRAMMAR: the .moo file describing the format expected for stdin
  ORACLE:  path to a binary that must correctly answer the problem
  SUBJECT: path to a binary to compare to the oracle

Options:
  -n int
    	number of generated test cases (default 5)
  -t duration
    	command timeout (default 1s)
  -v	enable verbose output
```

Grammar specification
---------------------

```
[int]: generates a random integer from -1000 (included) to 1000 (not included) ;
[int x,y]: generates a random integer from x (included) to y (not included) ;
[enum A,B,C]: prints A or B or C (indefinite number of parameters) ;
[loop x]mess[/loop]: repeats the \texttt{mess} sequence from 0 (included) to x times (not included) ;
[loop x,y]mess[/loop]: repeats the \texttt{mess} sequence from x (included) to y times (not included) ;
Other tokens are printed as-is (especially white-spaces).
```
