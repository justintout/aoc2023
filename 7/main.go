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

var cardValuesWithJoker map[string]int = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
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

func (b biddedHand) handTypeWithJoker() handType {
	var (
		four   bool
		three  bool
		pairs  int
		jokers int
	)
	for v, c := range b.cardCounts() {
		if c == 5 {
			return fiveOfKind
		}
		if v == "J" {
			if c == 4 {
				return fiveOfKind
			}
			jokers = c
			continue
		}
		if c == 4 {
			four = true
			continue
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
	case four && jokers == 1, three && jokers == 2, pairs == 1 && jokers == 3:
		return fiveOfKind
	case four, three && jokers == 1, pairs == 1 && jokers == 2, jokers == 3:
		return fourOfKind
	case three && pairs == 1, pairs == 2 && jokers == 1:
		return fullHouse
	case three, pairs == 1 && jokers == 1, jokers == 2:
		return threeOfKind
	case pairs == 2, pairs == 1 && jokers == 1:
		return twoPair
	case pairs == 1, jokers == 1:
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

	p2 := part2(hands)
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
		total += h.value(i + 1)
	}
	return total
}

func part2(hands []biddedHand) (total int) {
	slices.SortFunc(hands, cmpWithJoker)
	for i, h := range hands {
		total += h.value(i + 1)
	}
	return total
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

func cmpWithJoker(a, b biddedHand) int {
	at := a.handTypeWithJoker()
	bt := b.handTypeWithJoker()
	if at < bt {
		return -1
	}
	if at > bt {
		return 1
	}
	if at == bt {
		for i, v := range a.hand {
			if cardValuesWithJoker[v] < cardValuesWithJoker[b.hand[i]] {
				return -1
			}
			if cardValuesWithJoker[v] > cardValuesWithJoker[b.hand[i]] {
				return 1
			}
		}
	}
	return 0
}
