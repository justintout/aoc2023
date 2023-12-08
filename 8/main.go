package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cznic/mathutil"
)

var idRegex = regexp.MustCompile(`[0-9A-Z]{3}`)

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

	p2 := part2(ins, nodes)
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

// each starting node's solution is a factor of the total,
// so find the LCM of all of the starting node's steps.
// we'll use prime factorization
func part2(ins *ring.Ring, nodes map[string]node) (total int) {
	var startingNodes []node
	for id, n := range nodes {
		if strings.HasSuffix(id, "A") {
			startingNodes = append(startingNodes, n)
		}
	}
	var steps []int
	for _, n := range startingNodes {
		ins := ins
		current := n
		s := 0
		for !strings.HasSuffix(current.id, "Z") {
			if ins.Value == "L" {
				current = nodes[current.left]
			}
			if ins.Value == "R" {
				current = nodes[current.right]
			}
			ins = ins.Next()
			s++
		}
		steps = append(steps, s)
	}
	factors := make(map[int]struct{})
	for _, s := range steps {
		f := mathutil.FactorInt(uint32(s))
		for _, x := range f {
			factors[int(x.Prime)] = struct{}{}
		}
	}
	total = 1
	for x := range factors {
		total *= x
	}
	return total
}
