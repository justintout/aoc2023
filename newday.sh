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

	p1, err := part1(ss)
	if err != nil {
		panic("part 1 error: " + err.Error())
	}
	fmt.Printf("part 1: %d\n", p1)

	p2, err := part2(ss)
	if err != nil {
		panic("part 2 error: " + err.Error())
	}
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

func part1(lines []something) (total int, err error) {
	return total, nil
}

func part2(lines []something) (total int, err error) {
	return total, nil
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
	out, err := part1(in)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	in := parse(strings.Split(example, "\n"))
	out, err := part2(in)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if expected != out {
		t.Errorf("expected: %d, got: %d", expected, out)
	}
}
EOF