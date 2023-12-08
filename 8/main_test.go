package main

import (
	"strings"
	"testing"
)

var example = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func TestPart1(t *testing.T) {
	expected := 6
	ins, nodes := parse(strings.Split(example, "\n"))
	out := part1(ins, nodes)
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	ins, nodes := parse(strings.Split(example, "\n"))
	out, err := part2(ins, nodes)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}
