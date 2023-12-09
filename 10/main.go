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
