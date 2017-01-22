package main

import (
	"context"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func runTest(grammar, oracle, test string, n int) ([]byte, error) {
	ctx, done := context.WithTimeout(context.Background(), 5*time.Minute)
	defer done()

	cmd := exec.CommandContext(
		ctx,
		"go",
		"run",
		"main.go",
		"-n",
		strconv.Itoa(n),
		filepath.Join("testdata", grammar),
		"go run "+filepath.Join("testdata", oracle),
		"go run "+filepath.Join("testdata", test),
	)

	return cmd.CombinedOutput()
}

func TestRunOK(t *testing.T) {
	data, err := runTest("g1.moo", "c1_ok.go", "c1_ok.go", 50)
	if err != nil {
		t.Error("unexpected error:", err, "\n", string(data[:]))
	}

	if strings.Contains(string(data), "ERROR") {
		t.Error("got an error\n", string(data[:]))
	}
}

func TestRunKO(t *testing.T) {
	data, err := runTest("g1.moo", "c1_ok.go", "c1_ko.go", 50)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if !strings.Contains(string(data), "ERROR") {
		t.Error("expected ERROR string in output")
	}
}

func TestRunSyntaxError(t *testing.T) {
	data, err := runTest("gerr1.moo", "c1_ok.go", "c1_ko.go", 50)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if !strings.Contains(string(data), "expecting ENDLOOP at line 3") {
		t.Error("expected syntax error string in output")
	}
}

func TestRunInvalidFile(t *testing.T) {
	data, err := runTest("invalid_file.moo", "c1_ok.go", "c1_ko.go", 50)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if !strings.Contains(string(data), "no such file or directory") {
		t.Error("expected io error message output")
	}
}
