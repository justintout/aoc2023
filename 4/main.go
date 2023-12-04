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
	n := c.haveWinning()
	for i := 0; i < n; i++ {
		if i == 0 {
			v = 1
			continue
		}
		v *= 2
	}
	return v
}

func (c card) haveWinning() (n int) {
	for _, h := range c.have {
		for _, w := range c.winning {
			if h == w {
				n++
			}
		}
	}
	return n
}

func (c card) wonNumbers() (won []int) {
	n := c.haveWinning()
	for i := 1; i < n+1; i++ {
		won = append(won, c.number+i)
	}
	return won
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

func part2(cardList []card) (total int) {
	cards := make(map[int][]card)
	for _, c := range cardList {
		cards[c.number] = []card{c}
	}

	insertWon := func(c card) {
		wn := c.wonNumbers()
		for _, n := range wn {
			cards[n] = append(cards[n], cards[n][0])
		}
	}

	for _, c := range cardList {
		for _, c := range cards[c.number] {
			insertWon(c)
		}
	}

	for _, c := range cards {
		total += len(c)
	}
	return total
}
