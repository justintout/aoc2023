package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile(`\d+`)

type card struct {
	number  int
	winning []int
	have    []int
}

func (c card) value() (v int) {
	for _, h := range c.have {
		for _, w := range c.winning {
			if h == w && v == 0 {
				v = 1
				continue
			}
			if h == w {
				v *= 2
			}
		}
	}
	return v
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
	var cards []card
	for _, l := range lines {
		s := strings.Split(l, ":")
		number, err := strconv.Atoi(numberRegex.FindString(s[0]))
		if err != nil {
			panic(err)
		}
		c := card{
			number: number,
		}
		s = strings.Split(s[1], "|")
		ws := numberRegex.FindAllString(s[0], -1)
		for _, n := range ws {
			w, err := strconv.Atoi(strings.TrimSpace(n))
			if err != nil {
				panic(err)
			}
			c.winning = append(c.winning, w)
		}
		hs := numberRegex.FindAllString(s[1], -1)
		for _, n := range hs {
			h, err := strconv.Atoi(strings.TrimSpace(n))
			if err != nil {
				panic(err)
			}
			c.have = append(c.have, h)
		}
		cards = append(cards, c)
	}

	p1 := part1(cards)
	if err != nil {
		panic("part 1: " + err.Error())
	}
	fmt.Printf("part 1: %d\n", p1)

	p2 := part2(cards)
	if err != nil {
		panic("part 2: " + err.Error())
	}
	fmt.Printf("part 2: %d\n", p2)
}

func part1(cards []card) (total int) {
	for _, c := range cards {
		total += c.value()
	}
	return total
}

func part2(cards []card) int {
	return 0
}
