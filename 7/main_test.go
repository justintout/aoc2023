package main

import (
	"strings"
	"testing"
)

var example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPart1(t *testing.T) {
	expected := 6440
	in := parse(strings.Split(example, "\n"))
	out := part1(in)
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}

func TestPart2(t *testing.T) {
	expected := 5905
	in := parse(strings.Split(example, "\n"))
	out := part2(in)
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}
