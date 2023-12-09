package main

import (
	"strings"
	"testing"
)

var example = ``

func TestPart1(t *testing.T) {
	expected := 0
	in := parse(strings.Split(example, "\n"))
	out := part1(in)
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	in := parse(strings.Split(example, "\n"))
	out := part2(in)
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}
