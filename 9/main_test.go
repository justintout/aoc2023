package main

import (
	"strings"
	"testing"
)

var example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
4 22 46 72 96 114`

func TestPart1(t *testing.T) {
	expected := 236
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
