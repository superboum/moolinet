package main

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func runTest(grammar, oracle, test string, n int) ([]byte, error) {
	// Build binaries
	f1, _ := ioutil.TempFile("", "")
	f2, _ := ioutil.TempFile("", "")
	_, _ = exec.Command("go", "build", "-o", f1.Name(), filepath.Join("testdata", oracle+".go")).CombinedOutput()
	_, _ = exec.Command("go", "build", "-o", f2.Name(), filepath.Join("testdata", test+".go")).CombinedOutput()

	// Prepare cleanup
	defer func() { _ = os.Remove(f1.Name()) }()
	defer func() { _ = os.Remove(f2.Name()) }()

	// Run command
	ctx, done := context.WithTimeout(context.Background(), 30*time.Second)
	defer done()

	cmd := exec.CommandContext(
		ctx,
		"go",
		"run",
		"main.go",
		"-n",
		strconv.Itoa(n),
		filepath.Join("testdata", grammar),
		f1.Name(),
		f2.Name(),
	)

	return cmd.CombinedOutput()
}

func TestRunOK(t *testing.T) {
	cases := []string{"1", "2"}

	for _, c := range cases {
		data, err := runTest("g"+c+".moo", "c"+c+"_ok", "c"+c+"_ok", 100)
		if err != nil {
			t.Error("unexpected error:", err, "in case", c)
		}

		if strings.Contains(string(data), "ERROR") {
			t.Error("got an error in case", c)
		}
	}
}

func TestRunKO(t *testing.T) {
	cases := []string{"1", "2"}

	for _, c := range cases {
		data, err := runTest("g"+c+".moo", "c"+c+"_ok", "c"+c+"_ko", 100)
		if err == nil {
			t.Error("expected error, got nil in case", c)
		}

		if !strings.Contains(string(data), "ERROR") {
			t.Error("expected ERROR string in output in case", c)
		}
	}
}

func TestRunSyntaxError(t *testing.T) {
	data, err := runTest("gerr1.moo", "c1_ok", "c1_ko", 100)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if !strings.Contains(string(data), "expecting ENDLOOP at line 3") {
		t.Error("expected syntax error string in output")
	}
}

func TestRunInvalidFile(t *testing.T) {
	data, err := runTest("invalid_file.moo", "c1_ok", "c1_ko", 100)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if !strings.Contains(string(data), "no such file or directory") {
		t.Error("expected io error message output")
	}
}
