#!/bin/bash

mkdir $1

touch $1/input.txt

cat <<EOF >$1/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

type something struct{}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	ss := parse(lines)

	p1 := part1(ss)
	fmt.Printf("part 1: %d\n", p1)

	p2 := part2(ss)
	fmt.Printf("part 2: %d\n", p2)
}

func parse(lines []string) []something {
	var s []something
	for _, l := range lines {
		fmt.Printf("parse line: %v", l)
		s = append(s, something{})
	}
	return s
}

func part1(lines []something) (total int) {
	return total
}

func part2(lines []something) (total int) {
	return total
}
EOF

cat <<EOF >$1/main_test.go
package main

import (
	"strings"
	"testing"
)

var example = \`\`

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
EOF