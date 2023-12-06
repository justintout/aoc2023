package main

import (
	"strings"
	"testing"
)

var example = `Time:      7  15   30
Distance:  9  40  200`

func testPart1(t *testing.T) {
	in := parseRaces(strings.Split(example, "\n"))
	expected := 288
	out := part1(in)
	if out != expected {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}

func testPart2(t *testing.T) {
	in := strings.Split(example, "\n")
	expected := 71503
	out := part2(in)
	if out != expected {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}
