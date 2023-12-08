package main

import (
	"strings"
	"testing"
)

var example = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

var example2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestPart1(t *testing.T) {
	expected := 6
	ins, nodes := parse(strings.Split(example, "\n"))
	out := part1(ins, nodes)
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}

func TestPart2(t *testing.T) {
	expected := 6
	ins, nodes := parse(strings.Split(example2, "\n"))
	out := part2(ins, nodes)
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}
