package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id    int
	draws []draw
}

func newGame(s string) (game, error) {
	ss := strings.Split(s, ":")
	ids := strings.TrimPrefix(ss[0], "Game ")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return game{}, err
	}
	var draws []draw
	for _, d := range strings.Split(ss[1], ";") {
		dd, err := newDraw(d)
		if err != nil {
			return game{}, err
		}
		draws = append(draws, dd)
	}
	return game{id: id, draws: draws}, nil
}

type draw struct {
	blue  int
	red   int
	green int
}

func newDraw(s string) (draw, error) {
	rs := redRegex.FindStringSubmatch(s)
	if len(rs) < 2 {
		rs = []string{"", "0"}
	}
	r, err := strconv.Atoi(rs[1])
	if err != nil {
		return draw{}, err
	}

	gs := greenRegex.FindStringSubmatch(s)
	if len(gs) < 2 {
		gs = []string{"", "0"}
	}
	g, err := strconv.Atoi(gs[1])
	if err != nil {
		return draw{}, err
	}

	bs := blueRegex.FindStringSubmatch(s)
	if len(bs) < 2 {
		bs = []string{"", "0"}
	}
	b, err := strconv.Atoi(bs[1])
	if err != nil {
		return draw{}, err
	}

	return draw{red: r, green: g, blue: b}, nil
}

var (
	greenRegex = regexp.MustCompile("([0-9]+) green")
	redRegex   = regexp.MustCompile("([0-9]+) red")
	blueRegex  = regexp.MustCompile("([0-9]+) blue")
)

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
	var games []game
	for i, l := range lines {
		g, err := newGame(l)
		if err != nil {
			panic(fmt.Sprintf("line %d: %v", i, err))
		}
		games = append(games, g)
	}

	p1, err := part1(games)
	if err != nil {
		panic("part 1: " + err.Error())
	}
	fmt.Printf("part 1: %d\n", p1)

	p2, err := part2(games)
	if err != nil {
		panic("part 2: " + err.Error())
	}
	fmt.Printf("part 2: %d\n", p2)
}

func part1(games []game) (total int, err error) {
	var (
		red   = 12
		green = 13
		blue  = 14
	)
GAME:
	for _, g := range games {
		for _, d := range g.draws {
			if d.red > red {
				continue GAME
			}
			if d.green > green {
				continue GAME
			}
			if d.blue > blue {
				continue GAME
			}
		}
		total += g.id
	}

	return total, nil
}

func part2(games []game) (total int, err error) {
	for _, g := range games {
		var minr, ming, minb int
		for _, d := range g.draws {
			if d.red > minr {
				minr = d.red
			}
			if d.green > ming {
				ming = d.green
			}
			if d.blue > minb {
				minb = d.blue
			}
		}
		total += power(minr, ming, minb)
	}
	return total, nil
}

func power(red, green, blue int) int {
	return red * green * blue
}
