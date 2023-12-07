package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type handType int

const (
	highCard handType = iota
	pair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
)

var cardValues map[string]int = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type biddedHand struct {
	hand []string
	bid  int
}

func (b biddedHand) value(rank int) int {
	return b.bid * rank
}

func (b biddedHand) cardCounts() map[string]int {
	counts := make(map[string]int)
	for _, c := range b.hand {
		if _, ok := counts[c]; ok {
			counts[c]++
			continue
		}
		counts[c] = 1
	}
	return counts
}

func (b biddedHand) handType() handType {
	var (
		three bool
		pairs int
	)
	for _, c := range b.cardCounts() {
		if c == 5 {
			return fiveOfKind
		}
		if c == 4 {
			return fourOfKind
		}
		if c == 3 {
			three = true
			continue
		}
		if c == 2 {
			pairs++
			continue
		}
	}
	switch {
	case three && pairs > 0:
		return fullHouse
	case three && pairs == 0:
		return threeOfKind
	case !three && pairs == 2:
		return twoPair
	case !three && pairs == 1:
		return pair
	default:
		return highCard
	}
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
	hands := parse(lines)

	p1 := part1(hands)
	fmt.Printf("part 1: %d\n", p1)

	p2, err := part2(hands)
	if err != nil {
		panic("part 2: " + err.Error())
	}
	fmt.Printf("part 2: %d\n", p2)
}

func parse(lines []string) []biddedHand {
	var hands []biddedHand
	for _, l := range lines {
		s := strings.Split(l, " ")
		b, err := strconv.Atoi(s[1])
		if err != nil {
			panic("failed to parse bid: " + err.Error())
		}
		hands = append(hands, biddedHand{
			hand: strings.Split(s[0], ""),
			bid:  b,
		})
	}
	return hands
}

func part1(hands []biddedHand) (total int) {
	slices.SortFunc(hands, cmp)
	for i, h := range hands {
		total += h.bid * (i + 1)
	}
	return total
}

func part2(hands []biddedHand) (total int, err error) {
	return total, nil
}

func cmp(a, b biddedHand) int {
	at := a.handType()
	bt := b.handType()
	if at < bt {
		return -1
	}
	if at > bt {
		return 1
	}
	if at == bt {
		for i, v := range a.hand {
			if cardValues[v] < cardValues[b.hand[i]] {
				return -1
			}
			if cardValues[v] > cardValues[b.hand[i]] {
				return 1
			}
		}
	}
	return 0
}
