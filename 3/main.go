package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	p1, err := part1(lines)
	if err != nil {
		panic("part 1: " + err.Error())
	}
	fmt.Printf("part 1: %d\n", p1)

	p2, err := part2(lines)
	if err != nil {
		panic("part 2: " + err.Error())
	}
	fmt.Printf("part 2: %d\n", p2)
}

func part1(lines []string) (total int, err error) {
	n := parse(lines)
	for _, nn := range n {
		total += nn.number
	}
	return total, nil
}

func part2(lines []string) (total int, err error) {
	parts := parse(lines)
	gears := make(map[point][]part)
	for _, p := range parts {
	GEAR:
		for _, g := range p.gears {
			if pp, ok := gears[g]; ok {
				for _, ppp := range pp {
					if ppp.number == p.number {
						continue GEAR
					}
				}
			}
			gears[g] = append(gears[g], p)
		}
	}
	for _, parts := range gears {
		if len(parts) != 2 {
			continue
		}
		total += parts[0].number * parts[1].number
	}
	return total, nil
}

type part struct {
	number int
	gears  []point
}

type point struct {
	y int
	x int
}

// incredibly janky character parser
// trying not to go full blown rsc with it
func parse(lines []string) (parts []part) {
	var characters [][]string
	for _, l := range lines {
		s := strings.Split(l, "")
		characters = append(characters, s)
	}

	var (
		height               = len(characters) - 1
		width                = len(characters[0]) - 1
		currentNumber        string
		inNumber             bool
		currentTouchesSymbol bool
		currentGears         []point = nil
	)

	reset := func() {
		currentNumber = ""
		inNumber = false
		currentGears = nil
		currentTouchesSymbol = false
	}

	contiguous := func(y, x int) []point {
		if y == 0 && x == 0 {
			return []point{
				{0, 1}, {1, 0}, {1, 1},
			}
		}
		if y == 0 && x == width {
			return []point{
				{0, width - 1}, {1, width - 1}, {1, width},
			}
		}
		if y == height && x == 0 {
			return []point{
				{height - 1, 0}, {height - 1, 1}, {height, 1},
			}
		}
		if y == height && x == width {
			return []point{
				{height, width - 1}, {height - 1, width - 1}, {height - 1, width},
			}
		}
		if y == 0 {
			return []point{
				{0, x - 1}, {0, x + 1}, {1, x - 1}, {1, x}, {1, x + 1},
			}
		}
		if y == height {
			return []point{
				{height, x - 1}, {height, x + 1}, {height - 1, x - 1}, {height - 1, x}, {height - 1, x + 1},
			}
		}
		if x == 0 {
			return []point{
				{y - 1, 0}, {y - 1, 1}, {y, 1}, {y + 1, 0}, {y + 1, 1},
			}
		}
		if x == width {
			return []point{
				{y - 1, 0}, {y - 1, width - 1}, {y, width - 1}, {y + 1, width}, {y + 1, width - 1},
			}
		}
		return []point{
			{y - 1, x - 1},
			{y - 1, x},
			{y - 1, x + 1},
			{y, x - 1},
			{y, x + 1},
			{y + 1, x - 1},
			{y + 1, x},
			{y + 1, x + 1},
		}
	}

	touchesSymbol := func(y, x int) bool {
		for _, p := range contiguous(y, x) {
			if isSymbol(characters[p.y][p.x]) {
				return true
			}
		}
		return false
	}

	touchingGears := func(y, x int) (gears []point) {
		for _, p := range contiguous(y, x) {
			if isGear(characters[p.y][p.x]) {
				gears = append(gears, p)
			}
		}
		return gears
	}

	for y, l := range characters {
		for x, c := range l {
			if (!inNumber && !isDigit(c)) || (inNumber && !currentTouchesSymbol && !isDigit(c)) {
				reset()
				continue
			}
			if inNumber && !isDigit(c) && currentTouchesSymbol {
				pn, err := strconv.Atoi(currentNumber)
				if err != nil {
					panic(fmt.Sprintf("y: %d, x: %d, err: %v", y, x, err))
				}
				parts = append(parts, part{
					number: pn,
					gears:  currentGears,
				})
				reset()
				continue
			}
			if inNumber && isDigit(c) {
				currentNumber += c
				if !currentTouchesSymbol && touchesSymbol(y, x) {
					currentTouchesSymbol = true
				}
				currentGears = append(currentGears, touchingGears(y, x)...)
				continue
			}
			if isDigit(c) {
				inNumber = true
				currentNumber += c
				currentTouchesSymbol = touchesSymbol(y, x)
				currentGears = append(currentGears, touchingGears(y, x)...)
			}
		}
		if inNumber && currentTouchesSymbol {
			pn, err := strconv.Atoi(currentNumber)
			if err != nil {
				panic(fmt.Sprintf("y: %d, x: %d, err: %v", y, width, err))
			}
			parts = append(parts, part{
				number: pn,
				gears:  currentGears,
			})
		}
		reset()
	}

	return parts
}

func isDigit(c string) bool {
	return c == "1" ||
		c == "2" ||
		c == "3" ||
		c == "4" ||
		c == "5" ||
		c == "6" ||
		c == "7" ||
		c == "8" ||
		c == "9" ||
		c == "0"
}

func isSymbol(c string) bool {
	return !isDigit(c) && !(c == ".")
}

func isGear(c string) bool {
	return c == "*"
}
