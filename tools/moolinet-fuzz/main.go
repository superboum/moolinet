package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/superboum/moolinet/tools/moolinet-fuzz/parse"
)

var nb int
var verbose bool
var timeout time.Duration

func init() {
	flag.IntVar(&nb, "n", 5, "number of generated test cases")
	flag.BoolVar(&verbose, "v", false, "enable verbose output")
	flag.DurationVar(&timeout, "t", time.Second, "command timeout")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 3 {
		flag.Usage()
		return
	}

	g, o, s := args[0], args[1], args[2]

	// Read and parse grammar
	data, err := ioutil.ReadFile(g)
	check(err)

	grammar, err := parse.NewGrammar(string(data))
	check(err)

	for i := 0; i < nb; i++ {
		fmt.Printf("Test case %d/%d... ", i+1, nb)
		// Generate test case
		in, err := grammar.Render()
		check(err)
		debug(in, verbose, "Input")

		// Run on oracle
		expected, err := run(in, o)
		check(err)
		debug(expected, verbose, "Expected")

		// Run on test
		got, err := run(in, s)
		check(err)
		debug(got, verbose, "Got")

		// Verify
		if !bytes.Equal(expected, got) {
			fmt.Println("ERROR")
			debug(in, !verbose, "Input")
			debug(expected, !verbose, "Expected")
			debug(got, !verbose, "Got")
			os.Exit(1)
		}
		fmt.Println("OK")
	}
}

func run(in []byte, path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	args := strings.Split(path, " ")

	// We know that executing command from vars could be dangerous but this application
	// should be run in a sandbox OR with trusted programs
	// Without #nosec GAS Linter will throw an error
	cmd := exec.CommandContext(ctx, args[0], args[1:]...) // #nosec
	cmd.Stdin = bytes.NewReader(in)
	return cmd.Output()
}

func debug(content []byte, verbose bool, msg string) {
	if verbose {
		fmt.Println("-- " + msg + " -------")
		fmt.Printf("%s\n", content)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: moolinet-fuzz [OPTIONS] GRAMMAR ORACLE SUBJECT")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "A fuzzer for algorithmic competition programs")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Arguments:")
	fmt.Fprintln(os.Stderr, "  GRAMMAR: the .moo file describing the format expected for stdin")
	fmt.Fprintln(os.Stderr, "  ORACLE:  path to a binary that must correctly answer the problem")
	fmt.Fprintln(os.Stderr, "  SUBJECT: path to a binary to compare to the oracle")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
