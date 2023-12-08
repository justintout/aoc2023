package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var idRegex = regexp.MustCompile(`[A-Z]{3}`)

type instructions *ring.Ring

func newInstructions(s string) instructions {
	r := ring.New(len(s))
	for _, v := range strings.Split(s, "") {
		r.Value = v
		r = r.Next()
	}
	return r
}

type node struct {
	id    string
	left  string
	right string
}

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
	ins, nodes := parse(lines)

	p1 := part1(ins, nodes)
	fmt.Printf("part 1: %d\n", p1)

	p2, err := part2(ins, nodes)
	if err != nil {
		panic("part 2 error: " + err.Error())
	}
	fmt.Printf("part 2: %d\n", p2)
}

func parse(lines []string) (instructions, map[string]node) {
	ins := newInstructions(lines[0])

	n := make(map[string]node)
	for _, l := range lines[2:] {
		ids := idRegex.FindAllString(l, -1)
		n[ids[0]] = node{
			id:    ids[0],
			left:  ids[1],
			right: ids[2],
		}
	}
	return ins, n
}

func part1(ins *ring.Ring, nodes map[string]node) (total int) {
	current := nodes["AAA"]
	steps := 0
	for current.id != "ZZZ" {
		if ins.Value == "L" {
			current = nodes[current.left]
		}
		if ins.Value == "R" {
			current = nodes[current.right]
		}
		ins = ins.Next()
		steps++
	}
	return steps
}

func part2(ins instructions, nodes map[string]node) (total int, err error) {
	return total, nil
}
