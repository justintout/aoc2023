package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type history []int

func (h history) solve() [][]int {
	diffs := [][]int{h}
	row := 1
	done := false
	for !done {
		diffs = append(diffs, make([]int, len(diffs[row-1])-1))
		for i := 1; i < len(diffs[row-1]); i++ {
			d := diffs[row-1][i] - diffs[row-1][i-1]
			diffs[row][i-1] = d
		}
		done = allZero(diffs[row])
		row++
	}
	return diffs
}

func (h history) extrapolate() []int {
	diffs := h.solve()
	slices.Reverse(diffs)
	for i, d := range diffs {
		if i == 0 {
			diffs[i] = append(d, 0)
			continue
		}
		diffs[i] = append(d, diffs[i-1][len(diffs[i-1])-1]+d[len(d)-1])
	}
	slices.Reverse(diffs)
	// fmt.Println(diffs)
	fmt.Println(diffs[0])
	return diffs[0]
}

func (h history) next() int {
	e := h.extrapolate()
	return e[len(e)-1]
}

var numberRegex = regexp.MustCompile(`-?\d+`)

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
	h := parse(lines)

	p1 := part1(h)
	fmt.Printf("part 1: %d\n", p1)

	p2 := part2(h)
	fmt.Printf("part 2: %d\n", p2)
}

func parse(lines []string) []history {
	var ints []history
	for _, l := range lines {
		var i []int
		ns := numberRegex.FindAllString(l, -1)
		for _, s := range ns {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			i = append(i, n)
		}
		ints = append(ints, i)
	}
	return ints
}

func part1(histories []history) (total int) {
	for _, h := range histories {
		n := h.next()
		total += n
	}
	return total
}

func part2(histories []history) (total int) {
	return total
}

func allZero(ints []int) bool {
	for _, x := range ints {
		if x != 0 {
			return false
		}
	}
	return true
}
