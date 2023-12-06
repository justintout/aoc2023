package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numRegex = regexp.MustCompile(`\d+`)

type race struct {
	time     int
	distance int
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

	races := parseRaces(lines)

	fmt.Println("part 1: ", part1(races))
	fmt.Println("part 2: ", part2(lines))
}

func parseRaces(lines []string) []race {
	var times []int
	ts := numRegex.FindAllString(lines[0], -1)
	for _, s := range ts {
		t, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		times = append(times, t)
	}
	var distances []int
	ds := numRegex.FindAllString(lines[1], -1)
	for _, s := range ds {
		d, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		distances = append(distances, d)
	}

	if len(ts) != len(ds) {
		panic("time/distance list different lengths")
	}

	var races []race
	for i, t := range times {
		races = append(races, race{t, distances[i]})
	}
	return races
}

func part1(races []race) (margin int) {
	var wins []int
	for _, r := range races {
		var w []int
		var hold int
		for hold = 1; hold < r.time-1; hold++ {
			d := hold * (r.time - hold)
			if d > r.distance {
				w = append(w, hold)
			}
		}
		wins = append(wins, len(w))
	}
	margin = 1
	for _, w := range wins {
		margin *= w
	}
	return margin
}

func part2(lines []string) (margin int) {
	ts := strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time:"), " ", "")
	time, err := strconv.Atoi(ts)
	if err != nil {
		panic(err)
	}
	ds := strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance:"), " ", "")
	distance, err := strconv.Atoi(ds)
	if err != nil {
		panic(err)
	}

	var ways int
	for hold := 1; hold < time-1; hold++ {
		d := hold * (time - hold)
		if d > distance {
			ways++
		}
	}
	return ways
}
